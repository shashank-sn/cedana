# Cedana Kubernetes helper plugin

Sets up and maintains the Cedana daemon on Kubernetes nodes. You do not run it by hand: the cedana-helper DaemonSet, installed by the [cedana-helm chart](https://github.com/cedana/cedana-helm-charts), runs these commands on each node.

## Commands

| Command | What it does |
|---|---|
| `cedana k8s-helper --setup-host` | Prepares a node: installs prerequisites, copies binaries and sets up the Cedana service |
| `cedana k8s-helper destroy` | Cleans up a node after Cedana is stopped |
| `cedana k8s-helper --restart` | Restarts the Cedana service on the host |

`--setup-host` is not needed on nodes that run a Cedana machine image (AMI).

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
