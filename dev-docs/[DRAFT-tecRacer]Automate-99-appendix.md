(TODO: This should be general knowledge about AWS and probably be trained by partners or refer to the official documentation)


### Amazon’s Virtual Private Cloud (VPC)
#### *VPC requirements*
Amazon VPC, a virtual network dedicated to your AWS account that enables you to launch AWS resources into a virtual network. This virtual network resembles a traditional network that you had operate in your own data center, with the benefits of using the scalable infrastructure of AWS.

Amazon VPC is the networking layer for Amazon EC2. Amazon Elastic Compute Cloud (Amazon EC2) provides scalable computing capacity in the Amazon Web Services (AWS) Cloud. Using Amazon EC2 eliminates your need to invest in hardware up front, so you can develop and deploy applications faster. You can use Amazon EC2 to launch as many or as few virtual servers as you need, configure security and networking, and manage storage. Amazon EC2 enables you to scale up or down to handle changes in requirements or spikes in popularity, reducing your need to forecast traffic.

VPC creates an isolated virtual network environment in the AWS cloud, dedicated to your AWS account. Other AWS resources and services operate inside of VPC networks to provide cloud services. AWS VPC looks familiar to anyone used to running a physical Data Center (DC). A VPC behaves like a traditional TCP/IP network that can be expanded and scaled as needed. However, the DC components you are used to dealing with—such as routers, switches, VLANS, etc.—do not explicitly exist in a VPC. They have been abstracted and re-engineered into cloud software.

All VPCs are created and exist in one—and only one—AWS region. AWS regions are geographic locations around the world where Amazon clusters its cloud data centers.

The advantage of regionalization is that a regional VPC provides network services originating from that geographical area. If you need to provide closer access for customers in another region, you can set up another VPC in that region.

This aligns nicely with the theory of AWS cloud computing where IT applications and resources are delivered through the internet on-demand and with pay-as-you-go pricing. Limiting VPC configurations to specific regions allows you to selectively provide network services where they are needed, as they are needed.

Each Amazon account can host multiple VPCs. Because VPCs are isolated from each other, you can duplicate private subnets among VPCs the same way you could use the same subnet in two different physical data centers. You can also add public IP addresses that can be used to reach VPC-launched instances from the internet.

You can modify or use that VPC for your cloud configurations or you can build a new VPC and supporting services from scratch.
#### *Amazon’s Virtual Private Cloud (VPC) Limit*
Note: - You require a minimum of three node clusters for Elasticsearch and Postgres-sql instances.

AWS limits the size of each VPC; a user cannot change the size once the VPC has been created. Amazon VPC also sets a limit of 200 subnets per VPC, each of which can support a minimum of 14 IP addresses. AWS places further limitations per account / per region, including limiting the number of VPCs to five, the number of Elastic IP addresses to five, the number of Internet gateways per VPC to one, the number of virtual private gateways to five and the number of customer gateways to 50.

CIDR block -Classless Inter-Domain Routing. An internet protocol address allocation and route aggregation methodology. For more information, see Classless Inter-Domain Routing in Wikipedia. Subnet - A range of IP addresses in your VPC.

VPC IP address ranges are defined using Classless interdomain routing (CIDR) IPv4 and IPv6 blocks. You can add primary and secondary CIDR blocks to your VPC, if the secondary CIDR block comes from the same address range as the primary block.
