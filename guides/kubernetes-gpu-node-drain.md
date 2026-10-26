# Drain a Kubernetes GPU node without losing the work on it

This guide checkpoints the pods on a GPU node before you drain it, so each pod restores on another node from its checkpoint instead of starting again. Use it for maintenance, driver upgrades and node replacement.

## What you need first

- A Kubernetes cluster with Cedana installed from the [Helm chart](https://github.com/cedana/cedana-helm-charts). The install needs a Cedana account. Reach out to us through your contact at Cedana or founders@cedana.ai to get an account!
- Workloads run by a controller that recreates their pods, such as a Deployment or a Job, each with a `CEDANA_CHECKPOINT` environment variable. GPU pods also set `runtimeClassName: cedana`.
- At least one other node with the same GPU type and driver, with room for the pods.

## Steps

1. Stop new pods from landing on the node:

   ```sh
   kubectl cordon <node>
   ```

2. Checkpoint each pod on the node, from the pods page of the Cedana dashboard or through the [Cedana API](https://demo.cedana.ai/v1/scalar). A heartbeat checkpoint policy works too, if its last checkpoint is recent enough.

3. Drain the node. The Cedana helper is a DaemonSet, so it stays on the node:

   ```sh
   kubectl drain <node> --ignore-daemonsets --delete-emptydir-data
   ```

4. Each controller creates a new pod on another node, and Cedana restores it from the latest checkpoint with its `CEDANA_CHECKPOINT` label.

5. After the maintenance, let pods land on the node again:

   ```sh
   kubectl uncordon <node>
   ```

## Check the result

`kubectl get pods -o wide` shows each pod on another node, and each pod's log carries on from where it was at the checkpoint.

## Limits today

- A checkpoint restores only where versions match. The other node needs the same GPU, driver, engine and model versions, or the workload cold-starts instead. Upgrade the driver on the drained node after the pods have moved off it.
- Work done between the checkpoint and the drain runs again.
- Effects outside the machine do not roll back.

## Read more

[Draining a GPU node in Kubernetes without losing the work on it](https://cedana.com/blog/kubernetes-gpu-node-drain/), and [automatic checkpoint and restore](https://docs.cedana.ai/cedana-kubernetes/automation).
