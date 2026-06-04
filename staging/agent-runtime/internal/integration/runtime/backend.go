//go:build integration_runtime

// Package runtime holds Agent-Runtime integration tests.
//
// The design mirrors Talos' storage tests (one suite, swappable backend), but
// the swappable axis here is the cluster provider: a multi-cloud registry of
// ClusterBackends. The same RuntimeSuite runs against whichever is selected.
package runtime

import (
	"context"
	"fmt"
	"os"
)

// ClusterBackend provisions a Kubernetes cluster on some substrate and yields a
// path to a kubeconfig the suite can use. Implementations are swappable —
// adding a new cloud means adding a backend, not changing the suite.
type ClusterBackend interface {
	// Name is the registry key, e.g. "kind", "eks".
	Name() string
	// Provision brings up a cluster and returns a kubeconfig path.
	Provision(ctx context.Context) (kubeconfig string, err error)
	// Teardown destroys everything Provision created.
	Teardown(ctx context.Context) error
}

// backends is the multi-cloud cluster registry.
var backends = map[string]ClusterBackend{}

// RegisterBackend adds a backend to the registry. Called from each backend's
// init() so importing this package wires them all up.
func RegisterBackend(b ClusterBackend) {
	backends[b.Name()] = b
}

// SelectedBackend returns the backend named by AGENT_RUNTIME_BACKEND, defaulting
// to "kind" for local runs.
func SelectedBackend() (ClusterBackend, error) {
	name := os.Getenv("AGENT_RUNTIME_BACKEND")
	if name == "" {
		name = "kind"
	}

	b, ok := backends[name]
	if !ok {
		return nil, fmt.Errorf("unknown AGENT_RUNTIME_BACKEND %q; registered: %v", name, registered())
	}

	return b, nil
}

func registered() []string {
	names := make([]string, 0, len(backends))
	for n := range backends {
		names = append(names, n)
	}

	return names
}
