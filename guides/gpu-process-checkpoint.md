# Checkpoint a GPU process, and what Cedana adds to CRIU

This guide checkpoints a running CUDA process with the Cedana daemon and restores it on the same machine. On its own, CRIU does not checkpoint a process that holds a GPU. The Cedana daemon adds the GPU side with one of two plugins: Cedana's GPU plugin, or the CRIU community's CUDA plugin, which calls NVIDIA's cuda-checkpoint.

## What you need first

- Linux with an NVIDIA GPU. The Cedana GPU plugin supports drivers 452 to 570 on amd64 and arm64. The CRIU CUDA plugin supports driver 570 on amd64.
- The Cedana daemon installed and running, as in the [quick start](../README.md#quick-start).
- A Cedana account, because this guide installs either GPU plugin from Cedana's plugin registry. The CRIU CUDA plugin itself is open source. Reach out to us through your contact at Cedana or founders@cedana.ai to get an account!

## Steps

1. Install the Cedana GPU plugin:

   ```sh
   sudo cedana plugin install gpu
   ```

   To use the CRIU CUDA plugin instead, run `sudo cedana plugin install criu/cuda`. It handles one process at a time.

2. Check that the daemon and the plugin are ready:

   ```sh
   cedana daemon check --full
   ```

3. Get a small CUDA program to run. The samples repo has one:

   ```sh
   git clone https://github.com/cedana/cedana-samples
   ```

4. Run it as a Cedana job with GPU support. `--attach` keeps this terminal connected to the job, so open a second terminal for the next steps:

   ```sh
   cedana run process --attach --gpu-enabled --jid gpu-demo -- cedana-samples/gpu_smr/vector_add
   ```

5. In the second terminal, checkpoint the job:

   ```sh
   cedana dump job gpu-demo
   ```

6. Restore it. The terminal you run this in is attached to the restored job:

   ```sh
   cedana restore job --attach gpu-demo
   ```

For a job that coordinates its GPUs through NCCL, checkpoint with `cedana dump job <job_id> --gpu-freeze-type nccl`. The CRIU CUDA plugin does not support jobs with several processes.

## Check the result

`cedana ps` lists the job with a checkpoint time and size, and after the restore the program carries on from where it was checkpointed.

## Limits today

- A checkpoint restores only on a node with the same GPU, driver, engine and model versions it recorded. If any of them changes, the checkpoint is invalid and the workload cold-starts instead.
- NVIDIA GPUs only.
- The daemon needs root.

## Read more

[What CRIU and cuda-checkpoint do when you wire them together yourself](https://cedana.com/blog/criu-gpu-checkpointing/), and the [GPU guide](../docs/guides/gpu.md) for every option.
