# IBM Sterling Secure Proxy Engine Configuration

This file documents the configuration variables used for deploying IBM Sterling Secure Proxy Engine.

---

```yaml
# -----------------------------------------------------------------------------
# IBM SSP Engine Configuration
# -----------------------------------------------------------------------------
# This file documents the configuration variables used for deploying IBM SSP Engine.
# -----------------------------------------------------------------------------

# --- Core Settings ---
ssp_engine_release_name: auto-eng                        # Helm release name for SSP Engine deployment
ssp_engine_version: 6.2.0.1.02                           # SSP Engine version
ssp_engine_license_type: non-prod                         # License type: prod, non-prod, etc.

ssp_engine_namespace: ssp                                # Kubernetes namespace for SSP Engine
ssp_cm_namespace: ssp                                    # Kubernetes namespace for SSP CM (related)

ssp_engine_generate_certificates: false                  # Enable/disable automatic cert generation
ssp_engine_serviceaccount: ssp-engine-sa                 # Service account for SSP Engine
ssp_engine_pvc: ibm-ssp-engine-pvc                       # Persistent volume claim for SSP Engine

# --- Registry Configuration ---
ssp_engine_registry_secret: ibm-entitlement-key          # Secret to pull container images
ssp_engine_registry: cp.icr.io                            # Image registry URL
ssp_engine_registry_user: cp                              # Registry username
ssp_engine_registry_key:                                  # Optional registry key, leave empty if not used

# --- Passwords & Secrets ---
ssp_engine_secret: ibm-ssp-engine-secret                 # Kubernetes secret for SSP Engine sensitive data

ssp_engine_sys_passphrase: Passw0rd@123                   # System passphrase

ssp_engine_keycert_store_passphrase: password             # Keycert store passphrase
ssp_engine_keycert_encrypt_passphrase: password           # Keycert encryption passphrase
ssp_engine_keycert_secret: "cm-keycert"                   # Kubernetes secret containing keycert data
ssp_engine_custom_keycert_passphrase: password            # Custom keycert passphrase

# --- Version Info ---
versions:
  6.2.0.1.02:
    helm_chart_version: 1.5.4                              # Helm chart version for this SSP Engine version
```
---