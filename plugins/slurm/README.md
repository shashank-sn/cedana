# Cedana Slurm plugin

Adds Slurm support to the Cedana daemon. Slurm jobs can be checkpointed and restored, and a requeued job resumes as of its last checkpoint instead of starting its batch script again. Users keep using `sbatch`, `srun`, `squeue` and `scancel` with no change to arguments, queues or usage. Admins run one setup script per cluster.

## What the plugin installs

| File | Where | Role |
|---|---|---|
| `cedana-slurm` | binary directory | helper binary |
| `task_cedana.so` | Slurm plugin directory | task plugin, added to `TaskPlugin` |
| `spank_cedana.so` | Slurm plugin directory | SPANK plugin, listed in `plugstack.conf` |
| `cli_filter_cedana.so` | Slurm plugin directory | added to `CliFilterPlugins` |
| `job_submit_cedana.so` | Slurm plugin directory | added to `JobSubmitPlugins` |

They go on the worker nodes (slurmd), the controller nodes (slurmctld) and the database node (slurmdbd).

## Install

Recommended: the web installer. It installs the latest stable version and every plugin Slurm support needs. Run the same setup script on every node of the cluster, each time with that node's role (controller, worker or login). Installing Cedana on Slurm does not require a Slurm restart.

The installer needs a Cedana account, for the org URL, auth token and cluster ID below. Reach out to us through your contact at Cedana or founders@cedana.ai to get an account!

```sh
export CEDANA_URL=<your-org-url>
export CEDANA_AUTH_TOKEN=<your-auth-token>
export CEDANA_CLUSTER_ID=<cluster_id>
curl -fsSL "${CEDANA_URL}/install/slurm" -H "Authorization: Bearer ${CEDANA_AUTH_TOKEN}" | sudo -E bash -s -- --node-role <controller|worker|login>
```

To install the plugin files by hand instead, follow the manual steps in the [installation docs](https://docs.cedana.ai/cedana-slurm/installation).

No custom Slurm build is needed. The minimum is a controller and one compute node. Nodes without Cedana keep working, without checkpoint and restore.

## Use the plugin in a job

```sh
#SBATCH --export=CEDANA_ENABLE=1,CEDANA_CHECKPOINT=my-job
```

`CEDANA_ENABLE=1` turns checkpointing on for the job. `CEDANA_CHECKPOINT` labels its checkpoints, so a restore picks up the latest checkpoint with that label.

## What the plugin supports

From [`cedana features`](../../docs/get-started/features.md): dump and restore commands, helper commands, dump and restore middleware, freeze and unfreeze handlers, and serverless mode. For GPU jobs, see [checkpoint and restore on Slurm](https://docs.cedana.ai/cedana-slurm/cr). For sites where root is restricted, see [unprivileged mode](https://docs.cedana.ai/cedana-slurm/unprivileged-mode-non-root).

## Remove the plugin

Run on all nodes:

```sh
sudo cedana slurm destroy
```

## More

- [Slurm installation docs](https://docs.cedana.ai/cedana-slurm/installation)
- [Cedana installation: what changes on your cluster](https://cedana.com/blog/cedana-installation-cluster-changes/)
- [Slurm time limits: what to do after a job stops](https://cedana.com/blog/slurm-time-limit/)
