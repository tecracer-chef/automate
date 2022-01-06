
# Backup and restore

Backup configurations to be done before deploying cluster.

## Pre-backup configuration

### Elasticsearch configuration and setup

A shared file system is needed to create Elasticsearch snapshots. In order to register the snapshot repository with Elasticsearch it is necessary to mount the same shared filesystem to the same location on all master and data nodes. This location (or one of its parent directories) must be registered in the `path.repo` setting on all master and data nodes.

Assuming that the shared filesystem is mounted to `/mnt/automate_backups`, we can configure Automate to register the snapshot locations with Elasticsearch.

Ensure the shared file system is mounted on all Elasticsearch servers:
`mount /mnt/automate_backups`

Create elasticsearch sub-directory and set permissions, this will only need to be done on a single Elasticsearch server if the network mount is correctly mounted.

```shell
sudo mkdir /mnt/automate_backups/elasticsearch
sudo chown hab:hab /mnt/automate_backups/elasticsearch
```

Configure Elasticsearch `path.repo` setting by SSHing to a single Elasticsearch server and using the following steps:

Export the current Elasticsearch config from the Habitat supervisor. You will need to have root access to run the following commmands

```shell
source /hab/sup/default/SystemdEnvironmentFile.sh
automate-backend-ctl applied --svc=automate-ha-elasticsearch | tail -n +2 > es_config.toml
```

Edit es_config.toml and add the following settings to the end of the file.
Note: If credentials have never been rotated this file may be empty.

```toml
[es_yaml.path]
  repo = "/mnt/automate_backups/elasticsearch"
```

Apply updated `es_config.toml` config to Elasticsearch, this only needs to be done once. This will trigger a restart of the Elasticsearch services on each server.

`hab config apply automate-ha-elasticsearch.default $(date '+%s') es_config.toml`

Check elasticsearch service is up
`hab svc status`

Another way to check Elasticsearch. Check that all the indices are green
`curl -k -X GET "https://localhost:9200/_cat/indices/*?v=true&s=index&pretty" -u admin:admin`

#### Configure Automate to handle external Elasticsearch backups

Create a automate.toml file on the provisioning server

`touch automate.toml`

Add the following configuration to automate.toml on the provisioning host.

```toml
[global.v1.external.elasticsearch.backup]
enable = true
location = "fs"

[global.v1.external.elasticsearch.backup.fs]
# The `path.repo` setting you've configured on your Elasticsearch nodes must be a parent directory of the setting you configure here:
path = "/mnt/automate_backups/elasticsearch"

[global.v1.backups.filesystem]
path = "/mnt/automate_backups/backups"
```

After that patch the config. This will trigger also the deployment.

`chef-automate config patch automate.toml`

### Management Node IAM Privileges

(TODO: Rework - IAM role, attached policy, etc.)

In order to run the Terraform scripts, we need an IAM user with proper permissions. All the required permissions are mentioned in the next section. We need to make sure that we have the access key ID and secret access key for the user. If not, then regenerate a new access key and keep it handy.

Permissions to be provided:

We need to check if the EC2 Role/IAM user has all the required permissions or not. Listed below are the must the following AWS Managed Policy:

* `AdministratorAccess`

(TODO: This should be checked as it is a security problem - We can offer a training at tecRacer to learn Least Privilege :-) )

### S3 Configuration for Backup

We also have to create IAM role to give access of s3 to elasticsearch instances.

These permissions can either be directly added to the user or can be added via IAM Group.

After doing the above steps, we need to create toml file and patch the config. Please modify the below listed values in the file:

bucket name (bucket = "bucket-name" and name = "bucket-name")

```shell
mkdir configs
vi configs/automate.toml
```

**Put below content in automate.toml file:**

```toml
[global.v1.external.elasticsearch.backup]
enable = true
location = "s3"

[global.v1.external.elasticsearch.backup.s3]
# bucket (required): The name of the bucket
bucket = "bucket-name"
# base_path (optional):  The path within the bucket where backups should be stored
# If base_path is not set, backups will be stored at the root of the bucket.
base_path = "elasticsearch"
# name of an s3 client configuration you create in your elasticsearch.yml
# see <https://www.elastic.co/guide/en/elasticsearch/plugins/current/repository-s3-client.html>
# for full documentation on how to configure client settings on your
# Elasticsearch nodes
client = "default"

[global.v1.external.elasticsearch.backup.s3.settings]
## The meaning of these settings is documented in the S3 Repository Plugin
## documentation. See the following links:
## <https://www.elastic.co/guide/en/elasticsearch/plugins/current/repository-s3-repository.html>
## Backup repo settings

# compress = false
# server_side_encryption = false
# buffer_size = "100mb"
# canned_acl = "private"
# storage_class = "standard"

## Snapshot settings

# max_snapshot_bytes_per_sec = "40mb"
# max_restore_bytes_per_sec = "40mb"
# chunk_size = "null"

## S3 client settings
# read_timeout = "50s"
# max_retries = 3
# use_throttle_retries = true
# protocol = "https"

[global.v1.backups]
location = "s3"

[global.v1.backups.s3.bucket]
# name (required): The name of the bucket
name = "bucket-name"
# endpoint (required): The endpoint for the region the bucket lives in.
# See <https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region>
endpoint = "[https://s3.amazonaws.com](https://s3.amazonaws.com/)"

# base_path (optional):  The path within the bucket where backups should be stored
# If base_path is not set, backups will be stored at the root of the bucket.
base_path = "automate"

[global.v1.backups.s3.credentials]
# Optionally, AWS credentials may be provided. If these are not provided, IAM instance
# credentials will be used. It's also possible for these to be read through the standard
# AWS environment variables or through the shared AWS config files.
# Use the credentials obtained from here [AWS-Credential](https://github.com/chef/automate-as-saas/wiki/Bastion-Setup#aws-credentials)
access_key = "AKIARUQHMSKHGYTUJ&UI”
secret_key = "s3kQ4Idyf9WjAgRXyv9tLYCQgYTRESDFRFV"
```

After putting contents in automate.toml file, we need to eceute below command. This command will also trigger the deployment.

`chef-automate patch configs/automate.toml`

Back-up configurations to be done after deploying cluster

IAM Role: Assign the IAM Role to the all the elastic search instances in the cluster that we create above step.

### File System (EFS)Configuration for backup

Backup on share file system. (This section is specific for aws).

Create the EFS over the AWS.

Once EFS is ready there are 2 ways to mount (via DNS and via IP).

Open the port(2049) Proto(NFS) for EFS security group.

## Backup

To create a new backup run chef-automate backup create from a Chef-Automate front-end node.

`chef-automate backup create`

## Restore

Check status of all Chef Automate and Chef Infra Server front-end nodes.

`chef-automate status`

Shutdown Chef Automate service on all front-end nodes.

`sudo systemctl stop chef-automate`

Login to same instance of Chef Automate front-end node from which backup is taken run the restore command

`chef-automate backup restore --yes -b /mnt/automate\_backups/backups --patch-config /etc/chef-automate/config.toml`

Start all Chef Automate and Chef Infra Server front-end nodes.

`sudo systemctl start chef-automate`

**In case of S3 back-up:**

Login to same instance of Chef Automate front-end node from which backup is taken run the restore command

`chef-automate backup restore s3://bucket_name/path/to/backups/BACKUP_ID --skip-preflight --s3-access-key "Access_Key" --s3-secret-key "Secret_Key"`

Start all Chef Automate and Chef Infra Server front-end nodes.

`sudo systemctl start chef-automate`
