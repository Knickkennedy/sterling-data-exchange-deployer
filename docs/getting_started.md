# Getting Started

## Prerequisites

- Docker installed
- Go installed (optional, for building entrypoint)
- Ansible installed (optional, for running playbooks outside container)

## Building the Container

```bash
docker build -t sterling-deployer .
```

## Running the Container

```bash
docker run --rm \
  -v $(pwd)/playbooks:/app/playbooks \
  -v $(pwd)/roles:/app/roles \
  -v $(pwd)/examples:/app/examples \
  sterling-deployer
```

*Adjust volume mounts as needed for your environment.*
