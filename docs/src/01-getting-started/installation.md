
# Installation

Follow these steps to install the Sterling Data Exchange Deployer:

## 1. Clone the Repository

```bash
git clone https://github.com/Knickkennedy/sterling-data-exchange-deployer.git
cd sterling-data-exchange-deployer
```

## 2. Set Up the Environment

Ensure you have the following tools installed:

- Ansible  
- Docker  
- Go

## 3. Build the Container Image

```bash
docker build -t sde-deployer .
```

## 4. Run the Deployer

```bash
docker run --rm -v $(pwd):/app sde-deployer
```

---

For detailed configuration options, refer to the [Configuration](https://knickkennedy.github.io/sterling-data-exchange-deployer/02-usage/configuration) section.
