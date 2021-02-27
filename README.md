# auto-raise-hpa

This is a tool for raise MaxReplicas of hpa if Replicas of deployment scaled equal to MaxReplicas.

## Prerequisites

- [docker](https://www.docker.com/products/docker-desktop)
- [skaffold](https://skaffold.dev/docs/install/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)

## Usage

switch to kubernetes context which you want to clean evicted pod

```bash
kubectl config use-context <context> 
```

build image and deploy to kubernetes cluster

```bash
skaffold dev
```
