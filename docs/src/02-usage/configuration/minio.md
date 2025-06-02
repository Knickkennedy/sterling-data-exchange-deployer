# MinIO Configuration

This file documents the configuration variables used for deploying MinIO.

---
##    Core Settings 
```yaml
minio_namespace: minio                      # Namespace where MinIO will be deployed
minio_serviceaccount: minio                 # Service account for the MinIO pods
```
---
##    Credentials 
```yaml
minio_root_user: minio                      # Root username for accessing the MinIO UI and API
minio_root_password: minio123               # Root password for MinIO
```
---
##    Storage 
```yaml
minio_size: 250Gi                           # Persistent storage size allocated to MinIO
```
---