# Keep a training job's progress through a spot reclaim

This guide sets up a Kubernetes training Job so that, when its spot node is reclaimed, the replacement pod restores from the latest checkpoint instead of starting from the beginning. Cedana checkpoints the job on an interval, below the application, so the training code does not change.

## What you need first

- A Kubernetes cluster with Cedana installed from the [Helm chart](https://github.com/cedana/cedana-helm-charts). The install needs a Cedana account. Reach out to us through your contact at Cedana or founders@cedana.ai to get an account!
- A spot node pool that replaces a reclaimed node with one of the same GPU type, for example with Karpenter.

## Steps

1. Add one environment variable to the training container. Its value labels the job's checkpoints. GPU pods also set `runtimeClassName: cedana`:

   ```yaml
   apiVersion: batch/v1
   kind: Job
   metadata:
     name: train
   spec:
     backoffLimit: 3
     template:
       spec:
         restartPolicy: Never
         runtimeClassName: cedana
         containers:
           - name: train
             image: <your-training-image>
             env:
               - name: CEDANA_CHECKPOINT
                 value: train-run-1
             resources:
               limits:
                 nvidia.com/gpu: 1
   ```

2. Apply it:

   ```sh
   kubectl apply -f train-job.yaml
   ```

3. In the Cedana dashboard, turn on a heartbeat checkpoint policy for the job, so Cedana checkpoints it every few minutes. Pick the interval from how much work you can afford to run again.

4. When the spot node is reclaimed, the Job controller creates a new pod, and Cedana restores it from the latest checkpoint labelled `train-run-1`.

## Check the result

Before a real reclaim, test it: wait for a checkpoint, then delete the pod with `kubectl delete pod <pod-name>`. The new pod's log carries on from the step it had reached at the checkpoint, instead of from step 0.

## Limits today

- A reclaim loses the work done since the last checkpoint. Checkpoints run on an interval the operator sets.
- A checkpoint restores only where versions match: the new node needs the same GPU, driver, engine and model versions, or the job cold-starts instead.
- Effects outside the machine do not roll back. A metric logged or a file uploaded after the checkpoint is sent again after the restore.

## Read more

[What happens to a training job when a spot instance is reclaimed](https://cedana.com/blog/spot-instance-interruption-training/), and [jobs that resume after a restart](https://docs.cedana.ai/cedana-kubernetes/jobs).
