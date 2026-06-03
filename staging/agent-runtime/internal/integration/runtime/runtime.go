//go:build integration_runtime

package runtime

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// TestRuntimeIntegration is the entrypoint. Backends register via init(); the
// suite runs against the one selected by AGENT_RUNTIME_BACKEND.
//
// Run with: go test -tags integration_runtime ./internal/integration/runtime/...
func TestRuntimeIntegration(t *testing.T) {
	backend, err := SelectedBackend()
	require.NoError(t, err)

	t.Logf("running runtime suite against backend %q", backend.Name())
	suite.Run(t, &RuntimeSuite{backend: backend})
}
