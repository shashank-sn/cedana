# Cedana podman plugin

Adds podman support to the Cedana daemon: checkpoint and restore containers that podman runs, and run or manage podman containers as Cedana jobs, with or without GPUs. Only rootful podman is supported.

## Requirements

- A Cedana account, to install the plugins from the registry ([authentication](https://docs.cedana.ai/daemon/get-started/authentication)).
- The podman plugin: `sudo cedana plugin install podman`
- The oci-runtime plugin, the OCI runtime podman uses for containers that Cedana runs or restores: `sudo cedana plugin install oci-runtime`
- The plugin for podman's OCI runtime: `sudo cedana plugin install crun` (the default) or `sudo cedana plugin install runc`
- The daemon running. Check it with a [health check](https://docs.cedana.ai/daemon/get-started/health).

## Checkpoint and restore a podman container

```sh
sudo podman run -d --name <container> docker.io/library/nginx:latest
cedana dump podman <container> --dir <dump-dir>
sudo podman rm <container>
cedana restore podman --path <dump-dir>/<dump> --id <container>
```

The restored container is a regular podman container, so `podman ps`, `logs`, `exec` and `stop` work on it as usual.

## Run a container as a managed job

`--attach` keeps the terminal connected to the job, so run `cedana dump` and `cedana restore` in a second terminal.

```sh
cedana run podman --attach --jid <job_id> docker.io/library/nginx:latest
cedana dump job <job_id>
cedana restore job --attach <job_id>
```

With GPUs, add `--gpu-enabled --gpus <gpus>` to `cedana run podman`. The restored container gets the same GPUs. GPU restore needs podman's `overlay` storage driver (the default).

## Full guide

[Checkpoint and restore with podman](../../docs/guides/podman.md)
