# 🧩 Kubernetes Resource Overview: `sde-deployer.yaml`

This document outlines the Kubernetes artifacts deployed by the provided YAML file. These resources are used to configure and launch a deployment automation job for a suite of enterprise IBM applications.

---

## 📦 1. **Namespace**
```yaml
kind: Namespace
```
Creates a dedicated namespace `sde-deployer` to logically isolate and manage the associated resources.

---

## 🔐 2. **Secret**
```yaml
kind: Secret
```
`entitlement-key-secret` stores the IBM Container Registry entitlement key needed to authenticate when pulling protected container images.

---

## ⚙️ 3. **ConfigMap**
```yaml
kind: ConfigMap
```
`sde-config` provides a collection of application-specific configuration YAML files (e.g., for B2Bi, ITXA, Connect:Direct). This is mounted into the deployer container for dynamic use at runtime.

---

## 👤 4. **ServiceAccount**
```yaml
kind: ServiceAccount
```
`sde-deployer-sa` defines an identity used by the deployment job to authenticate within the cluster.

---

## 🔐 5. **RoleBinding**
```yaml
kind: RoleBinding
```
Grants the `sde-deployer-sa` service account permission to use the `privileged` Security Context Constraint (SCC), allowing more permissive operations (e.g., running as root or accessing host-level resources if required).

---

## 🌐 6. **ClusterRoleBinding**
```yaml
kind: ClusterRoleBinding
```
Gives the `sde-deployer-sa` service account `cluster-admin` permissions, enabling full access across the cluster — typically necessary for complex, multi-namespace deployments.

---

## ⚙️ 7. **Job**
```yaml
kind: Job
```
`sde-deployer-job` runs a one-time container using the image `knickkennedy/sde-deployer:v6.2.1.0`. It mounts the configuration files and uses the entitlement key to deploy IBM applications (e.g., B2Bi, ITXA) across the cluster according to the contents of the `sde-config` ConfigMap.

Key features:
- Uses the service account `sde-deployer-sa`.
- Mounts the config files at `/app/config`.
- Pulls secrets and values dynamically at runtime.

---

## ✅ Summary

| Resource           | Purpose                                                       |
|--------------------|---------------------------------------------------------------|
| Namespace          | Isolates the deployment artifacts                             |
| Secret             | Stores registry credentials securely                          |
| ConfigMap          | Supplies configuration for application components             |
| ServiceAccount     | Provides an identity for the deployer job                     |
| RoleBinding        | Grants SCC privileges for enhanced pod capabilities           |
| ClusterRoleBinding | Grants full cluster access for deployment                     |
| Job                | Executes the actual deployment logic                          |

