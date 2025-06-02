# IBM Event Streams Configuration

This file documents the configuration variables used for deploying IBM Event Streams.

---

```yaml
# -----------------------------------------------------------------------------
# Kafka Configuration
# -----------------------------------------------------------------------------
# Configuration for deploying Kafka using the IBM Event Streams operator.
# Includes versioning, secrets, storage settings, and authentication.
# -----------------------------------------------------------------------------

## --- Core Settings ---
kafka_version: latest                          # Kafka version to deploy
kafka_operator_version: v3.7.0                 # Operator version used to manage Kafka
kafka_channel: v3.7                            # Operator subscription channel
kafka_namespace: kafka                         # Namespace where Kafka will be deployed

## --- Registry Access ---
registry: cp.icr.io                            # IBM Container Registry
registry_user: cp                              # IBM registry username
kafka_registry_secret: ibm-registry-secret     # Secret used for registry access

## --- Kafka CRD and Resources ---
kafka_crd_name: kafka                          # Name of the Kafka custom resource
kafka_zookeeper_storage_size: 100Gi            # Storage size for Zookeeper
kafka_node_storage_size: 50Gi                  # Storage size for Kafka nodes

## --- Authentication ---
kafka_appuser_password: "Passw0rd"             # Password for the Kafka application user
```
---