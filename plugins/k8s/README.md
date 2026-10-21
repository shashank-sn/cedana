# Cedana Kubernetes helper plugin

Sets up the Cedana daemon on Kubernetes nodes. You do not run it by hand: the cedana-helper DaemonSet, installed by the [cedana-helm chart](https://github.com/cedana/cedana-helm-charts), runs `cedana k8s setup` on each node, and the chart's uninstaller runs `cedana k8s destroy`.

## Commands

| Command | What it does |
|---|---|
| `cedana k8s setup` | Sets up and starts Cedana on the node: installs dependencies, the Cedana binaries and plugins and the Cedana service, then starts the helper |
| `cedana k8s destroy` | Removes Cedana from the node and cleans up |

Both are in the [CLI reference](../../docs/references/cli/cedana_k8s.md).

The plugin also provides the query command that finds pods and containers on the node (see [`cedana features`](../../docs/get-started/features.md)).

## How a workload opts in

Add one environment variable to the container. Its value labels the workload's checkpoints, and applying the same YAML again starts the container from the latest checkpoint with that label:

```yaml
env:
  - name: CEDANA_CHECKPOINT
    value: my-workload
```

GPU pods also set `runtimeClassName: cedana`. Full examples: [jobs that resume after a restart](https://docs.cedana.ai/cedana-kubernetes/jobs).

## Install Cedana on a cluster

See the [Kubernetes installation docs](https://docs.cedana.ai/cedana-kubernetes/installation) and the [chart README](https://github.com/cedana/cedana-helm-charts).

## More

- [Cedana installation: what changes on your cluster](https://cedana.com/blog/cedana-installation-cluster-changes/)
- [Kubernetes GPU node drain: preserve running work](https://cedana.com/blog/kubernetes-gpu-node-drain/)
