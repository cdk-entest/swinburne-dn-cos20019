---
title: setup lamp stack on amazon linux2
author: haimtran
date: 21/01/2024
---

## Setup LAMP Stack

Let setup LAMP stack for Amazon Linux 2.

- Option 1. Step by step install [HERE](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-lamp-amazon-linux-2.html)
- Option 2. UserData

Let update

```bash
sudo yum update -y
```

Install mariadb

```bash
sudo amazon-linux-extras install mariadb10.5
```

Install php

```bash
sudo amazon-linux-extras install php8.2
```

Install http server

```bash
sudo yum install -y httpd
```

Enable http service

```bash
sudo systemctl start httpd
sudo systemctl enable httpd
sudo systemctl is-enabled httpd
```

Check the result by open the EC2 public IP address on browser, please ensure SecurityGroup open port 80 for 0/0. Finally here is full UserData

```sh
#!/bin/bash
sudo su ec2-user
sudo yum update -y
yes | sudo amazon-linux-extras install mariadb10.5
yes | sudo amazon-linux-extras install php8.2
yes | sudo yum install -y httpd
sudo systemctl start httpd
sudo systemctl enable httpd
sudo systemctl is-enabled httpd
```

## Photo Upload Page

> [!IMPORTANT]
> Le change the ownership of /var/www/html directory so ec2-user can write to it

```bash
sudo shown -R /var/www/html ec2-user:ec2-user
```

Then create the photo upload lage

```bash
cd /var/www/html/
touch photouploader.php
```

Edit the content of photouploader.php

```bash
vim photouploader.php
```

Here is the content of photouploader.php

```php

```

## Photo Lookup Page

Similarly, create photolookup.php

```bash
touch photolookup.php
```

Edit its content

```bash
vim photolookup.php
```

Here is content of photolookup.php

```php

```

## Vim Basic Command

Save and quite command

```bash
:wq!
```
