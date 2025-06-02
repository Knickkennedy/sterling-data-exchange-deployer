# IBM Sterling PEM STD Configuration

This file documents the configuration variables used for deploying IBM Sterling PEM STD.

---

```yaml
# -----------------------------------------------------------------------------
# IBM PEM Standard Configuration (WIP)
# -----------------------------------------------------------------------------
# This file documents the configuration variables used for deploying IBM
# Partner Engagement Manager (PEM) Standard Edition.
# 
# ⚠️ This configuration is a work in progress and has not been fully tested.
# -----------------------------------------------------------------------------

# --- Core Settings ---
pem_std_helm_release_name: auto                             # Helm release name
pem_std_version: 6.2.4.0.1_standard                          # PEM version
pem_std_license_type: nonprod                               # License type
pem_std_namespace: pem-std                                  # Kubernetes namespace
pem_std_serviceaccount: pem-std-sa                          # Service account for PEM pods

# --- Registry Configuration ---
pem_std_registry_secret: ibm-registry-secret                # Secret for image registry access
pem_std_registry: cp.icr.io                                 # Docker image registry URL
pem_std_registry_user: cp                                   # Registry username
pem_std_registry_key:                                       # Registry access key (to be filled)

# --- Passwords & Credentials ---
pem_std_default_password: Passw0rd123!                      # Default PEM password
pem_std_default_passphrase: P@ssPhr4s3_890@                 # PEM system passphrase
pem_std_keystore_password: Change1t@                        # Java keystore password

# --- Admin & SMTP Settings ---
pem_std_admin_email: admin@company.com                      # PEM admin email
pem_std_smtp_host: smtp.company.com                         # SMTP mail server
pem_std_smtp_port: 25                                       # SMTP server port

# --- JMS Settings (Optional) ---
pem_std_jms_username: ""
pem_std_jms_password: ""
pem_std_jms_keystore_password: ""
pem_std_jms_truststore_password: ""

# --- Database Setup Flags ---
pem_std_dbsetup_enabled: true                               # Enable database setup on deploy
pem_std_dbsetup_upgrade: false                              # Allow DB upgrade during setup
pem_std_dbdriver: com.ibm.db2.jcc.DB2Driver                 # JDBC driver class name

# --- Secrets ---
pem_std_secret: ibm-pem-secret                              # Main PEM secret
pem_std_db_secret: ibm-pem-db-secret                        # Database credentials secret
pem_std_cm_prod_secret: ibm-pem-cm-prod-secret              # CM prod environment secret
pem_std_cm_nonprod_secret: ibm-pem-cm-nonprod-secret        # CM non-prod environment secret
pem_std_keyalias: pem_default                               # Default keystore key alias
pem_std_cert_jks_secret: pem-keystore.jks                   # JKS certificate secret
pem_std_cert_p12_secret: pem-keystore.p12                   # P12 certificate secret
pem_std_cert_crt_secret: pem-cert.crt                       # PEM certificate secret

# --- B2Bi Integration (Prod) ---
pem_std_b2bi_prod_url: ""
pem_std_b2bi_prod_user: ""
pem_std_b2bi_prod_pw: ""
pem_std_b2bi_prod_passphrase: ""
pem_std_b2bi_prod_apiurl: ""
pem_std_b2bi_prod_apiuser: ""
pem_std_b2bi_prod_apipw: ""
pem_std_b2bi_prod_dburl: ""
pem_std_b2bi_prod_dbuser: ""
pem_std_b2bi_prod_dbpw: ""
pem_std_b2bi_prod_jpa_dialect: com.pe.pcm.config.database.dialect.DB2ExtendedDialect

# --- B2Bi Integration (Non-prod) ---
pem_std_b2bi_nonprod_url: ""
pem_std_b2bi_nonprod_user: ""
pem_std_b2bi_nonprod_pw: ""
pem_std_b2bi_nonprod_passphrase: ""
pem_std_b2bi_nonprod_apiurl: ""
pem_std_b2bi_nonprod_apiuser: ""
pem_std_b2bi_nonprod_apipw: ""
pem_std_b2bi_nonprod_dburl: ""
pem_std_b2bi_nonprod_dbuser: ""
pem_std_b2bi_nonprod_dbpw: ""
pem_std_b2bi_nonprod_jpa_dialect: com.pe.pcm.config.database.dialect.DB2ExtendedDialect

# --- JMS Configuration ---
jms_user: root                                              # JMS admin user
jms_vendor: IBMMQ                                           # JMS vendor
jms_connection_factory: com.ibm.mq.jms.MQQueueConnectionFactory
jms_queue_name: DEV.QUEUE.1                                 # JMS queue name
jms_channel: DEV.CHANNEL.SVRCONN                            # JMS channel
jms_enable_ssl: false                                       # Enable JMS over SSL

# --- Storage & Logging ---
pem_std_app_logs_size: 1Gi                                  # Application log volume size
pem_std_app_documents_enabled: false                        # Enable document service
pem_std_app_documents_size: 1Gi                             # Volume size for documents

# --- Data Setup Flags ---
pem_std_data_setup_enabled: true                            # Enable data setup at deploy
pem_std_data_setup_upgrade: false                           # Enable data upgrade

# --- Additional Secrets ---
pem_std_system_passphrase_secret: pem-std-system-passphrase-secret
pem_std_system_passphrase: password                         # System-level PEM passphrase

pem_std_liberty_keystore_password_secret: pem-std-liberty-secret
pem_std_liberty_keystore_password: changeit                # Liberty keystore password

# --- Security & Compliance ---
pem_std_nist_enabled: "off"                                 # Enable/disable NIST compliance

# --- Network ---
asi_backend_nodeport_range: 30351-30353                     # NodePort range for backend services

# --- Version Details ---
versions:
  6.2.4.0.1_standard:
    helm_chart_version: 1.4.2
    resourcesInit_version: 6.2.1.0
    dataSetup_version: 6.2.1.0
    seas_version: 1.0
    itx_version: 10.1.2.1.20241122
    itxa_version: 10.0.1.10-x86_64
    itxa_resourcesInit_version: 10.0.1.10-x86_64
    test_version: 1.1.67
    purge_version: 6.2.1.0
    documentService_version: 6.2.1.0

# --- DB2 Tuning Parameters ---
db_values:
  STMT_CONC: "LITERALS"
  SELF_TUNING_MEM: "ON"
  LOCKLIST: "AUTOMATIC"
  MAXLOCKS: "AUTOMATIC"
  PCKCACHESZ: "AUTOMATIC"
  SHEAPTHRES_SHR: "AUTOMATIC"
  SORTHEAP: "AUTOMATIC"
  NUM_IOCLEANERS: "AUTOMATIC"
  NUM_IOSERVERS: "AUTOMATIC"
  DFT_PREFETCH_SZ: "AUTOMATIC"
  MAXAPPLS: "AUTOMATIC"
  APPLHEAPSZ: "AUTOMATIC"
  APPL_MEMORY: "AUTOMATIC"
  DBHEAP: "AUTOMATIC"
  LOGFILSIZ: "65536"
  LOGPRIMARY: "40"
  LOGSECOND: "12"
  NUM_LOG_SPAN: "40"
  DFT_DEGREE: "1"
  LOGBUFSZ: "2152"
  STMTHEAP: "51200 AUTOMATIC"
  CATALOGCACHE_SZ: "300"
  UTIL_HEAP_SZ: "AUTOMATIC"
  CHNGPGS_THRESH: "80"
```
---