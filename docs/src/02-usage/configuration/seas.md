# IBM SEAS Configuration

This file documents the configuration variables used for deploying IBM SEAS.

---

```yaml
# -----------------------------------------------------------------------------
# IBM SEAS Configuration
# -----------------------------------------------------------------------------
# This file documents the configuration variables used for deploying IBM SEAS.
# -----------------------------------------------------------------------------

## --- Core Settings ---
seas_release_name: auto-seas                              # Helm release name for SEAS deployment
seas_version: 6.1.0.2.03                                  # SEAS version
seas_license_type: non-prod                               # License type: prod, non-prod, etc.

seas_namespace: seas                                     # Kubernetes namespace for SEAS
seas_serviceaccount: seas-sa                             # Service account for SEAS
seas_pvc: seas-pvc                                       # Persistent volume claim for SEAS data

## --- Registry Configuration ---
seas_registry_secret: ibm-entitlement-key                # Secret to pull container images
seas_registry: cp.icr.io                                  # Image registry URL
seas_registry_user: cp                                    # Registry username
seas_registry_key:                                        # Optional registry key, leave empty if not used

## --- Passwords & Secrets ---
seas_secret: ibm-seas-secret                              # Kubernetes secret for SEAS sensitive data
seas_sys_passphrase: Passw0rd@123                         # System passphrase for SEAS
seas_admin_password: Passw0rd@123                         # SEAS admin password
seas_keystore_passphrase: Change1t@                       # Keystore password
seas_truststore_passphrase: Change1t@                     # Truststore password

## --- Version Info ---
versions:
  6.1.0.2.03:
    helm_chart_version: 1.5.4                             # Helm chart version for this SEAS version
```
---