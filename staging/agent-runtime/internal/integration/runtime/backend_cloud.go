//go:build integration_runtime

package runtime

import (
	"context"
	"fmt"
)

// The multi-cloud entries of the cluster registry. Each is the analogue of a
// Talos storage backend: a swappable provider the one RuntimeSuite runs against
// unchanged. Implement Provision/Teardown per cloud (eksctl, gcloud, az).
func init() {
	RegisterBackend(&cloudBackend{name: "eks", tool: "eksctl"})
	RegisterBackend(&cloudBackend{name: "gke", tool: "gcloud"})
	RegisterBackend(&cloudBackend{name: "aks", tool: "az"})
}

type cloudBackend struct {
	name string
	tool string // provisioning CLI
}

func (b *cloudBackend) Name() string { return b.name }

func (b *cloudBackend) Provision(ctx context.Context) (string, error) {
	// TODO: provision a real cluster via b.tool and write a kubeconfig.
	// e.g. eks: `eksctl create cluster ...`; gke: `gcloud container clusters create ...`.
	return "", fmt.Errorf("%s backend not implemented yet (provision via %s)", b.name, b.tool)
}

func (b *cloudBackend) Teardown(ctx context.Context) error {
	// TODO: delete the cluster created in Provision.
	return fmt.Errorf("%s backend not implemented yet", b.name)
}
