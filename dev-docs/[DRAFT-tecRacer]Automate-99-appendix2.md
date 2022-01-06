# Appendix
## [What to change in config.toml](https://progresssoftware.sharepoint.com/sites/ChefCoreC/_layouts/15/doc.aspx?sourcedoc=%7bac26b0b0-9621-4d83-a6ef-47c363a9aaf7%7d&action=edit)
\1. Specify the ssh username and the ssh\_key\_file path. This path should be from bastion and If you scroll down in config.toml then you will find 	 here this key pair name and key file both should have a same content. Suppose you have mentioned the "a2ha-aws" name in key\_pair section then put that file content in ssh\_key\_file path's file.



\2. Assign permision for this file "ssh\_key\_file"

chmod 400 /root/.ssh/id

\3. Provide the number of node for the respective cluster for chef\_automate and chef\_server otherwise 1 is fine. For PG,ES and chef-server we have to just maintain the cluster number.



Type of instance (Before proceeding ahead, ensure that the instance type mentioned inside a2ha.rb file is supported in the region) volume type, size and iops aws region ssh\_key\_pair\_name (this is same key-pair name what we have used to provision the bastion ec2 machine). In Section-I we have define the ssh\_key\_file, this should point to the same key file as mentioned in 1st step. Setup the secrets management key and any needed passwords. The default location for the secrets key and secret storage is set in the config file. The default location for the key is /etc/chef-automate/secrets.key and the secret store file is in /hab/a2\_deploy\_workspace/secrets.json.





\4. You have to provide vpcid and cidr block. a2ha will not create a vpc so you have to provide vpc. You can use the default vpc that you'll find in aws VPC section. And also provide cidr block. Use this doc for cidr block [what to write in cidr block](#_What_to_write)

## W[hat to write in cidr block](https://progresssoftware.sharepoint.com/:w:/r/sites/ChefCoreC/_layouts/15/Doc.aspx?sourcedoc=%7B157FBD37-787E-420F-A9F5-9D03A63F2199%7D&file=Vpc%20Instruction.docx&action=default&mobileredirect=true)
How to set vpc

\1. Copy VPC ID from aws web portal as shown in the below image in this field in config.toml.


aws\_vpc\_id = “vpc-c2011ba6”





\2. For Cidr block first go into the subnet section as shown below image.






\3. Then check into the Ipv4 CIDR field in the subnet section in aws web console.





Now as shown in above image IPv4 CIDR we have to pick a unique one. To choose value for ‘aws\_cidr\_block\_addr’ follow below approach

Use This to Select Correct CIDR Block

<https://www.calculator.net/ip-subnet-calculator.html?cclass=a&csubnet=20&cip=172.31.0.0&ctype=ipv4&printit=0&x=82&y=36>

Network class: 172.31.0.0 is a class B ip address so in above image B is selected.

Subnet field: if your vpc IPv4 CIDR block is 172.31.0.0/16 then just add plus 2 to make 18 because we have set /18 in our system. So accordingly, you have to set from subnet drop down. That’s why in above image subnet is selected to /20.

IP Address: Write down 172.31.0.0 from 172.31.0.0/16 from IPv4 CIDR block and press calculate.

After that just scroll up to see below table. In this table you’ll find Network Address field. From that just pick one of them and put into this field‘aws\_cidr\_block\_addr= ‘and do not forgot check that selected address is already occupied or not. Refer step 3 to check.



So as above table give us network address pick unique one. Make sure it should not be conflicted 		already created subnet. See 3rd step aws subnet image. As seen in that image 	all addresses have been picked up so we can’t use that vpc id for this deployment. So, we have to change vpc.

However, if you feel subnet is not possible using above calculator then follow ccna docs. Because 		networking itself a wider topic so not possible to include all the things here.




||||
| :- | :-: | -: |
