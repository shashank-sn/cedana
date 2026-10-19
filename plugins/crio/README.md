# Cedana CRI-O plugin

This plugin will add CRI-O support to the Cedana daemon. It is not supported yet: checkpoint, restore and run are not available for containers that CRI-O manages.

## Status

| Feature | Supported |
|---|---|
| Checkpoint (dump) | No |
| Restore | No |
| Run | No |

`cedana features` lists dump and restore middleware and a run handler for this plugin. These are building blocks, not user-facing commands.

## What to use today

- On Kubernetes, install Cedana with the [Helm chart](https://github.com/cedana/cedana-helm-charts). It works through the container runtimes in the [support table](../../README.md#container-runtimes-and-vms).
- On one machine, use [runc](../../docs/guides/runc.md), [crun](../../docs/guides/crun.md), [containerd](../../docs/guides/containerd.md) or [podman](../../docs/guides/podman.md).

This README changes when CRI-O support ships. To ask about CRI-O support, open an issue in this repo.
