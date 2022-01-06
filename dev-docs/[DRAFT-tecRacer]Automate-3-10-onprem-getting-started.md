
# On-premises Configuration

This section is for configuration related information for on-prem deployment. For this type of deployment user will their own provisioned VMs based on the [System requirements](#_System_and_software). Hence no provisioning step is required for on-prem deployment.

[Bastion host setup](#_Bastion_host)

## Prerequisite

List of VM with public and private IP. Public IP is only mandatory for Elasticsearch. (TODO: This makes no sense for us when nothin is exposed to the internet)

All the VM must expose the port 22 for SSH.

We need to open certain port across the VM to make the communication. Please refer this doc for [Firewall and security settings](#_Security_and_firewall) that need to be done before deployment.

| **Component**             | **Port** |
| :------------------------ | :------- |
| Habitat gossip (UDP)      | **9638** |
| Habitat http API          | **9631** |
| PostgreSQL                | **5432** |
| Pgleaderchk               | **6432** |
| HaProxy                   | **7432** |
| Elasticsearch (https)     | **9200** |
| Elasticsearch (transport) | **9300** |
| Kibana                    | **5601** |

## Configuration

Setup configuration file for HA Deployment on on-premises
`chef-automate init-config-ha existing_infra` (TODO: We would prefer a name like `datacenter` or `on-premises`; if this is for existing infrastructure then the title about on-premises might be wrong named?)

This will create configuration for deployment on existing nodes. `config.toml` is the config file where you need to make changes for any change in Automate HA. By default, config file will look like below:

(TODO: Insert config file)

At the end in the `existing_infra` config part, you need to provide IPs of your on premise details separated by comma. we must mention the list of IP address for the cluster. In the below image there are options for private and public IPs (public IP is needed for Elasticsearch only). (TODO: See comment above related to Elasticsearch and public availability)
