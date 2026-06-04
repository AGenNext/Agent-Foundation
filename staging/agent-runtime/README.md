# Agent-Runtime (scaffold)

> Staged in Agent-Foundation for hand-off. Move this folder to
> `AGenNext/Agent-Runtime`.

Agent-Runtime is the reference implementation of the `RuntimeSpec` contract
(`runtime.proto` in Agent-Foundation): it launches and manages agents across
substrates (Kubernetes, serverless, wasm, dapr, edge), including disposable,
scale-to-zero runtimes.

## Integration tests: one suite, swappable backend

Modeled on
[Talos' k8s integration tests](https://github.com/siderolabs/talos/tree/main/internal/integration/k8s).
Talos runs **one storage suite** against **swappable storage backends**
(Longhorn / OpenEBS / Rook). We apply the same idea on a different axis:

```
Talos:         one storage suite   ×  backend ∈ {longhorn, openebs, rook}
Agent-Runtime: one runtime suite   ×  backend ∈ {kind, eks, gke, aks}
                                       └─ a multi-cloud cluster registry
```

The `RuntimeSuite` is written once and asserts the same behaviour — an `Agent`
scheduled from a `RuntimeSpec` becomes ready, then scales to zero — regardless
of which cloud provisioned the cluster. Backends are swappable via a registry
and selected with `AGENT_RUNTIME_BACKEND` (default `kind`).

## Layout

```
internal/integration/runtime/
  runtime.go        build-tagged entrypoint; registers backends, runs the suite
  suite.go          the shared RuntimeSuite (same assertions for every backend)
  backend.go        ClusterBackend interface + registry (the cluster registry)
  backend_kind.go   local kind backend
  backend_cloud.go  EKS/GKE/AKS backends (stubs)
  constants.go      namespace, timeouts
go.mod
```

## Running

```bash
# default local backend
go test -tags integration_runtime ./internal/integration/runtime/...

# swap the backend — same suite, different cloud
AGENT_RUNTIME_BACKEND=eks go test -tags integration_runtime ./...
```

## Status

Skeleton with `TODO`s. The suite shells out to `kubectl` against the
backend-provided kubeconfig to avoid heavy client deps in the scaffold; swap for
client-go in the real implementation.
