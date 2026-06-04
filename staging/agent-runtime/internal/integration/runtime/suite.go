//go:build integration_runtime

package runtime

import (
	"context"
	"os/exec"
	"strings"
	"time"

	"github.com/stretchr/testify/suite"
)

// RuntimeSuite is written once and run against every ClusterBackend. It asserts
// the behaviour the RuntimeSpec contract promises, independent of which cloud
// provisioned the cluster.
type RuntimeSuite struct {
	suite.Suite

	backend    ClusterBackend
	kubeconfig string
	ctx        context.Context
	cancel     context.CancelFunc
}

func (s *RuntimeSuite) SetupSuite() {
	s.ctx, s.cancel = context.WithTimeout(context.Background(), clusterProvisionTimeout)

	kubeconfig, err := s.backend.Provision(s.ctx)
	s.Require().NoError(err, "provision %s cluster", s.backend.Name())
	s.kubeconfig = kubeconfig
}

func (s *RuntimeSuite) TearDownSuite() {
	if s.backend != nil {
		s.NoError(s.backend.Teardown(context.Background()), "teardown %s cluster", s.backend.Name())
	}

	if s.cancel != nil {
		s.cancel()
	}
}

// TestAgentSchedules applies an Agent derived from a RuntimeSpec and asserts the
// workload becomes ready. Same assertion for kind, EKS, GKE, AKS.
func (s *RuntimeSuite) TestAgentSchedules() {
	// TODO: render a Deployment from a foundation.v1.RuntimeSpec (image,
	// resources, replicas) and apply it.
	s.applyManifest(testAgentManifest)
	s.Require().True(
		s.waitFor("deployment/"+testAgentName, "condition=Available", readyTimeout),
		"agent should become Available on %s", s.backend.Name(),
	)
}

// TestScaleToZero asserts an ephemeral runtime (min_replicas=0) scales down when
// idle — the disposable, on-demand runtime promise.
func (s *RuntimeSuite) TestScaleToZero() {
	s.T().Skip("TODO: drive autoscaler to 0 and assert no running pods")
}

// --- helpers (shell out to kubectl to keep the scaffold dependency-light) ---

func (s *RuntimeSuite) kubectl(args ...string) (string, error) {
	cmd := exec.CommandContext(s.ctx, "kubectl", append([]string{"--kubeconfig", s.kubeconfig}, args...)...)
	out, err := cmd.CombinedOutput()

	return strings.TrimSpace(string(out)), err
}

func (s *RuntimeSuite) applyManifest(manifest string) {
	cmd := exec.CommandContext(s.ctx, "kubectl", "--kubeconfig", s.kubeconfig, "apply", "-f", "-")
	cmd.Stdin = strings.NewReader(manifest)
	out, err := cmd.CombinedOutput()
	s.Require().NoError(err, "kubectl apply: %s", out)
}

func (s *RuntimeSuite) waitFor(resource, condition string, timeout time.Duration) bool {
	_, err := s.kubectl("wait", resource, "--for", condition, "--timeout", timeout.String())

	return err == nil
}
