# Connect Direct Configuration

This file documents the configuration variables used for deploying IBM Sterling Connect Direct.

---

```yaml
##    Core Settings 
cd_release_name: auto-cd                         # Helm release name for Connect:Direct
cd_version: 6.4.0.1_iFix029                      # Version of Connect:Direct
cd_license_type: non-prod                        # License type: 'non-prod' or 'prod'
cd_namespace: connect-direct                     # Kubernetes namespace to deploy into
cd_serviceaccount: cd-sa                         # Kubernetes ServiceAccount used by pods
cd_dashboard_enabled: true                       # Enables the dashboard component
cd_pvc: cd-pvc                                   # Persistent Volume Claim for data
cd_nodename: CDNODE01                            # Node name registered in Connect:Direct
cd_deploy_sum_enabled: 1                         # Whether to deploy the SUM component (1 = true)

##    Passwords 
cd_admin_password: passw0rd                      # Password for the Connect:Direct admin user
cd_appuser_password: passw0rd                    # Password for the application user
cd_local_cert_phrase: changeit                   # Passphrase for the local certificate
cd_keystore_password: changeit                   # Password for the keystore file

##    LDAP Configuration 
cd_ldap_enabled: false                           # Enable LDAP integration
cd_ldap_host: ""                                 # LDAP server host
cd_ldap_port: ""                                 # LDAP server port
cd_ldap_domain: ""                               # LDAP domain name
cd_ldap_tls: ""                                  # Use TLS for LDAP connection
cd_ldap_start_tls: ""                            # Use StartTLS for LDAP connection
cd_ldap_ca_cert: ""                              # Path or reference to LDAP CA certificate
cd_ldap_tls_req_cert: ""                         # TLS certificate requirement policy
cd_ldap_default_bind_dn: ""                      # Default bind distinguished name
cd_ldap_default_auth_token_type: ""              # Auth token type (e.g., simple, SASL)
cd_ldap_default_auth_token: ""                   # Token used for default authentication
cd_ldap_client_validation: ""                    # Whether to validate LDAP client certificate
cd_ldap_client_cert: ""                          # LDAP client certificate (if needed)
cd_ldap_client_key: ""                           # LDAP client private key (if needed)

##    Secrets and Certificates 
cd_deploy_registry_secret: ibm-entitlement-key   # Secret for accessing IBM container registry
cd_deploy_secret: cd-secret                      # Secret for Connect:Direct application credentials
cd_deploy_cert_secret: cd-cert-secret            # Kubernetes TLS secret for SSL certs
cd_cert_crt: /tmp/cdcert.crt                     # Path to the certificate file
cd_cert_key: /tmp/cdkey.pem                      # Path to the private key file
cd_cert_pem: /tmp/cdcert.pem                     # Path to combined certificate PEM file

##    Version Mapping 
versions:
  6.4.0.1_iFix029:
    helm_chart_version: 1.4.4                    # Helm chart version associated with this app version
```
---