# Cedana Storage plugin

Stores checkpoints in Cedana Storage: global storage for checkpoints, backed by several cloud providers, with low latency and high availability. It is the fastest way to start with remote checkpoint and restore, because you only need a Cedana account.

## Requirements

- A Cedana account ([authentication](https://docs.cedana.ai/daemon/get-started/authentication)).
- The plugin: `sudo cedana plugin install storage/cedana`
- The daemon running. Check it with a [health check](https://docs.cedana.ai/daemon/get-started/health).

## Checkpoint and restore

```sh
cedana dump job my-job-1 --dir cedana://my-checkpoints
cedana restore job --attach my-job-1
```

The restore picks up the latest checkpoint for `my-job-1`. To restore an unmanaged checkpoint, pass `--path cedana://<path>/<dump>.tar`. To use Cedana Storage by default, set `Checkpoint.Dir` to a `cedana://` path in the configuration.

Use compression for large checkpoints and for GPU workloads (`lz4` is a good compromise). Streaming also works with this plugin. On Amazon EKS, [Amazon S3](../storage-s3/README.md) will likely be faster.

## Full guide

[Cedana Storage](../../docs/guides/storage/cedana.md)
