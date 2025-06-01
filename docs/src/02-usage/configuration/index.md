# Configuration

## Component Configurations

There are examples of each component being defined [here](https://github.com/Knickkennedy/sterling-data-exchange-deployer/tree/main/examples/configs) that will work for demo environments. The [global.yml](https://github.com/Knickkennedy/sterling-data-exchange-deployer/blob/main/examples/configs/global.yml) config is the only one that's required to deploy other products, as it contains metadata needed by the ansible installation role. The rest of the individual products can be run as-is if your environment allows, or with customizations based on the variables provided on their respective configuration page.

## ansible.cfg

Customize Ansible behaviors like inventory path and roles.

## Environment Variables

Define environment variables like `CONFIG_DIR` and `ENTITLEMENT_KEY` for runtime configurations.

## Secrets

Use Kubernetes Secrets or Ansible Vault for sensitive information management.