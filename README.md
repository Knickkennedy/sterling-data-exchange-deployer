
# Sterling Data Exchange Deployer

A containerized deployer framework designed to streamline the deployment and management of IBM Sterling Data Exchange environments using Ansible and Go.

See Documentation [here](https://knickkennedy.github.io/sterling-data-exchange-deployer/01-getting-started/overview/) to get started.

---

## 🚀 Overview

This project provides a modular and reproducible approach to automate the setup of IBM Sterling Data Exchange environments. It leverages:

- **Ansible**: For defining and executing infrastructure as code via playbooks and roles.
- **Go**: For orchestrating and managing the execution flow through a custom entrypoint.
- **Containerization**: Encapsulates the deployer logic within a container for portability and consistency.

---

## 📁 Repository Structure

```
.
├── .github/workflows/    # CI/CD workflows
├── docs/                 # Documentation files
├── examples/             # Sample configurations and usage examples
├── playbooks/            # Ansible playbooks for deployer tasks
├── roles/                # Ansible roles defining reusable tasks
├── Containerfile         # Definition for building the container image
├── ansible.cfg           # Ansible configuration file
├── entrypoint.go         # Go-based entrypoint for orchestrating tasks
├── go.mod                # Go module definition
├── .dockerignore         # Specifies files to ignore during Docker build
└── .gitignore            # Specifies files to ignore in Git
```

---

## ⚙️ Configuration

- **Ansible Configuration**: Customize `ansible.cfg` to set inventory paths, roles path, and other Ansible behaviors.
- **Environment Variables**: Set environment variables as needed for your playbooks.
- **Secrets Management**: Utilize Ansible Vault or Docker secrets to manage sensitive information securely.

---

## 📄 Documentation

Comprehensive documentation is available in the [docs/](docs/) directory. It includes:

- **Setup Guides**: Instructions for setting up the environment.
- **Usage Examples**: Sample playbooks and configurations.
- **Best Practices**: Recommendations for maintaining and scaling your deployer.

To build and serve the documentation locally:

```
make venv
make serve
```

Then navigate to `http://127.0.0.1:8000/` in your browser.

---

## 📦 Release Information

- **Latest Release**: v6.2.1.0
- **Release Date**: May 26, 2025
- **Changelog**: See [Releases](https://github.com/Knickkennedy/sterling-data-exchange-deployer/releases) for detailed information.

---

## 🤝 Contributing

Contributions are welcome! Please fork the repository and submit a pull request. For major changes, open an issue first to discuss what you would like to change.

---

# License

This project is licensed under the [Apache License 2.0](https://www.apache.org/licenses/LICENSE-2.0).

---

## 📬 Contact

For questions or support, please open an issue in the repository.
