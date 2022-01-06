# Upgrade

## Upgrade the Complete Cluster

Using this below command it will upgrade both the bundles and deploy it on their respective nodes

`chef-automate upgrade run --upgrade-airgap-bundles`

## Upgrade only Frontend Servers

Using this below command it will upgrade only frontend bundles and deploy it.

`chef-automate upgrade run --upgrade-frontends`

## Upgrade only Backend Servers

Using this below command it will upgrade only backend bundles and deploy it.

`chef-automate upgrade run --upgrade-backends`

In every case you can use skip-deploy flag to only update bundles without deploying them.
For example update only backends:

`chef-automate upgrade run --upgrade-backends --skip-deploy`
