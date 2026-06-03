//go:build integration_runtime

package runtime

import "time"

const (
	testNamespace = "agent-runtime-it"
	testAgentName = "it-agent"

	clusterProvisionTimeout = 15 * time.Minute
	readyTimeout            = 5 * time.Minute
)

// testAgentManifest is a placeholder workload. TODO: generate this from a
// foundation.v1.RuntimeSpec instead of hardcoding it.
const testAgentManifest = `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: it-agent
  namespace: agent-runtime-it
spec:
  replicas: 1
  selector:
    matchLabels: {app: it-agent}
  template:
    metadata:
      labels: {app: it-agent}
    spec:
      containers:
        - name: agent
          image: ghcr.io/agennext/it-agent:latest
          resources:
            requests: {cpu: "100m", memory: "128Mi"}
`
