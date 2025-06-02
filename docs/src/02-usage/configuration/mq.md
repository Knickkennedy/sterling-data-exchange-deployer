# IBM MQ Configuration

This file documents the configuration variables used for deploying IBM MQ.

---

##    Core Settings 

```yaml
mq_version: v3.5.3                                           # IBM MQ version
mq_operator_version: 9.4.2.1-r2                              # MQ operator version
mq_channel: v3.5                                             # MQ channel version
mq_namespace: mq                                             # Kubernetes namespace for MQ
mq_crd_name: mq                                              # Custom Resource Definition name
mq_storage_size: 20Gi                                        # Persistent volume size for MQ data
```

---

##    Registry Configuration 

```yaml
mq_registry_secret: ibm-registry-secret                      # Secret for accessing image registry
registry: cp.icr.io                                          # Image registry URL
registry_user: cp                                            # Registry username
```

---

##    Passwords & Secrets 

```yaml
mq_admin_password: "Passw0rd"                                # Administrator password for MQ
mq_appuser_password: "Passw0rd"                              # Password for application user
mq_deploy_secret: "mq-secret"                                # Kubernetes secret for deployment
```

---