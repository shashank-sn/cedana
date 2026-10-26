# Checkpoint a running container and restore it on another node

This guide checkpoints a containerd container on one node, copies the checkpoint to a second node, and restores the container there. A migration is a restore on a different node, so the steps are a checkpoint, a copy and a restore.

## What you need first

- Two Linux nodes with containerd, each with the Cedana daemon installed and running, as in the [quick start](../README.md#quick-start).
- A Cedana account, to install the plugins from Cedana's plugin registry. Reach out to us through your contact at Cedana or founders@cedana.ai to get an account!
- On both nodes, the containerd plugin and the plugins for the runtime underneath. For the default runc runtime:

  ```sh
  sudo cedana plugin install containerd containerd/runtime-runc runc
  ```

- A way to copy files from the first node to the second, such as SSH.

## Steps

1. On the first node, run a container:

   ```sh
   sudo ctr run docker.io/library/nginx:latest web-1
   ```

2. In a second terminal on the first node, checkpoint it. `cedana dump` prints the checkpoint's path:

   ```sh
   cedana dump containerd web-1 --dir <dump-dir>
   ```

3. Copy the checkpoint to the same path on the second node:

   ```sh
   rsync -a <dump-path> <second-node>:<dump-path>
   ```

4. On the second node, make sure the same image is there, then restore the container from the checkpoint:

   ```sh
   sudo ctr image pull docker.io/library/nginx:latest
   cedana restore containerd --path <dump-path> --id web-1
   ```

5. Once the restored container is serving, delete the original on the first node:

   ```sh
   sudo ctr container delete web-1
   ```

## Check the result

`sudo ctr task ls` on the second node shows `web-1` running, and the container's files and memory are as they were at the checkpoint.

## Limits today

- A checkpoint restores only where versions match. For a GPU container, the second node needs the same GPU, driver, engine and model versions the checkpoint recorded, or the workload cold-starts instead.
- The container's writable layer comes back only with containerd's default `overlayfs` snapshotter.
- Effects outside the machine do not roll back. A request that went out after the checkpoint runs again after the restore.

## Read more

To keep checkpoints in a bucket instead of copying them, see the [S3 storage plugin](../plugins/storage-s3/README.md).

[Does Kubernetes, Docker, Slurm or Nextflow already restore a running job?](https://cedana.com/blog/container-checkpoint-restore/), and the [containerd plugin README](../plugins/containerd/README.md).
