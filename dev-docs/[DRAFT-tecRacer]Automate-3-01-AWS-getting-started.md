# Getting started

## Package download

Chef-automate is the main utility used for installation of chef-automate. If you are doing installation on fresh server where you don’t have chef-automate utility, you can download it using below link

`curl https://packages.chef.io/files/current/latest/chef-automate-cli/chef-automate_linux_amd64.zip | gunzip - > /usr/local/bin/chef-automate && chmod +x /usr/local/bin/chef-automate`

## Configuration and Provisioning - AWS

This section is only for cloud deployment. Currently we support AWS based provisioning and deployment.

### Sizing

Please refer to [Common System Requirements](#_System_and_software) for general requirement guidelines.

The following recommendations are depending on the number of managed nodes.

#### Production

(TODO: Check, if Elasticsearch is correct; changed the vCPU to the one AWS is providing with that instance type)

| Instance          | Type      | RAM   | vCPU   | Root Volume Size |
| :---------------- | :-------- | :---- | :----- | ---------------- |
| PostgreSQL        | m5.large  | 8 GB  | 2 vCPU | 50 GB            |
| Elasticsearch     | m5.xlarge | 16 GB | 4 vCPU | 50 GB            |
| Chef Automate     | m5.large  | 8 GB  | 2 vCPU | 50 GB            |
| Chef Infra Server | m5.large  | 8 GB  | 2 vCPU | 50 GB            |

#### Test Recommendation

| Instance          | Type      | RAM  | vCPU   | Root Volume Size |
| :---------------- | :-------- | :--- | :----- | ---------------- |
| PostgreSQL        | t3.medium | 4 GB | 2 vCPU | 50 GB            |
| Elasticsearch     | t3.large  | 8 GB | 2 vCPU | 50 GB            |
| Chef Automate     | t3.medium | 4 GB | 2 vCPU | 50 GB            |
| Chef Infra Server | t3.medium | 4 GB | 2 vCPU | 50 GB            |

ES volume size also depends on the number of nodes and frequency of Chef Infra Client runs and compliance scans. The above table includes AWS instance types.

For Elasticsearch and PostgreSQL, a minimum of three node clusters is required. (TODO: Check, if this is true)

### Configuration

Create a config.toml for AWS using below command

`chef-automate init-config-ha aws`

The config.toml is the config file where you need to make changes for any change in Automate HA.

By default, config file will look like this:

(TODO: add example)

Here, you should make all the changes required for AWS deployment. Refer to the [config.toml documentation](#_What_to_change) for details.

### Provisioning

Now, to provision the cloud infrastructure as per configuration provided enter the following command referencing your `config.toml` file

`chef-automate provision-infra config.toml`

This will run a deployment on AWS. It will download habitat, create a workspace in `/hab/a2_deploy_workspace` on the current system and will provision the infrastructure on AWS via Terraform.
