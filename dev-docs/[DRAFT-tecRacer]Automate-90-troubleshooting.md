# Troubleshooting Guide

## Restore issues

Below are a few frequently encountered issues in Restore and steps on how to resolve them:

### Error: Database is being accessed by other users

**How to resolve:** Some services may be still in running state and are referring to the database, while restore service is trying to drop the database. Please check if all the front end and backend services are stopped.

Error: Cached artifact not found in offline mode

How to resolve: In case of this error, we need to use the `--airgap-bundle` option along with the restore command. Please find the name of the airgap bundle from the path `/var/tmp`. the airgap bundle file would be something like `frontend-20210908093242.aib`

Command:
`chef-automate backup restore s3://bucket_name/path/to/backups/BACKUP_ID --patch-config </path/to/patch.toml> --skip-preflight --s3-access-key "Access_Key" --s3-secret-key "Secret_Key" --airgap-bundle /var/tmp/<airgap-bundle>`

### Error: Existing arch does not match the requested one

How to Resolve:

Execute below commands from bastion from any location.

```shell
sed -i 's/deployment/aws/' /hab/a2_deploy_workspace/terraform/.tf_arch
sed -i 's/architecture "deployment"/architecture "aws"/' /hab/a2_deploy_workspace/a2ha.rb
```

### Other Errors

After running the following deployment command, the deployment repeatedly fails due to `UnhealthyStatusError` (refer screenshot)

`chef-automate deploy config.toml`

Those error can occur during deployment time.

In this above all cases, do the things below.

1. ssh into all frontends (Automate and Chef Infra Server)
2. From all frontends nodes remove `/hab` dir and also remove `/var/tmp` contents.

(TODO: What happens when other software is also using the /var/tmp dir? lsof is not mandatory to be installed on every os)

```shell
rm –rf /hab /var/tmp/*
sudo kill -9 $(sudo lsof -t -i:9631)
sudo kill -9 $(sudo lsof -t -i:9638)
```

1. Do `terraform destroy` of deployment

`(cd /hab/a2_deploy_workspace/terraform && terraform destroy)`

After that, deploy again using:

`chef-automate deploy config.toml`

## Certificates issues

Create `server_cert_ext.cnf` and `client_cert_ext.cnf` files and paste this into them:

`extendedKeyUsage = clientAuth, serverAuth`

Both files should be in same dir where the following script is placed.

Use the script below to generate the required certificates. The contents of the Bash script below needs to be copied in the same directory like the two files from above. Then execute it on the management node with `source cert.sh`.

The script will copy the generated certificates into `/hab/a2_deploy_workspace/certs` directory. It is required to have the openssl utility to be installed.

Contents of `cert.sh`

```shell
#!/bin/bash

openssl genrsa -out ca_root.key 2048
openssl req -x509 -new -key ca_root.key -sha256 -out ca_root.pem -subj '/C=US/ST=Washington/L=Seattle/O=Chef Software Inc/CN=chefrootca'
openssl genrsa -out admin-pkcs12.key 2048
openssl pkcs8 -v1 "PBE-SHA1-3DES" -in admin-pkcs12.key -topk8 -out es_admin_ssl_private.key -nocrypt
openssl req -new -key es_admin_ssl_private.key -out admin.csr -subj '/C=US/ST=Washington/L=Seattle/O=Chef Software Inc/CN=chefadmin'
openssl x509 -req -in admin.csr -CA ca_root.pem -CAkey ca_root.key -CAcreateserial -out es_admin_ssl_public.pem -sha256 -extfile server_cert_ext.cnf
openssl genrsa -out ssl-pkcs12.key 2048
openssl pkcs8 -v1 "PBE-SHA1-3DES" -in ssl-pkcs12.key -topk8 -out es_ssl_private.key -nocrypt
openssl req -new -key es_ssl_private.key -out ssl.csr -subj '/C=US/ST=Washington/L=Seattle/O=Chef Software Inc/CN=chefnode'
openssl x509 -req -in ssl.csr -CA ca_root.pem -CAkey ca_root.key -CAcreateserial -out es_ssl_public.pem -sha256 -extfile client_cert_ext.cnf

cp ca_root.pem /hab/a2_deploy_workspace/certs/ca_root.pem
cp es_admin_ssl_public.pem /hab/a2_deploy_workspace/certs/es_admin_ssl_public.pem
cp es_admin_ssl_private.key /hab/a2_deploy_workspace/certs/es_admin_ssl_private.key
cp es_ssl_public.pem /hab/a2_deploy_workspace/certs/es_ssl_public.pem
cp es_ssl_private.key /hab/a2_deploy_workspace/certs/es_ssl_private.key
cp es_admin_ssl_private.key /hab/a2_deploy_workspace/certs/kibana_ssl_private.key
cp es_admin_ssl_public.pem /hab/a2_deploy_workspace/certs/kibana_ssl_public.pem
cp es_ssl_private.key /hab/a2_deploy_workspace/certs/pg_ssl_private.key
cp es_ssl_public.pem /hab/a2_deploy_workspace/certs/pg_ssl_public.pem
```

Now apply the new SSL certificates from `/hab/a2_deploy_workspace`:

`./scripts/credentials set ssl --es-ssl`
`./scripts/credentials set ssl --kibana-ssl`
`./scripts/credentials set ssl --pg-ssl`

After applying the certificates successfully, you will see an output like this.
(TODO: Give example here)

Go to Chef Automate and Chef Infra Server instance and check the chef service status with `chef-automate status`. If you see a service down or critical, then just wait for 3-4 minutes because after applying the certs it will take some time.

To rotate every SSL certificate, you can also run
`./scripts/credentials set ssl --rotate-all`

If you are running the command once before running the cert.sh script, it will generate a skeleton of required files.
