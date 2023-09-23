# Terraform Beginner Bootcamp 2023

## Semantic Versioning
My project is using semantic versioning for tagging.
Visit [semver.org](https://semver.org) for more info.

The format implemented is the **MAJOR.MINOR.PATCH** format. eg. ``1.0.1``

- **MAJOR** version when you make incompatible API changes
- **MINOR** version when you add functionality in a backward compatible manner
- **PATCH** version when you make backward compatible bug fixes

## Install Terraform CLI

### Considerations
I had to modify the original install instructions in my [.gitpod.yml](.gitpod.yml) file due to gpg keyring changes.
Updated instructions to install the terraform CLI provided by Hashicorp.

[Install Terraform CLI](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli)

### Refactor into Bash Script

The latest installation instructions were quite lenghty so it was refactored into a bash script. This script is then run at the startup of the Gitpod environment.

Script: [./bin/install_terrafrom_cli](./bin/install_terraform_cli)

- This allows for quick and easy debug and also manual execution.
- Portability in case other projects need the Terraform CLI installed.