# IBM DB2 Configuration

This file documents the configuration variables used for deploying IBM DB2 with the operator.

---

##    Core Settings 

```yaml
db2_version: s11.5.9.0-cn3                      # Version of DB2 container image to use
db2_channel: v110509.0                          # DB2 update channel or tag, if required
db2_namespace: db2                              # Kubernetes namespace for the DB2 deployment
```
---

##    Registry Settings 

```yaml
registry: cp.icr.io                             # IBM container registry URL
registry_user: cp                               # IBM container registry username
db2_registry_secret: ibm-registry-secret        # Kubernetes secret to authenticate with IBM registry
```
---

##    Secrets and Access 

```yaml
db2_secret: db2-secret                          # Secret containing DB2 credentials
db2_service_account: db2-sa                     # Service account used by the DB2 pods
```
---

##    Database Configuration 

```yaml
db2_crd_name: db2                               # Custom Resource Definition (CRD) name for the DB2 instance
db2_username: db2inst1                          # Username for accessing the DB2 database
db2_password: passw0rd                          # Password for the DB2 database user
db2_dbname: BLUDB                               # Name of the DB2 database
db2_storage_size: 100Gi                         # Size of the persistent volume for DB2 storage
```
---