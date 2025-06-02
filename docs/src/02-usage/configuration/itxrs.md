# Sterling ITX Runtime Services Configuration

This file documents the configuration variables used for deploying IBM Sterling ITX Runtime Services.

---

```yaml
## --- Core Settings ---
itxrs_release_name: auto-itxrs                    # Helm release name for the ITXRS deployment
itxrs_version: 11.0.1.1.20250425                  # Version of the ITXRS container/image
itxrs_license_type: non-prod                     # License type: 'prod' or 'non-prod'
itxrs_namespace: itxrs                            # Kubernetes namespace for ITXRS

## --- Registry Settings ---
itxrs_registry: cp.icr.io/cp                      # IBM container registry path
itxrs_registry_user: cp                           # IBM container registry username
itxrs_registry_key:                               # IBM container registry key (usually provided via secret)
itxrs_registry_secret: ibm-entitlement-key        # Kubernetes secret for pulling images from the registry

## --- Access & Permissions ---
itxrs_serviceaccount: itxrs-sa                    # Kubernetes service account used by ITXRS pods

## --- Version-specific Helm Chart Mapping ---
versions:
  11.0.1.1.20250425:
    helm_chart_version: 3.0.1                     # Helm chart version corresponding to this ITXRS version
```
---