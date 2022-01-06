# HA Overview

## Cluster introduction

**High availability (HA)** refers to a system or application (such as a network, a server array, or cluster) that offers a high level of operational performance and quality over a relevant time with maximum potential uptime and accessibility for the content stored on it.

While a more basic system will be adequate to serve content to a low or medium number of users, it may include a single point of failure. This means that if one server goes down, whether due to traffic overload or any number of other issues, the entire site or application could become unavailable.

HA simply means the application remains available with no interruption. We achieve high availability when an application continues to operate when one or more underlying components fail. For example, a router, switch, firewall, or server that fails.

Thus, HA is designed to avoid loss of service by reducing or managing failures and minimizing unscheduled downtime (when your system or network is not available for use or is unresponsive) that happens due to power outages or failure of a component.

“Availability” includes two periods of time: how much time a service is accessible and how much time the system needs to respond to user requests. When it comes to measuring availability, several factors are salient. These include recovery time and both scheduled and unscheduled maintenance periods. Typically, availability is expressed as a percentage of uptime defined by service level agreements (SLAs). A score of 100 percent characterizes a system that never fails or experiences zero downtime by being 100% operational.

## Components

This section lists the various Chef Automate High Availability (HA) components and their purpose.

**Automate-ha-ctl**

Aids connect the backend (postgres and elasticsearch) databases using an automate configuration file and Terraform without any manual intervention.

**Automate-ha-curator**

Elasticsearch curator aids in curating and managing the Elasticsearch indices and snapshots by obtaining the entire actionable list of indices (or snapshots) from the cluster. This component is the same as the default curator. It’s written in a HAB package to merge applications in a HAB environment.

**Automate-ha-deployment**

Aids in setting up a workspace for Chef Automate HA environment. For example, /hab/a2\_deploy\_workspace. It also includes terraform code, some necessary scripts, inspecs, tests, Makefile and so on.

**Automate-ha-elasticsearch**

Includes the Elasticsearch configuration and builds the Elasticsearch package. It is installed in the backend nodes.

**Automate-ha-elasticsidecar**

Provides a sidecar service for automate-backend-elasticsearch that reads user’s credentials and passwords of the Elasticsearch binding and applies it to Elasticsearch using the odfe tooling.

**Automate-ha-haproxy**

Aids in sending a request to the leader node and is placed on postgres cluster.

**Automate-ha-pgleaderchk**

This component is used in a proxy health check to determine where to route SQL requests. A golang service that checks the local PostgreSQL instance to view if it’s a Leader.

**Automate-ha-postgresql**

This component is a wrapper package of core/postgresql11 for Chef Automate that provides a backend HA PostgreSQL.

## Architecture

### Reference Architecture

This section includes Chef Automate High Availability (HA) high-level reference architecture that interacts with the HA backend components on different providers or in different environments.

The following Chef Automate HA architecture diagram shows the components involved in the Chef Automate HA that works on Leader-Follower strategy. We are creating the cluster of the Chef Automate, Chef Server, Postgres, and Elasticsearch for Chef Automate HA.

The Chef Automate HA architecture involves two different clusters part of the main cluster, which are:

#### Automate Backend Cluster

The backend components connect into the frontend habitat supervisor cluster. In the habitat supervisor, postgres and Elasticsearch instances runs. A minimum of three nodes is required for Postgres and Elasticsearch databases, where one becomes a leader, and others are followers.

#### Automate Frontend Cluster

Chef Automate and Chef Server act as frontend nodes and serve as a web UI with load balancer configurations.

These clusters comprise four different servers with HA mode, which are as follows:

Chef-automate

Chef Infra Server

Elasticsearch - an open-source search and analytics engine based on Apache Lucene and built with 	Java. It is a NoSQL database that stores data in an unstructured way.

PostgreSQL - an open-source relational database management system (RDBMS) emphasizing 		extensibility and SQL compliance.

Elastic Search internally manages the communication and backup and does not follow any leader-follower strategy.

## Deployment Support Types

Currently, Chef Automate HA supports two types of deployment, which are

[Cloud deployment - AWS](#_Configuration_and_Provisioning)

[Bare Infrastructure Deployment](#_Configuration_On-prem)

### AWS Deployment

In [AWS deployment](#_Configuration_and_Provisioning), the entire Chef Automate HA infrastructure is built into the AWS cloud. If you choose AWS as a reference architecture, a standard Terraform script handles AWS deployment. This deployment terraform script first sets up all the prerequisites like creating a VPC, EC2, load balancer, security groups, subnets. Then, ensures the VPCID is provided for security purposes, and the cidr block is created manually based on respective VPC.

Later, following series of configurations and installation is performed:

* Installing and configuring PostgreSQL into the postgres instances
* Configuring and installing Elasticsearch into Elasticsearch instances
* Installing a Chef Habitat and creation of a supervisor network.
* Installing automate into the automate instances
* Installing Chef Infra Server in all chef-server instances

### Bare Infrastructure Deployment

Bare infrastructure deployment is for those customers who already have basic network infrastructure with VMs, networks, load balancers in their environment. This environment can be on-premises or in the cloud, and the respective organizations might not want to provide access to create items like VMs. In such cases, IPs of their instances are used to set up Chef Automate HA on their network premises.

If you don’t let Terraform create them, or the customer has already made those by themselves, or customers have on-premises servers, or the customers just want to configure Chef Automate HA (automate, chef-server, Elasticsearch, PostgreSQL) in those servers, and then the customer should choose existing\_node reference architecture.

You can also utilize Terraform script for this scenario; however, then this script only handles installing and configuring components and does not create instances on the cloud providers.
