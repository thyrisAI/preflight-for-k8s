# Preflight for Thyris Runtime on Kubernetes

`preflight-for-k8s` is a Thyris preflight check tool for Kubernetes clusters. It verifies several prerequisites and configurations before running AI workloads on your cluster. The script checks storage, CPU, memory, and Longhorn configurations to ensure the environment is ready for intensive workloads.

## Pre-requisites

Before running the preflight script, ensure the following requirements are met:

1. **Kubernetes Cluster**: You must have a running Kubernetes cluster (version 1.29+ recommended).
2. **kubectl**: Ensure `kubectl` is installed and configured to interact with your Kubernetes cluster.
4. **Kubernetes Permissions**: The Kubernetes cluster should have proper RBAC roles to allow the preflight script to check necessary resources (e.g., storage, CPU, RAM).

### Longhorn Installation (Optional)
If you're using Longhorn for storage, make sure Longhorn is properly installed and running on your cluster.

```bash
kubectl apply -f https://raw.githubusercontent.com/longhorn/longhorn/master/deploy/longhorn.yaml
```
### Apply deploy.yaml (Optional)

```bash
kubectl apply -f https://raw.githubusercontent.com/thyrisAI/preflight-for-k8s/refs/heads/main/deploy/deploy.yaml
```




