---
title: assignment 2
author: haitran
date: 01/07/2204
---

## Step 1. Read Requirements

- Read requirements and rubric
- Download code and specification
- Open spec and see the architecture

## Step 2. VPC and RDS

- Create a VPC with 2 public subnets, 2 private subnets

  - public 1: 10.0.1.0/24
  - public 2: 10.0.2.0/24
  - private 1: 10.0.3.0/24
  - private 2: 10.0.4.0/24

- Create Internet Gateway, and attach to the VPC
- Create one or two NAT gateway in public subnets (NAT should be in PUBLIC SUBNET)
- Create public-route-table

  - create route dest 0/0, target IGW
  - associate with pub-1, pub-2 subnets

- Create a private-route-table

  - create route dest 0/0, target NAT Gateway
  - associate with private-1, private-2, subnets

- Create web-server-security-group, db-security-group, alb-security-group

  - create web-server-security-group, open port 80, 443 for 0/0
  - create db-security-gorup, open port 3306 for web-server-sg
  - create alb-security-group, open port 80, 443 for internet
  - **update web-server-security-group, open 80, 443 for alb-security-group **

- Create a RDS in private subnet
  - MySQL VER 8.0.34, Free Tier
  - In private subnets
  - Take node username and password: admin, Admin2024
  - Disable monitoring, backup for faster launch (NOT PRODUCTION!!!)

Take note RDS endpoint

```bash
database-1.cgzxsxf0wggm.us-east-1.rds.amazonaws.com
username: admin
pass: Admin2024
```

- Install PHP code
- Install DB Admin App and connect to RDS

## Step 3. Web Server and AMI

- Create a EC2 (web server) in public subnet

  - t2.micro
  - Amazon Linux 2 (please make sure this!!!)
  - Public subnet, enable assign public ip
  - web-server-security-group please
  - key-pair (new or already existing with your pem is OK)
  - select LabInstanceProfile
  - Download userdata from assignment 1a (Install_PHP_AWS.rtf)!!!!
  - Paste UserData

```bash
###
#!/bin/bash
yum update -y
amazon-linux-extras install -y lamp-mariadb10.2-php7.2 php7.2
service httpd start
yum install -y httpd mariadb-server php-mbstring php-xml
sed -i "s/upload_max_filesize = 2M/upload_max_filesize = 10M/g" /etc/php.ini
systemctl start httpd
systemctl enable httpd
usermod -a -G apache ec2-user
chown -R ec2-user:apache /var/www
chmod 2775 /var/www
find /var/www -type d -exec sudo chmod 2775 {} \;
find /var/www -type f -exec sudo chmod 0664 {} \;
echo "<?php echo '<h2>Welcome to COS80001. Installed PHP version: ' . phpversion() . '</h2>'; ?>" > /var/www/html/phpinfo.php
```

- Check that the EC2 web server is running
  - Connect, remote access (SSH port 22 open check)
  - Update SecurityGroup open port 22 for 0/0
  - Check that PHP install in /var/www/html/
  - Goto /phpinfo.php to check web server PHP running
  - Good it working

## Step 4. Install MySQL Admin Application

- Download install guide from assignment 1b
- Run some commands to install the admin app
- Update configruation in EC2 to connect RDS

**Because it is not easy sometimes to edit config.inc.php via SSH WEB so let SSH via key pair and edit**

- SSH with key pair
  - Update permission chmod 700 keypair.pem
  - SSH: ssh -i "hai-assignment-2.pem" ec2-user@ec2-44-204-167-89.compute-1.amazonaws.com
  - update RDS endpoint in config.inc.php, localhost
  - check access to the MySQL Admin Page

## Step 5. Update Web Server Code

- Download AWS SDK PHP
- Unzip

```
*	1. Download the zip file that contains AWS SDK PHP onto /var/www/html directory
*	wget -P /var/www/html http://docs.aws.amazon.com/aws-sdk-php/v3/download/aws.zip
*	2. Unzip the downloaded file onto a new directory called "aws", which sits in /var/www/html directory
*	unzip /var/www/html/aws.zip -d /var/www/html/aws
```

- Check project structure
- Let create photoalbum directory and files inside it
- Now update the constants.php
- Create a S3 bucket for storing images
- update S3 bucket name
- update RDS endpoint
- update password to rds
- let create a db in RDS and a table for image metadata
- make sure the schema is correct (score!!!)
- upload an image to s3 and insert a item to db
- update S3 bucket policy so the image can be publicly accessible

```json
{
  "Version": "2012-10-17",
  "Id": "PublicAccess",
  "Statement": [
    {
      "Sid": "PublicAccess",
      "Principal": "*",
      "Effect": "Allow",
      "Action": "s3:GetObject",
      "Resource": "arn:aws:s3:::hai-assignment-2-01072024/*"
    }
  ]
}
```

- now update code in EC2 web server
- now let check album web server
- it work!!!
- let test upload

If you see this permissions error, need to update uploads folder owner permissions !!!

```bash
Permission denied in /var/www/html/photoalbum/photouploader.php on line 51
```

**Next error is due to no lambda setup yet. So it is OK, we will do it later on **

- let check the image in uploads
- let upload another image
- OK it work, can upload

## Step 6. Lambda Function Setup

- Download the lambda code .zip from canvas lambda-deployment-package-0.1.zip
- Create lambda form aws consonle
- Select correct IAM role and Python version 3.11, ARM, name
- Function name: CreateThumbnail
- Let me check role for lambda (**LabRole**)
- Deploy the lambda-deployment-package-0.1.zip
- **Can increase timeout to 10 seconds (cold start issue)**
- Test the web upload again!!!
- No more error !!!! Check S3 to see if Lambda output some images in S3?
- Lambda output resized-cat.png GOOD!

## Step 7. High Available Arch

- Add Auto Scaling Group

  - From running web server create AMI image
  - Create launch template
  - Create Auto Scaling Group
  - Select private subnet
  - You can select desired 2, min 2, max 4, please research yourself

- Add Application Load Balancer

  - ALB should be in public subnet
  - Select ALB security Group
  - Have to create a Target Group only select the webserver for testing! later on select ASG (auto scaling)
  - So we are creating ALB => Target Group => Web Server
  - Wait a few minutes for ALB ON
  - Listener port 80
  - Target Group
  - Finally let update ASG to connect ALB
  - **Then update ALB => Target Group => Auto Scaling Group!!! DONE!**
  - Let me check why no healthy and come back

  - I tried to change health check path to /phpinfo.php
  - De-register targets and Re-register all targets again
  - Now all green and healthy
  - **If having issue here, please let me know**
  - ALB working and let me share my ALB endpoint

```bash
http://htranalb-511854011.us-east-1.elb.amazonaws.com/photoalbum/album.php
```

## Update NACL

Add NACL to further protect web servers (ASG group).

- Create PrivateSubnetNACL
- Block ICMP bidirectional to/from WebServer
- Apply this NACL to private subnets

Here is simple inbound rules

- 100: Open 80 for 10.0.0.0/16
- 200: Open 443 for 10.0.0.0/16
- **300: Open 3306 (MySQL) for 10.0.0.0/16**
- 400: Deny ICMP for source 10.0.0.0/16

Outbound rules

- 100: All all traffic for destination 0/0
- Can be modified to more security!! Think about it your self

Check the ALB endpoint still working!!!

## Update S3 Bucket Policy

Update S3 bucket policy with condition so that only ALB Endpoint can access images in S3 bucket. **You can set up an S3 bucket policy that restricts access to a specific HTTP referrer**. Please update your ALB endpoint. My ALB endpoint: http://htranalb-511854011.us-east-1.elb.amazonaws.com/*

```json
{
  "Version": "2012-10-17",
  "Id": "PublicAccess",
  "Statement": [
    {
      "Sid": "PublicAccess",
      "Principal": "*",
      "Effect": "Allow",
      "Action": "s3:GetObject",
      "Resource": "arn:aws:s3:::hai-assignment-2-01072024/*",
      "Condition": {
        "StringLike": {
          "aws:Referer": [
            "http://htranalb-511854011.us-east-1.elb.amazonaws.com/*"
          ]
        }
      }
    }
  ]
}
```

From ALB can access image

```bash
http://htranalb-511854011.us-east-1.elb.amazonaws.com/photoalbum/album.php
```

- EC2 web server => (Image AccessDenied )
- Open Image directly => (AccessDenied)
