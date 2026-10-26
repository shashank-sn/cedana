# Cedana Google Cloud Storage plugin

Checkpoint and restore to Google Cloud Storage is not implemented today, so do not use this plugin for checkpoints.

## What to use today

| Need | Plugin |
|---|---|
| Checkpoints in your own cloud bucket | [storage/s3](../storage-s3/README.md), which also works with S3-compatible endpoints through `AWS.Endpoint` |
| The fastest start, with only a Cedana account | [storage/cedana](../storage-cedana/README.md) |
| Cedana's NFS-backed storage (CSX) | [storage/csx](../storage-csx/README.md) |

## More

- [Google Cloud Storage guide](../../docs/guides/storage/gcs.md)
- [`cedana features`](../../docs/get-started/features.md) shows checkpoint storage as not implemented for this plugin today.
