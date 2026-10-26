# Cedana containerd plugin

Adds containerd support to the Cedana daemon: checkpoint containers that containerd runs, and run and manage new containerd containers as Cedana jobs, with or without GPUs.

## Requirements

- A Cedana account, to install the plugin from the registry: `sudo cedana plugin install containerd`
- The plugin for the OCI runtime underneath, and the Cedana runtime (shim) plugin for the runtime type:
  - runtime type `io.containerd.runc.v2`: `sudo cedana plugin install containerd/runtime-runc`
  - runtime binary `runc` (the default): `sudo cedana plugin install runc`
  - runtime binary `crun`: `sudo cedana plugin install crun`
- The daemon running. Check it with a [health check](https://docs.cedana.ai/daemon/get-started/health).

## Checkpoint a running container

```sh
sudo ctr run docker.io/library/nginx:latest <container_id>
cedana dump containerd <container_id> --dir <dump-dir>
```

To restore it, delete the old container, then restore from the dump path that `cedana dump` prints:

```sh
sudo ctr container delete <container_id>
cedana restore containerd --path <dump-path> --id <container_id>
```

## Run a container as a managed job

`--attach` keeps the terminal connected to the job, so run `cedana dump` and `cedana restore` in a second terminal.

A CPU example, with nginx:

```sh
cedana run containerd --attach --jid <job_id> docker.io/library/nginx:latest
cedana dump job <job_id>
cedana restore job --attach <job_id>
```

A GPU example:

```sh
cedana run containerd --attach --gpu-enabled --jid <job_id> <your-gpu-image>
cedana dump job <job_id>
cedana restore job --attach <job_id>
```

Cedana's regression tests cover containerd checkpoint and restore, including with crun as the runtime.

## Full guide

[Checkpoint and restore with containerd](../../docs/guides/containerd.md)
