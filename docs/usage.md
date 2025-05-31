# Usage

## Running Playbooks

Use the containerized environment or run Ansible directly.

### Container Example

```bash
docker run --rm \
  -v $(pwd)/playbooks:/app/playbooks \
  -v $(pwd)/roles:/app/roles \
  sterling-deployer ansible-playbook playbooks/deploy_sterling_data_exchange.yml
```

### Running Locally

```bash
source ~/.venv/sde-deployer-doc/bin/activate
ansible-playbook playbooks/deploy_sterling_data_exchange.yml
```

## Environment Variables

- `CONFIG_DIR`: Directory path where configs are mounted, e.g. `/app/config`
- `ENTITLEMENT_KEY`: Loaded from Kubernetes secret or environment variable
