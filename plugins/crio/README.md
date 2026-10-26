# Cedana CRI-O plugin

This plugin adds CRI-O support to the Cedana daemon. CRI-O is partially supported: the plugin's dump, restore and run middleware are implemented, but its dump, restore and run commands are not. You cannot yet checkpoint, restore or run a container that CRI-O manages with the Cedana CLI.

## Status

From [`cedana features`](../../docs/get-started/features.md).

| Feature | Implemented |
|---|---|
| Dump, restore and run middleware | Yes |
| Dump command | No |
| Restore command | No |
| Run command | No |

## What to use today

- On Kubernetes, install Cedana with the [Helm chart](https://github.com/cedana/cedana-helm-charts). It works through the container runtimes in the [support table](../../README.md#container-runtimes-and-vms).
- On one machine, use [runc](../../docs/guides/runc.md), [crun](../../docs/guides/crun.md), [containerd](../../docs/guides/containerd.md) or [podman](../../docs/guides/podman.md).

To ask about CRI-O support, open an issue in this repo.
