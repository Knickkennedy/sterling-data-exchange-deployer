# Sterling ITXA Configuration

This file documents the configuration variables used for deploying IBM Sterling ITXA.

---

```yaml
##    Core Settings 
itxa_helm_release_name: auto                     # Helm release name for ITXA
itxa_version: 10.0.1.10-x86_64                   # ITXA version to deploy
itxa_license_type: non-prod                      # License type: 'prod' or 'non-prod'

##    Namespace & Access 
itxa_namespace: itxa                             # Kubernetes namespace for ITXA
itxa_serviceaccount: itxa-sa                     # Service account used by ITXA pods

##    Registry Settings 
itxa_registry: cp.icr.io                         # IBM container registry
itxa_registry_user: cp                           # IBM registry username
itxa_registry_key:                               # Registry key (leave blank if provided via secret)
itxa_registry_secret: ibm-registry-secret        # Secret containing credentials for image pulls

##    Admin Credentials & TLS 
itxa_admin_secret: itxa-admin-secret             # Secret containing admin credentials
itxa_admin_password: passw0rd                    # Admin password

itxa_tls_secret: itxa-tls-secret                 # TLS secret used for secure communication
itxa_tls_keystore_password: Passw0rd             # Password for the TLS keystore

##    Storage & Logs 
itxa_app_logs_size: 5Gi                          # Log volume size for the application

##    Data Setup Controls 
itxa_data_setup_enabled: true                    # Enable data setup
itxa_data_setup_upgrade: false                   # Run data setup in upgrade mode

##    Database Configuration 
itxa_db_vendor: DB2                              # Database vendor: DB2 or Oracle
itxa_use_oracle_service_name: false              # Only set to true if using Oracle service names
itxa_db_secret: itxa-db-secret                   # Secret containing DB credentials

##    Ingress & Services 
itxa_ingress_secret: itxa-ingress-secret         # Secret used for ingress TLS termination
itxa_svc_web: "{{ itxa_helm_release_name }}-ibm-itxa-prod-itxauiserver"  # Internal web service name

##    Optional Component Deployments 
edi_deploy: false                                # Deploy EDI module
fsp_deploy: false                                # Deploy Financial Services Processor module
hc_deploy: false                                 # Deploy Healthcare module

##    Version-Specific Helm Chart Mapping 
versions:
  10.0.1.10-x86_64:
    helm_chart_version: 1.0.4                     # Helm chart version for this ITXA release
    itx_version: 11.0.1.1.20250425                # Associated ITX Runtime version
```
---