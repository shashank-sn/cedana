# Cedana Google Cloud Storage plugin

This plugin will store checkpoints in your own Google Cloud Storage buckets. Checkpoint and restore to Google Cloud Storage is work in progress, so do not use this plugin for checkpoints yet.

## What to use today

| Need | Plugin |
|---|---|
| Checkpoints in your own cloud bucket | [storage/s3](../storage-s3/README.md), which also works with S3-compatible endpoints through `AWS.Endpoint` |
| The fastest start, with only a Cedana account | [storage/cedana](../storage-cedana/README.md) |
| Checkpoints on your own NFS-backed storage | [storage/csx](../storage-csx/README.md) |

On a GKE cluster, Google Cloud Storage will likely be faster than Cedana Storage once this plugin ships.

## More

- [Google Cloud Storage guide](../../docs/guides/storage/gcs.md)
- [`cedana features`](../../docs/get-started/features.md) shows checkpoint storage as not implemented for this plugin today.

This README changes when Google Cloud Storage support ships.
