# Global Configuration

This file documents the global configuration variables used to control the deployment of various components in the Sterling Data Exchange Deployer. These variables are typically defined in group variables or passed as extra variables during playbook execution.

---

## Deployment Toggles

```yaml
# Deployment toggles — set true to deploy the component, false to skip

deploy_ibm_catalog_source: false            # Deploy IBM Operator catalog source for IBM operators
deploy_b2bi: true                           # Deploy IBM Sterling B2B Integrator
deploy_ssp_cm: false                        # Deploy Secure Proxy Configuration Manager (SSP CM)
deploy_ssp_engine: false                    # Deploy Secure Proxy Engine (core SSP service)
deploy_ssp_ps: false                        # Deploy Secure Proxy Perimeter Server (SSP PS)
deploy_connect_direct: false                # Deploy IBM Connect:Direct managed file transfer
deploy_connect_direct_web_services: false   # Deploy web services for Connect:Direct
deploy_seas: false                          # Deploy Sterling External Authentication Server (SEAS)
deploy_itx_rs: false                        # Deploy IBM Transformation Extender Runtime Services
deploy_itxa: true                           # Deploy IBM Transformation Extender Advanced (ITXA)
deploy_minio: false                         # Deploy MinIO object storage backend
deploy_kafka: false                         # Deploy Apache Kafka for event-driven messaging

# Container registry settings

registry: cp.icr.io                         # Container image registry URL
registry_user: cp                           # Registry username (often 'cp' for IBM Cloud)

# Storage classes for persistent volumes

file_storage_class: ocs-external-storagecluster-cephfs      # File-based storage class
block_storage_class: ocs-external-storagecluster-ceph-rbd   # Block-based storage class
```

## Product Deployment Explanation

Each of the following boolean flags controls whether a specific component is deployed.

- **deploy_ibm_catalog_source**:  
  Set to `true` to deploy the IBM Operator catalog source. This is typically required if you are installing IBM operators not already available in your cluster.

- **deploy_b2bi**:  
  Enables deployment of IBM Sterling B2B Integrator. Set to `true` if you want to deploy B2Bi in your environment.

- **deploy_ssp_cm**:  
  Deploys the Secure Proxy Configuration Manager component. Set to `true` if SSP CM is part of your architecture.

- **deploy_ssp_engine**:  
  Deploys the Secure Proxy Engine. Set to `true` to include the core engine for secure proxy services.

- **deploy_ssp_ps**:  
  Deploys the Secure Proxy Perimeter Server. Needed for edge security in SSP configurations.

- **deploy_connect_direct**:  
  Enables deployment of IBM Connect:Direct for managed file transfers.

- **deploy_connect_direct_web_services**:  
  Enables the web services interface for IBM Connect:Direct.

- **deploy_seas**:  
  Enables deployment of Sterling External Authentication Server (SEAS).

- **deploy_itx_rs**:  
  Enables deployment of IBM Transformation Extender Runtime Services.

- **deploy_itxa**:  
  Deploys IBM Transformation Extender Advanced (ITXA), used for data transformation and mapping.

- **deploy_minio**:  
  Deploys MinIO as an object storage backend. Use this for S3-compatible storage if no other object store is available.

- **deploy_kafka**:  
  Enables deployment of Apache Kafka, typically for event-driven integrations.

---

## Registry Configuration

- **registry**:  
  The container image registry to pull images from. Example: `cp.icr.io` (IBM Cloud Container Registry).

- **registry_user**:  
  Username for accessing the image registry. Typically set to `cp` for IBM Cloud.

---

## Storage Classes

- **file_storage_class**:  
  The default file-based storage class to use for persistent volumes. Example: `ocs-external-storagecluster-cephfs`.

- **block_storage_class**:  
  The default block-based storage class for components that require block storage. Example: `ocs-external-storagecluster-ceph-rbd`.

---