//go:build integration_runtime

package runtime

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func init() {
	RegisterBackend(&kindBackend{clusterName: "agent-runtime-it"})
}

// kindBackend provisions a local cluster with kind. It is the default backend,
// the equivalent of running Talos' suite against the simplest driver.
type kindBackend struct {
	clusterName string
	kubeconfig  string
}

func (b *kindBackend) Name() string { return "kind" }

func (b *kindBackend) Provision(ctx context.Context) (string, error) {
	b.kubeconfig = filepath.Join(os.TempDir(), b.clusterName+".kubeconfig")

	if out, err := run(ctx, "kind", "create", "cluster",
		"--name", b.clusterName, "--kubeconfig", b.kubeconfig); err != nil {
		return "", fmt.Errorf("kind create cluster: %w: %s", err, out)
	}

	if out, err := run(ctx, "kubectl", "--kubeconfig", b.kubeconfig,
		"create", "namespace", testNamespace); err != nil {
		return "", fmt.Errorf("create namespace: %w: %s", err, out)
	}

	return b.kubeconfig, nil
}

func (b *kindBackend) Teardown(ctx context.Context) error {
	_, err := run(ctx, "kind", "delete", "cluster", "--name", b.clusterName)

	return err
}

func run(ctx context.Context, name string, args ...string) (string, error) {
	out, err := exec.CommandContext(ctx, name, args...).CombinedOutput()

	return string(out), err
}
