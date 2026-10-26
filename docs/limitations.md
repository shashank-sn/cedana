# Limits

What Cedana does not do today, and where each limit is written down in full.

## Job shape

- Cedana works with single-GPU, multi-GPU on a single node and multi-node GPUs. For a job that coordinates through NCCL, checkpoint with `--gpu-freeze-type nccl`. ([GPU plugin docs](https://docs.cedana.ai/daemon/checkpoint-restore/gpu))

## Where a checkpoint can be restored

- A checkpoint records the GPU, driver, engine and model versions it was taken against. If any of them changes, the checkpoint is invalid and the workload cold-starts instead. Applies to every job shape. ([post](https://cedana.com/blog/what-cedana-saves/))

## What a checkpoint does not contain

- Effects that already left the machine: tool calls, database writes, emails. They do not roll back, and the work after the checkpoint runs again. ([post](https://cedana.com/blog/gpu-checkpointing-limitations/))
- Work done after the last checkpoint. Checkpoints run on an interval the operator sets.

## Processes that cannot be checkpointed

- A process holding a raw hardware device other than the GPU, such as a capture card.
- A process under a debugger.
- A process holding a file that belongs to something outside what is being saved, until that file is closed.
- Cedana does not support io-uring based checkpoint restores.

## Hardware and software

- NVIDIA GPUs only.
- Driver range: see the GPU table in the [README](../README.md#gpus-and-drivers).
- Container runtimes: see the support table in the [README](../README.md#container-runtimes-and-vms). CRI-O is partially supported: its dump, restore and run commands are not implemented yet.

## Access and privileges

- The daemon requires root. On Kubernetes, the helper DaemonSet runs with privilege escalation, and the controller can list, checkpoint and restore all pods.
- Released plugins, including the GPU plugin, need a Cedana account.
