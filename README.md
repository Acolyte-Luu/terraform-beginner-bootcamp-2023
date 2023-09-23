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



## Working With Environment Variables

### env command

We can list out all enviroment variables (Env Vars) using the `env` command.

We can filter specific env vars using grep eg. `env | grep AWS_`

### Setting and Unsetting Env Vars

we can set using `export HELLO='world` in the terminal.

we can unset using `unset HELLO` in the terminal.

We can also set an env var temporarily when just running a command.

```sh
HELLO='world' ./bin/print_message
```
Within a bash script we can set env without writing export eg.

```sh
#!/usr/bin/env bash

HELLO='world'

echo $HELLO
```

### Printing Vars

We can print or display an env var using the echo command eg. `echo $HELLO`

### Scoping of Env Vars

When you open up new bash terminals in VSCode it will not be aware of env vars that you have set in another window.

If you want your Env Vars to persist across all future bash terminals that are open you need to set env vars in your bash profile. eg. `.bash_profile`

### Persisting Env Vars in Gitpod

We can persist env vars into gitpod by storing them in Gitpod Secrets Storage.

```
gp env HELLO='world'
```

All future workspaces launched will set the env vars for all bash terminals opened in those workspaces.

You can also set env vars in the `.gitpod.yml` but this can only contain non-senstive env vars.


## AWS CLI Installation

AWS CLI is installed for the project via the bash script [`./bin/install_aws_cli`](./bin/install_aws_cli)


[Getting Started Install (AWS CLI)](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
[AWS CLI Env Vars](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html)

We can check if our AWS credentials is configured correctly by running the following AWS CLI command:
```sh
aws sts get-caller-identity
```

If it is succesful you should see a json payload return that looks like this:

```json
{
    "UserId": "AIEAVUO15ZPVHJ5WIJ5KR",
    "Account": "123456789012",
    "Arn": "arn:aws:iam::123456789012:user/terraform-beginner-bootcamp"
}
```

We'll need to generate AWS CLI credentials from the IAM User in order for the user to use AWS CLI.