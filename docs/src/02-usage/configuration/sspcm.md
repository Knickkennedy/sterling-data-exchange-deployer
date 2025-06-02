# IBM Sterling Secure Proxy Configuration Manager Configuration

This file documents the configuration variables used for deploying IBM Sterling Secure Proxy Configuration Manager.

---

##    Core Settings 

```yaml
ssp_cm_release_name: auto                                # Helm release name for SSP CM deployment
ssp_cm_version: 6.2.0.1.02                              # SSP CM version
ssp_cm_license_type: non-prod                            # License type: prod, non-prod, etc.

ssp_cm_namespace: ssp                                   # Kubernetes namespace for SSP CM
ssp_engine_namespace: ssp                               # Kubernetes namespace for SSP engine

ssp_cm_generate_certificates: true                      # Enable generation of certificates
ssp_cm_serviceaccount: ssp-cm-sa                        # Service account for SSP CM
ssp_cm_pvc: ibm-ssp-cm-pvc                              # Persistent volume claim for SSP CM
```
---

##    Registry Configuration 

```yaml
ssp_cm_registry_secret: ibm-entitlement-key            # Secret to pull container images
ssp_cm_registry: cp.icr.io                              # Image registry URL
ssp_cm_registry_user: cp                                # Registry username
ssp_cm_registry_key:                                    # Optional registry key, leave empty if not used
```
---

##    Passwords & Secrets 

```yaml
ssp_cm_secret: ibm-ssp-cm-secret                        # Kubernetes secret for SSP CM sensitive data

ssp_cm_sys_passphrase: Passw0rd@123                     # System passphrase
ssp_cm_admin_passphrase: Passw0rd@123                   # Admin passphrase

ssp_cm_keycert_store_passphrase: password               # Keycert store passphrase
ssp_cm_keycert_encrypt_passphrase: password             # Keycert encryption passphrase
ssp_cm_common_cert_passphrase: password                  # Common certificate passphrase
ssp_cm_eng_cert_password: password                       # Engine certificate password
ssp_cm_client_cert_password: password                    # Client certificate password
ssp_cm_cert_password: password                           # General certificate password
ssp_cm_server_cert_password: password                    # Server certificate password
ssp_cm_web_cert_password: password                       # Web certificate password
ssp_cm_export_cert_password: password                    # Export certificate password

ssp_cm_keycert_secret: "engine-keycert"                 # Kubernetes secret containing keycert data
```
---

##    Version Info 

```yaml
versions:
  6.2.0.1.02:
    helm_chart_version: 1.5.4                            # Helm chart version for this SSP CM version
```
---