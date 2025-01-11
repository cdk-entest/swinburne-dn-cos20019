---
title: assignment 2 note
date: 11/03/2024
---

## Create VPC

- create a new VPC, enable DNS
- add internet gateway, at least 1 NAT gateway in public subnet
- add 2 public subnet, 2 private subnet
- create DB-SG, ALB-SG, WebServer-SG

## Create MySQL RDS

- OK now create a RDS MySQL
- username: admin, password: admin123
- while waiting let create an EC2 for developing the web application

MySQL 8.0.34, private subnet, DB-SG open 3306 for WebServer-SG

## Create S3 Bucket

Let start with a simple s3 bucket policy

- All can get object
- All can put object

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "Test all can get and put object",
      "Effect": "Allow",
      "Principal": "*",
      "Action": ["s3:GetObject", "s3:PutObject"],
      "Resource": "arn:aws:s3:::student-swinburne-11/*"
    }
  ]
}
```

Then update, awllow only LabRole ARN to access the S3 bucket, condition on IP address for more security.

> Please update your bucket name and role arn accordingly.

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "Test all can get and put object",
      "Effect": "Allow",
      "Principal": "*",
      "Action": ["s3:GetObject"],
      "Resource": "arn:aws:s3:::student-swinburne-11/*"
    },
    {
      "Sid": "Only LabRole can get object",
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::654654198339:role/LabRole"
      },
      "Action": ["s3:GetObject", "s3:PutObject"],
      "Resource": "arn:aws:s3:::student-swinburne-11/*",
      "Condition": {
        "IpAddress": {
          "aws:SourceIp": ""
        }
      }
    }
  ]
}
```

## Test uploading image

- create uploads folder
- update s3 bucket policy so that LabInstnaceProfile ARN can put objec to S3 bucket
- use your role ARN
- let test uploading image
- look like an error because no lambda function deploy yet

> Next step: deploy lambda function

## Deploy Lambda Function

- Download the lambda-deployment-package.zip from the assignment page

- Please choose ARM! choose LabRole

- Test the function to see if it can convert image

- Increase time out to 10 seconds if needed

- Now test web server uploading image

- It look like working, let fix image not display late

## Upload S3 Bucket Policy

- Any principale can GetObject
- Only LabRole can PutObject and GetObject
- Condition on IP address later on
- This mean only request from the ALB endpoint can get or download the image
- You can expriment more with S3 bucket policy and write explaination!

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "allow only get request from a specific HTTP",
      "Effect": "Allow",
      "Principal": {
        "AWS": "*"
      },
      "Action": "s3:GetObject",
      "Resource": "arn:aws:s3:::haitran-11032024/*",
      "Condition": {
        "StringLike": {
          "aws:Referer": [
            // please replcae this by your ALB endpoint
            "http://htran-1688150984.us-east-1.elb.amazonaws.com/*"
          ]
        }
      }
    },
    {
      "Sid": "Only LabRole can get object",
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::654654198339:role/LabRole"
      },
      "Action": ["s3:GetObject", "s3:PutObject"],
      "Resource": "arn:aws:s3:::haitran-11032024/*"
    }
  ]
}
```

## ALB and AutoScaling

- AMI Image and Launch Template
- Application Load Balancer, Target Group, Auto Scaling Group

When creating a launch template please take note the following

- Please select LabInstanceRole
- WebServer Security Group

When creating ALB, Target Group please take note

- Heath check path /photoalbum/album.php
- ASG EC2 private subnet with NAT !
- Wait to see a new EC2 ready then update the target group (ASG) with ALB
- Wait a minute to see it become healthy!
- Look like working well!

## Create WebServer

- Create an EC2 instance in public subnet
- Please select amazon linux 2
- Please select LabInstanceProfile
- Use UserData to install PHP

Here is the UserData

```php
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

- Install PHP Admin

```php
cd /var/www/html
wget https://files.phpmyadmin.net/phpMyAdmin/5.2.1/phpMyAdmin-5.2.1-english.zip
unzip phpMyAdmin-5.2.1-english.zip
mv phpMyAdmin-5.2.1-english phpmyadmin
```

Then update config. Let copy database endpoint first

```bash
mv config.sample.inc.php config.inc.php
vim config.inc.php
$cfg['Servers'][$i]['host'] = 'database-1.cdqg6iq4gr9p.us-east-1.rds.amazonaws.com';
```

Update the constants.php. Here is an examples

<details>
<summary>constants.php</summary>

```php
<?php
/**
* 	All constants defined here
*
*	@author Swinburne University of Technology
*
*	============ READ ME !!! ============
* 	============ READ ME !!! ============
* 	============ READ ME !!! ============
*	This project requires AWS SDK for PHP.
*	SSH into your EC2 and execute the following two commands to install the SDK:
*	1. Download the zip file that contains AWS SDK PHP onto /var/www/html directory
*	wget -P /var/www/html http://docs.aws.amazon.com/aws-sdk-php/v3/download/aws.zip
*	2. Unzip the downloaded file onto a new directory called "aws", which sits in /var/www/html directory
*	unzip /var/www/html/aws.zip -d /var/www/html/aws
*
*	Make sure the directory structure is correct so that the AWS SDK can be invoked.
*	var
*	└───www
*		└───html
*			└───aws (created above in Command 2. This directory contains AWS SDK for PHP)
*	    		│   AWS
*	    		│   aws-autoloader.php
*	   			│   ...
*	   			│
*			└───photoalbum (this directory contains source files of the PhotoAlbum website)
*	    		│   uploads (this directory stores images before they are uploaded to S3, for more deets see photouploader.php)
*	   			│   album.php 					(executable) display all images in DB
*	   			│   constants.php 				Constants defined here
*	   			│   defaultstyle.css			CSS style for the website
*	   			│   mydb.php					Interact with RDS DB
*	   			│   photo.php					Photo object class
*	   			│   photouploader.php			(executable) upload image to S3 and RDS
*	   			│   photouploadtemplate.html	HTML template for the photo uploading function
*	   			│   utils.php					some supporting methods
*	   			│
*
*	The values of the constant variables with "[ACTION REQUIRED]" in the comment must be updated. The current values are just examples.
*	You need to replace the values of those constant variables with values specific to your setup.
*
* 	============ READ THE ABOVE !!! ============
* 	============ READ THE ABOVE !!! ============
* 	============ READ THE ABOVE !!! ============
*/

// PLEASE UPDATE THESE YOUR SELF
// [ACTION REQUIRED] your full name
define('STUDENT_NAME', 'Hai Tran');
// [ACTION REQUIRED] your Student ID
define('STUDENT_ID', '100743836');
// [ACTION REQUIRED] your tutorial session
define('TUTORIAL_SESSION', 'Sunday 09:15AM');

// [ACTION REQUIRED] name of the S3 bucket that stores images
define('BUCKET_NAME', 'haitran-11032024');
// [ACTION REQUIRED] region of the above bucket
define('REGION', 'us-east-1');
define('S3_BASE_URL','https://haitran-11032024.s3.amazonaws.com/');

// [ACTION REQUIRED] name of the database that stores photo meta-data (note that this is not the DB identifier of the RDS instance)
define('DB_NAME', 'photoalbum');
// [ACTION REQUIRED] endpoint of RDS instance
// FOR EXAMPLE UPDATE DB ENDPOINT
define('DB_ENDPOINT', 'database-1.cdqg6iq4gr9p.us-east-1.rds.amazonaws.com');
// [ACTION REQUIRED] username of your RDS instance
define('DB_USERNAME', 'admin');
// [ACTION REQUIRED] password of your RDS instance
define('DB_PWD', 'admin123');

// [ACTION REQUIRED] name of the DB table that stores photo's meta-data
define('DB_PHOTO_TABLE_NAME', 'photo_metadata');
// The table above has 5 columns:
// [ACTION REQUIRED] name of the column in the above table that stores photo's titles
define('DB_PHOTO_TITLE_COL_NAME', 'title');
// [ACTION REQUIRED] name of the column in the above table that stores photo's descriptions
define('DB_PHOTO_DESCRIPTION_COL_NAME', 'description');
// [ACTION REQUIRED] name of the column in the above table that stores photo's creation dates
define('DB_PHOTO_CREATIONDATE_COL_NAME', 'creationdate');
// [ACTION REQUIRED] name of the column in the above table that stores photo's keywords
define('DB_PHOTO_KEYWORDS_COL_NAME', 'keywords');
// [ACTION REQUIRED] name of the column in the above table that stores photo's links in S3
define('DB_PHOTO_S3REFERENCE_COL_NAME', 'reference');

// [ACTION REQUIRED] name (ARN can also be used) of the Lambda function that is used to create thumbnails
// create later on
define('LAMBDA_FUNC_THUMBNAILS_NAME', 'CreateThumbnail');

?>
```

</details>

- After modifing the code locally, let upload to remote EC2 using SCP command

```bash
scp -i ~/Downloads/hai.pem * ec2-user@ec2-34-235-148-34.compute-1.amazonaws.com:/var/www/html/photoalbum/
```

## Troubleshooting

- Forget to create uploads folder
- No permissions writting to uploads folder
- S3 bucket policy
- Lamdba function time out (increase to 10 seconds)
- Lambda test event

No permissions for php to write to uploads. Find out which user running httpd -> USER

```bash
ps aux | grep httpd
```

grant owner permissions

```bash
sudo chown -R USER uploads
```

grant write permissions

```bash
sudo chmod -R 0755 USER uploads
```

Client failed to upload to S3 bucket due to no IAM role attached to the EC2

Minh S3

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "allow only get request from a specific HTTP",
      "Effect": "Allow",
      "Principal": {
        "AWS": "*"
      },
      "Action": "s3:GetObject",
      "Resource": "arn:aws:s3:::haitran-11032024/*",
      "Condition": {
        "StringLike": {
          "aws:Referer": "http://htran-1688150984.us-east-1.elb.amazonaws.com/*"
        }
      }
    },
    {
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::654654198339:role/LabRole"
      },
      "Action": ["s3:GetObject", "s3:PutObject"],
      "Resource": "arn:aws:s3:::haitran-11032024/*"
    }
  ]
}
```

Sample lambda test event

```json
{
  "bucketName": "haitran-11032024",
  "fileName": "whale.png"
}
```
