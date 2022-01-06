# Security and Firewall

Automate Cluster requires several ports to be open between the Frontend and Backend servers in order to operate. Below it a breakdown of those ports and what needs to be open to each set of servers.

## Incoming frontends network traffic

Provisioning server => Frontends

TCP 22 - This allows terraform to ssh in and configure services

TCP 9631 - This allows our tools to query information from the backend to configure Automate

Users/chef-servers/chef-clients => Frontends

TCP 443 - This allows users to reach the UI and chef-servers to reach the api for reporting. If chef-clients report directly or download profiles, they'll need 443 access as well

TCP 80 - Optional, however if not in place something should redirect users to 443 before they reach it

## Incoming Elastic-search backend network traffic

Provisioning server => elasticsearch-backends

TCP 22 - This allows Terraform to ssh in and configure services

TCP 9631 - This allows our tools to query information from the backend to configure Automate

admin-users => elasticsearch-backends

TCP 5601 - Optional, this allows admins to reach Kibana on the Elasticsearch servers. This hosts an operations dashboard that shows metrics for the Elasticsearch and Postgres servers.

Frontend => elasticsearch-backends

TCP 9200 - This allows Automate to talk to the Elasticsearch API

TCP 9631 - This allows our tools to query information from the backend to configure Automate

elasticsearch-backends <=> elasticsearch-backends

TCP 9300 - This allows Elasticsearch to distribute data in its cluster

TCP/UDP 9638 - This allows Habitat to communicate configuration changes between elasticsearch nodes

TCP 9631 - This allows the Habitat API to be reachable from services on the elasticsearch nodes

postgres-backends <=> elasticsearch-backends

TCP 9200 - This allows Metricbeats to send data to elasticsearch for use in Kibana

TCP/UDP 9638 - This allows Habitat to communicate configuration changes between all backend nodes

TCP 9631 - This allows the Habitat API to be reachable from services on the all backend nodes

Incoming PostgreSQL backend network traffic

Provisioning server => postgres-backends

TCP 22 - This allows Terraform to ssh in and configure services

TCP 9631 - This allows our tools to reach the habitat API for configuration

Frontend => postgres-backends

TCP 7432 - This allows Automate to connect to haproxy which forwards to the psql leader

TCP 9631 - This allows our tools to query information from the backend to configure Automate

postgres-backends <=> postgres-backends

TCP 5432 - This allows HAProxy on postgres-backends to forward connections to the leader

TCP/UDP 9638 - This allows Habitat to communicate configuration changes between Postgres nodes

TCP 9631 - This allows the Habitat API to be reachable from services on the Postgres nodes

elasticsearch-backends <=> postgres-backends

TCP/UDP 9638 - This allows Habitat to communicate configuration changes between all backend nodes

TCP 9631 - This allows the Habitat API to be reachable from services on all backend nodes
