

# Migration
## Chef server (HA- backend) to Automate HA
If existing customer wants to move its existing chef infrastructure to our new a2-ha-backend cluster, it needs to do migration.

**For that we have identified there can be 2 scenarios**

\1. Migrating from standalone chef-server to automate chef-server which is part of a2-ha-backend frontend nodes cluster

\2. Migrating from chef-backend cluster to automate chef-server which is part of a2-ha-backend frontend nodes cluster

In both the cases we need to take backup using knife-ec-backup utility and then move the backup folder on the new chef-server where will take restore using the same utility. This backup will migrate all the cookbooks, users, data-bags, policies and organisations.

knife-ec-backup can backup and restore the data in an Enterprise Chef Server installation, preserving the data in an intermediate, editable text format. It is similar to the knife download and knife upload commands and uses the same underlying libraries, but also includes workarounds for objects not yet supported by those tools and various Server API deficiencies. The long-run goal is to improve knife download, knife upload and the Chef Infra Server API and deprecate this tool.

**This plugin currently supports Enterprise Chef 11 and Chef Infra Server 12+ .**
### Backup on your existing chef-server:
Install habitat

`curl <https://raw.githubusercontent.com/habitat-sh/habitat/master/components/hab/install.sh> | sudo bash`

Install the habitat package for knife-ec-backup

`hab pkg install chef/knife-ec-backup`

Generate a knife tidy server report to examine stale nodes and unused cookbooks

`hab pkg exec chef/knife-ec-backup knife tidy server report --node-threshold 60`


--node-threshold NUM\_DAYS    Maximum number of days since last checking before node is considered stale

Initiate a backup of your Chef Server data

`hab pkg exec chef/knife-ec-backup knife ec backup -c /etc/opscode/chef-server.rb backup\_$(date '+%Y%m%d%H%M%s') --webui-key /etc/opscode/webui\_priv.pem --with-user-sql --with-key-sql`

`--with-user-sql` This is required to correctly handle user passwords and to ensure user-specific association groups are not duplicated.

`--with-key-sql` is to handle cases where customers have users with multiple pem keys associated with their user or clients. The current chef-server API only dumps the default key and sometimes users will generate and assigned additional keys to give additional users access to an account but still be able to lock them out later without removing everyones access.

Run knife tidy server cleans to delete unused data from the reports above.

`hab pkg exec chef/knife-ec-backup knife tidy server clean --backup-path /path/to/an-ec-backup`

Copy backup

Copy the ec backup directory to any Chef Server frontend for the restore target Automate Cluster HA using rsync, NFS, etc or simply copy the folder to chef-server

`scp -i /path/to/key backup\_$(date '+%Y%m%d%H%M%s') user@host:/home/user`


### Restore to Chef Automate HA chef-server

Install the habitat package for knife-ec-backup

`hab pkg install chef/knife-ec-backup`


Restore the backup

`hab pkg exec chef/knife-ec-backup knife ec restore /home/centos/backup\_2021061013191623331154 -yes   --concurrency 1  --webui-key /hab/svc/automate-cs-oc-erchef/data/webui\_priv.pem --purge -c /hab/pkgs/chef/chef-server-ctl/14.1.0/20210225010004/omnibus-ctl/spec/fixtures/pivotal.rb`

## Existing A2HA to Automate HA

In some scenario we are required to migrate A2HA data to Automate HA cluster (as we have new HA implementation in Automate). For that here are some steps that you'll have to follow.

Take backup of chef-automate from A2HA (Old) using below command, this command can be executed from any of front-end (chef-automate) node
in case of multiple frontends. Usually if you don't specify any location for backup in config.toml then that backup will be store on /var/opt/chef-automate/backup location if you hit below command.

sudo chef-automate backup create


Make a tar file of backup that we have taken in step 1. Here make sure that you are also taking backup of .tar directory. Otherwise you'll face some issue related to metadata.json.

E.g. tar -cvf backup.tar.gz path/to/backup/20201030172917/ /path/to/backup/automatebackup-elasticsearch/ /path/to/backup/.tmp/

Create a aib file from any of chef-automate frontend node of A2HA. This will create a bundle of all necessary keys. Like pivotal.pem, secret key etc. Usually this will not be included in regular backup(step 1) so make sure you create a bundle for that.

sudo chef-automate bootstrap bundle create bootstrap.abb

Copy tar file and aib file that we created in step 2 and 3 respectively to any of the Automate HA chef-automate instance and extract it on specific location that you mentioned in config.toml file. Its is not necessary to extract that backup on below location. But make sure that restore command is able to read backup or not from your defined location on step1.

E.g /mnt/automate-backup

Restore A2HA backup on Automate HA. Read this docs for [chef-automate restore](https://docs.chef.io/automate/restore/). In below command there is also a frontend-20210624095659.aib generated in new Automate HA file mention that's because while restoration we also keep in mind all the services habitat release version. Because during restoration time A2HA restoration will try to find A2HA habitat pkg version so there can be a scenario occure where all(A2HA and automate HA frontends (automate's)) packages version can't be the same. That's why we are passing current Automate HA packages. You can find frontend aib file in /var/tmp directory of your Automate HA chef-automate instance.

E.g. sudo chef-automate backup restore /mnt/automate\_backups/backups/20210622065515/ --patch-config /etc/chef-automate/config.toml --airgap-bundle /var/tmp/frontend-20210624095659.aib --skip-preflight

After that if you not do this step then you will might face warning when you'll try to open chef-automate UI It looks like you do not have permission to access any data in automate So make sure you have unpacked the. aib file. Otherwise you'll not see login page. To unpack the bootstrap file that we copied from A2HA chef-automate using below command

Sudo chef-automate bootstrap bundle unpack bootstrap.abb

Copy the bootstrap.aib file to another automate node and chef node also if you are having multiple automate and chef instances. Because the secrets we have restored by unpacking the bootstrap file would be different for another automate instance. So we need to make that all the automate and chef instance would be in sync.

Important command and notes

Using the below command you can see what bootstrap includes in aib file and abb file. Aib file will only include keys related data while abb file will include service packages also. You can use below command and can compare both the file's data

`tail -n +3 bootstrap.aib | tar -tf -`

After using above command, if you want to see the data of service like secret service then you would see those services into /hab/svc directory. This'll be needed if you want to compare aib data between multiple FE(In respective chef-automate and chef-server) nodes.

E.g For secret service `cat /hab/svc/secrets-service/data/secrets\_key`

## A2 to Automate HA
