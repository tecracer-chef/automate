
# System and software requirements

This section lists the recommended operating systems requirements, virtual machine instances requirements for implementing Chef Automate High Availability (HA) for your network infrastructure or systems or applications or services.

## Platform support

| Operating Systems                    | Tested with      |
| :----------------------------------- | :--------------- |
| Red Hat Enterprise Linux (64 Bit OS) | 7, 8             |
| Ubuntu (64 Bit OS)                   | 16.04.x, 18.04.x |
| Centos (64 Bit OS)                   | 7                |

For RHEL 8 or above versions, SELinux configuration must be permissive. By default, in RHEL 8 SELinux configuration is enforced. Red Hat Enterprise Linux derivatives include Amazon Linux v1/v2.

(TODO: What about Ubuntu 20.04? What about Amazon Linux 2022 which is Fedora-based?))

## Virtual Machine (VM) Instances Type

Based on number of nodes (TODO: Align with AWS getting started document)

| Instance          | RAM               | Volume-size |
| :---------------- | :---------------- | :---------- |
| PostgreSQL        | 4 GB RAM for test | 50 GB       |
| Elasticsearch     | 8 GB RAM for test | 50 GB       |
| Chef Automate     | 4 GB RAM for test | 50 GB       |
| Chef Infra Server | 4 GB RAM for test | 50 GB       |

Elasticsearch volume size also depends on the number of nodes and frequency of Chef Infra Client runs and compliance scans. For all the above instances’ RAM and volume size will only for test setup. For production it will depend on number of nodes and frequency of Chef Infra Client runs and compliance scans.

**For Elasticsearch and PostgreSQL, a minimum of three node clusters is required.**

## Bastion Host

(TODO: We would recommend renaming this as "Management Node", as a Bastion Host is mostly considered a "stepping stone" type of setup. A Management Host with full access of an infrastructure, including Terraform, is not really a hardened or secure system)

### Introduction

A [Bastion Host](https://en.wikipedia.org/wiki/Bastion_host) is a special-purpose computer or server on a network specifically designed and configured to withstand attacks. This serve type generally hosts a single application or process, for example, a proxy server or load balancer, and all other services are limited to reduce the threat to the computer.

Its purpose is to provide access to a private network from an external network, such as the Internet or outside of a firewall and involves access from untrusted networks or computers. These computers are also equipped with special networking interfaces to withstand high-bandwidth attacks through the internet.

Bastion servers are instances that resides within your public subnet and accessed using SSH. The purpose of a bastion host is to restrict access to a private network from an external network. Once remote connectivity establishes with the bastion host, it allows you to use SSH to login to other instances (within private subnets) deeper within your network.

The bastion hosts provide secure access to Linux instances located in the private and public subnets.

### AWS Specifics

With AWS Systems Manager Session Manager, access to servers is possible without opening any SSH port solely by authenticating with your AWS credentials. This reduces the risk of open ports and adds auditing features to manual interaction.

Please refer to the AWS documentation on how to enable Systems Manager for your instances and how to connect to them.
