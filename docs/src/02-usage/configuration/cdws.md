# Connect Direct Web Services Configuration

This file documents the configuration variables used for deploying IBM Sterling Connect Direct Web Services.

---


##    Core Settings 

```yaml
cd_ws_release_name: auto-cd-ws                 # Helm release name for the deployment
cd_ws_version: 6.4.0.2_ifix001                 # Version of Connect:Direct Web Services to deploy
cd_ws_license_type: non-prod                   # License type (e.g., non-prod, prod)
cd_ws_namespace: connect-direct-web-services   # Kubernetes namespace for deployment
cd_ws_serviceaccount: cd-ws-sa                 # Kubernetes ServiceAccount to use
cd_ws_dashboard_enabled: true                  # Enables the dashboard component
cd_ws_pvc: cd-ws-pvc                           # Persistent Volume Claim for storage
cd_ws_nodename: localhost                      # Node name used within the application
cd_ws_deploy_sum_enabled: 1                    # Whether to deploy SUM (1 = true, 0 = false)
```

---

##    Passwords

```yaml 
cd_ws_keystore_password: Change1t@             # Password for the keystore
cd_ws_truststore_password: Change1t@           # Password for the truststore
cd_ws_cacert_password: Change1t@               # Password for the CA certificate
```

##    Secrets and Certificates

```yaml
cd_ws_deploy_registry_secret: ibm-entitlement-key  # Secret name for registry credentials
cd_ws_deploy_secret: cd-ws-secret              # Secret used for deployment credentials
cd_ws_deploy_cert_secret: cd-ws-cert-secret    # Secret containing TLS certs
cd_ws_cert_crt: /tmp/cdwscert.crt              # Path to certificate CRT file
cd_ws_cert_key: /tmp/cdwskey.pem               # Path to certificate KEY file
cd_ws_cert_pem: /tmp/cdwscert.pem              # Path to certificate PEM file
```


##    Version Mapping

```yaml
versions:
  6.4.0.2_ifix001:
    helm_chart_version: 1.1.6                  # Helm chart version corresponding to the release
```

---