# B2Bi Configuration

This file documents the configuration variables used for deploying IBM Sterling B2B Integrator (B2Bi).

---

## 🧱 Core Deployment Settings

```yaml
b2bi_version: 6.2.1.0                      # Version of B2Bi to deploy
```

---

## 🗃️ DB2 Configuration if deploying DB2 alongside B2Bi

```yaml
deploy_db2: true                          # Deploy DB2 for B2Bi
db2_version: s11.5.9.0-cn3                # Version of DB2 to deploy
db2_namespace: b2bi-db2                   # Kubernetes namespace for DB2
db2_registry_secret: ibm-registry-secret  # Image pull secret for DB2
db2_secret: db2-secret                    # DB2 secret name containing credentials
db2_service_account: db2-sa               # DB2 service account
db2_crd_name: b2bi-db2                    # Name of the DB2 custom resource
db2_username: db2inst1                    # DB2 username
db2_password: passw0rd                    # DB2 password
db2_dbname: BLUDB                         # DB2 database name
db2_storage_size: 100Gi                   # Persistent storage size for DB2
```

---

## 💬 MQ Configuration if deploying MQ alongside B2Bi

```yaml
deploy_mq: true                           # Deploy IBM MQ
mq_version: v3.5.3                        # Version of IBM MQ to deploy
mq_namespace: b2bi-mq                     # Namespace for MQ deployment
mq_registry_secret: ibm-registry-secret   # Image pull secret for MQ
mq_secret: b2bi-mq-secret                 # MQ secret name
jms_user: app                             # JMS app user
mq_appuser_password: Passw0rd             # MQ app user password
mq_crd_name: qmgr                         # MQ custom resource name
mq_storage_size: 20Gi                     # Storage allocated for MQ
```

---

## 🚀 B2Bi Helm Deployment Settings

```yaml
b2bi_helm_release_name: auto              # Helm release name
b2bi_license_type: non-prod               # License type (prod or non-prod)
b2bi_namespace: b2bi                      # Namespace for B2Bi deployment
b2bi_serviceaccount: b2bi-sa              # Service account for B2Bi
```

---

## 🖼️ Registry Settings

```yaml
b2bi_registry_secret: ibm-registry-secret # Registry secret for pulling images
b2bi_registry: cp.icr.io                  # Container registry
b2bi_registry_user: cp                    # Registry username
b2bi_registry_key: ""                     # Registry key/token (if used)
```

---

## 🔐 JMS Settings

```yaml
jms_secret: b2bi-jms-queue
jms_keystore_path: ""
jms_keystore_secret: ""
jms_truststore_path: ""
jms_truststore_secret: ""
jms_protocol: TLSv1.2
jms_user: root
jms_vendor: IBMMQ
jms_connection_factory: com.ibm.mq.jms.MQQueueConnectionFactory
jms_queue_name: DEV.QUEUE.1
jms_channel: DEV.CHANNEL.SVRCONN
jms_enable_ssl: false
```

---

## 📦 Application Storage

```yaml
b2bi_app_logs_size: 1Gi
b2bi_app_documents_enabled: false
b2bi_app_documents_size: 1Gi
```

---

## ⚙️ Data Setup Options

```yaml
b2bi_data_setup_enabled: true
b2bi_data_setup_upgrade: false
```

---

## 🔌 Optional Integrations

```yaml
b2bi_seas_enabled: false
b2bi_itx_enabled: false
b2bi_itx_data_setup_enabled: false
b2bi_itxa_enabled: false
b2bi_itxa_data_setup_enabled: false
b2bi_itxa_app_secret: ""
b2bi_itxa_secure_dbconnection: false
b2bi_itxa_db_cert_secret: ""
b2bi_itxa_sso_host: ""
b2bi_itxa_sso_port: 443
b2bi_itxa_ssl_enabled: true
b2bi_itxa_resources_init_enabled: false
```

---

## 🔁 B2Bi Feature Toggles

```yaml
b2bi_sfg_enabled: false
b2bi_ebics_enabled: false
b2bi_financial_services_enabled: false
b2bi_file_operation_enabled: false
b2bi_fips_enabled: false
```

---

## 🏗️ Setup Configuration

```yaml
setup_cfg_upgrade: false
setup_cfg_base_port: 50000
b2bi_db_vendor: DB2
b2bi_db_drivers: db2jcc4.jar
b2bi_db_create_schema: true
b2bi_use_oracle_service_name: false
b2bi_cfg_use_ssl: false
b2bi_db_secret: b2bi-db-secret
b2bi_db_truststore: ""
b2bi_db_truststore_secret: ""
b2bi_db_keystore: ""
b2bi_db_keystore_secret: ""
b2bi_admin_email: admin@temp.net
b2bi_smtp_host: localhost
```

---

## 🔒 Secrets

```yaml
b2bi_system_passphrase_secret: b2bi-system-passphrase-secret
b2bi_system_passphrase: password
b2bi_liberty_keystore_password_secret: b2bi-liberty-secret
b2bi_liberty_keystore_password: changeit
```

---

## 🛡️ Security and NIST Compliance

```yaml
b2bi_nist_enabled: "off"  # Accepts "on" or "off"
```

---

## 🧩 Component Configuration

```yaml
asi_replica_count: 1
asi_backend_nodeport_range: 30351-30353
asi_autoscaling_enabled: false
asi_autoscaling_min_replicas: 1
asi_autoscaling_max_replicas: 2
asi_autoscaling_cpu_utilization_target_percent: 60
asi_fg_access_port:
asi_fg_protocol:

ac_replica_count: 1
ac_backend_nodeport_range: 30501-30600
ac_autoscaling_enabled: false
ac_autoscaling_min_replicas: 1
ac_autoscaling_max_replicas: 2
ac_autoscaling_cpu_utilization_target_percent: 60
ac_fg_access_port:
ac_fg_protocol:

api_replica_count: 1
api_autoscaling_enabled: false
api_autoscaling_min_replicas: 1
api_autoscaling_max_replicas: 2
api_autoscaling_cpu_utilization_target_percent: 60
```

---

## 🧹 Purge Job Configuration

```yaml
b2bi_purge_enabled: true
b2bi_purge_schedule: "0 0 * * *"
```

---

## 📄 Document Service

```yaml
b2bi_document_service_enabled: false
b2bi_document_ssl_enabled: true
b2bi_document_use_grpc: true
b2bi_document_replica_count: 1
```

---

## 🧪 Version-Specific Settings

```yaml
versions:
  6.2.1.0:
    helm_chart_version: 3.1.0
    ocp_min_version: 4.14
    k8s_min_version: 1.27
    k8s_max_version: 1.29
    resourcesInit_version: 6.2.1.0
    dataSetup_version: 6.2.1.0
    seas_version: 1.0
    itx_version: 10.1.2.1.20241122
    itxa_version: 10.0.1.10-x86_64
    itxa_resourcesInit_version: 10.0.1.10-x86_64
    test_version: 1.1.68
    purge_version: 6.2.1.0
    documentService_version: 6.2.1.0
```
