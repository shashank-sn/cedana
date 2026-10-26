# Cedana S3 storage plugin

Stores checkpoints in your own Amazon S3 buckets, and restores from them. General-purpose and directory buckets both work. Custom S3-compatible endpoints work through `AWS.Endpoint`.

## Requirements

- A Cedana account, to install the plugin from the registry ([authentication](https://docs.cedana.ai/daemon/get-started/authentication)).
- The plugin: `sudo cedana plugin install storage/s3`
- AWS credentials, in one of three modes set by `AWS.CredentialsMode`:
  - `static` (the default): set `AWS.AccessKeyID` and `AWS.SecretAccessKey`, or `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY`.
  - `eksPodIdentity`: EKS Pod Identity. The Pod Identity Agent injects the credentials into the Cedana helper pods.
  - `ambient`: the AWS default credential chain, for example IRSA, EC2 instance roles or shared AWS configuration.

## Checkpoint and restore

```sh
cedana dump job my-job-1 --dir s3://checkpoints-bucket
cedana restore job --attach my-job-1
```

The restore picks up the latest checkpoint for `my-job-1`. To restore an unmanaged checkpoint, pass `--path s3://<bucket>/<path>/<dump>.tar`. To use S3 by default, set `Checkpoint.Dir` to an `s3://` path in the configuration.

For large checkpoints, and for GPU workloads, use compression (`lz4` is a good compromise). Streaming also works with this plugin.

## Full guide

[Amazon S3](../../docs/guides/storage/s3.md)
