# Cedana Storage Express (CSX) plugin

Stores checkpoints through CSX, Cedana's NFS-backed storage. CSX uses tmpfs, local disk and NFS to checkpoint and restore quickly. The plugin talks to a local CSX daemon over a Unix socket.

## Requirements

- A Cedana account, to install released plugins ([authentication](https://docs.cedana.ai/daemon/get-started/authentication)).
- The plugin: `sudo cedana plugin install storage/csx`
- The CSX binary: `sudo cedana plugin install csx`
- The Cedana daemon running, then the CSX daemon started with `sudo ./csx`. It needs root because it creates `tmpfs` mounts, and it must keep running while you checkpoint or restore. Its socket is set by `CSX.SockAddr`.

## Checkpoint and restore

```sh
cedana dump job test --dir csx://
cedana restore job test
```

Cedana records the checkpoint path with the `csx://` prefix, so a managed job restores without a path. To restore an unmanaged checkpoint, pass `--path csx://<dump-path>`. To use CSX by default, set `CEDANA_CHECKPOINT_DIR=csx://` or `Checkpoint.Dir` to `csx://`, then restart the daemon.

Compression works. Streaming does not work with CSX yet, so leave `--streams` unset or at `0`.

## Full guide

[Cedana Storage Express](../../docs/guides/storage/csx.md)
