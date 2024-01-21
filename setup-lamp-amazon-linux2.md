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
<html>

<head>
    <title>Upload Page</title>
    <style>

        .container {
            max-width: 800px;
            margin-left: auto;
            margin-right: auto;
        }

        .form-grid {
            display: grid;
            row-gap: 10px;
            grid-template-columns: repeat(1, minxmax(0, 1fr));
        }

        .form-upload {
            padding: 2em 2em;
            border: solid 1px black;
        }

        .button-submit {
            max-width: 150px;
        }

    </style>
</head>

<body>
    <div class="container">
        <h1>Photo Uploader</h1>
        <form class="form-upload">
            <div class="form-grid">
                <div>
                    <label for="title">Photo Title</label>
                    <input type="text" id="title" class="input-title" />
                </div>
                <div>
                    <label for="file">Select a photo</label>
                    <input type="file" id="file" class="input-file" />
                </div>
                <div>
                    <label for="description">Description</label>
                    <input type="text" id="description" class="input-desc" />
                </div>
                <div>
                    <label for="date">Date</label>
                    <input type="date" id="date" class="input-date" />
                </div>
                <div>
                    <label for="keyword">Keywords(comman-delimited, e.g. keyword1, keyword2, ...)</label>
                    <input type="text" id="keyword" class="input-keyword" />
                </div>
                <button type="submit" class="button-submit">Upload</button>
            </div>
        </form>
        <a href="/photolookup.php">Photo Lookup</a>
    </div>
</body>

</html>
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
<html>

<head>
    <title>Photo Lookup Page</title>
    <style>

        .container {
            max-width: 800px;
            margin-left: auto;
            margin-right: auto;
        }

        .form-grid {
            display: grid;
            row-gap: 10px;
            grid-template-columns: repeat(1, minxmax(0, 1fr));
        }

        .form-upload {
            padding: 2em 2em;
            border: solid 1px black;
        }

        .button-submit {
            max-width: 150px;
        }

    </style>
</head>

<body>
    <div class="container">
        <h1>Photo Lookup</h1>
        <form class="form-upload">
            <div class="form-grid">
                <div>
                    <label for="title">Photo Title</label>
                    <input type="text" id="title" class="input-title" />
                </div>
                <div>
                    <label for="keyword">Keywords</label>
                    <input type="text" id="keyword" class="input-keyword" />
                </div>
                <div>
                    <label for="from-date">From Date</label>
                    <input type="date" id="from-date" class="input-date" />
                </div>
                <div>
                    <label for="to-date">To Date</label>
                    <input type="date" id="to-date" class="input-date" />
                </div>
                <button type="submit" class="button-submit">Search</button>
            </div>
        </form>
        <a href="/photouploader.php">Photo Uploader</a>
    </div>
</body>

</html>
```

## Vim Basic Command

Save and quite command

```bash
:wq!
```
