# IBM Sterling Secure Proxy Perimeter Server Configuration

This file documents the configuration variables used for deploying IBM Sterling Secure Proxy Perimeter Server.

---

```yaml
# -----------------------------------------------------------------------------
# IBM SSP PS Configuration
# -----------------------------------------------------------------------------
# This file documents the configuration variables used for deploying IBM SSP PS.
# -----------------------------------------------------------------------------

# --- Core Settings ---
ssp_ps_release_name: auto-ps                             # Helm release name for SSP PS deployment
ssp_ps_version: 6.2.0.1.02                              # SSP PS version
ssp_ps_license_type: non-prod                            # License type: prod, non-prod, etc.

ssp_ps_namespace: ssp                                   # Kubernetes namespace for SSP PS
ssp_ps_serviceaccount: ssp-ps-sa                        # Service account for SSP PS

# --- Registry Configuration ---
ssp_ps_registry_secret: ibm-entitlement-key            # Secret to pull container images
ssp_ps_registry: cp.icr.io                              # Image registry URL
ssp_ps_registry_user: cp                                # Registry username
ssp_ps_registry_key:                                    # Optional registry key, leave empty if not used

# --- Version Info ---
versions:
  6.2.0.1.02:
    helm_chart_version: 1.5.4                           # Helm chart version for this SSP PS version
```
---