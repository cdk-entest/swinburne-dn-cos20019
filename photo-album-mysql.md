---
title: build a photo album with php and mariadb
author: haimtran
date: 26/01/2024
---

## Setup LAMP

> [!IMPORTANT]
> This script is for Amazon Linux 2!

Let install PHP and MariaDB. Here is basic UserData for getting started

```bash
#!/bin/bash
# install php and mariadb
sudo dnf update -y
sudo dnf install -y httpd wget php-fpm php-mysqli php-json php php-devel
sudo dnf install mariadb105-server
# download code
wget https://github.com/cdk-entest/swinburne-dn-cos20019/archive/php.zip
unzip php.zip
cd swinburne-dn-cos20019-php/
# run the webserver
php -S localhost:3000
```

First, let install PHP and MariaDB.

```bash
sudo dnf update -y
sudo dnf install -y httpd wget php-fpm php-mysqli php-json php php-devel
sudo dnf install mariadb105-server
```

Start the apache web server

```bash
sudo systemctl start httpd
```

Enable the httpd service

```bash
sudo systemctl is-enabled httpd
```

Change folder /var/www/html permissions

```bash
sudo usermod -a -G apache ec2-user
exit
groups
```

Change permissions

```bash
sudo chown -R ec2-user:apache /var/www
sudo chmod 2775 /var/www && find /var/www -type d -exec sudo chmod 2775 {} \;
find /var/www -type f -exec sudo chmod 0664 {} \;
```

Create a index.php in /var/www/html/index.php

```php
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>PHP - Hello, World!</title>
</head>
<body>
        <h1>Hello, World!</h1>
</body>
```

Run the php web server

```bash
php -S 0.0.0.0:3000
```

## Install MariaDB

Start MariabDB

```bash
sudo systemctl start mariadb
```

Stop MariaDB

```bash
sudo systemctl stop mariadb
```

Enable as a service when start the service

```bash
sudo systemctl enable mariadb
```

Since we install it locally, can access db by

```bash
sudo mysql
sudo mysql -h localhost -P 3306 -u root
```

Secure installation [here](https://docs.aws.amazon.com/linux/al2023/ug/ec2-lamp-amazon-linux-2023.html#secure-mariadb-lamp-server-2023)

```bash
sudo mysql_secure_installation
```

## Remote Access MariaDB

Install a client, please follow [here](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_ConnectToMariaDBInstance.html)

```bash
# amazon linux 2023
sudo dnf install mariadb105
# amazon linux 2
sudo yum install mariadb
```

Connect to mariadb

```bash
mysql -h localhost -P 3306 -u <mymasteruser> -p
```

For example

```bash
mysql -h localhost -P 3306 -u dev -p
```

Test

```bash
sudo mysql -h localhost -P 3306 -u root
```

## MariaDB as Admin

Access as root

```bash
sudo mysql
```

List all users

```sql
select user from mysql.user;
```

List databases and table

```sql
show databases;
```

Select a database

```sql
use mysql;
```

List schema

```sql
show schemas;
```

List table of a database

```sql
show tables;
```

Create a database

```sql
CREATE DATABASE IF NOT EXISTS demo;
```

Use the newly created db

```sql
use demo;
```

Create a book table

```sql
CREATE TABLE IF NOT EXISTS book (
    id int auto_increment primary key,
    author text,
    title text,
    amazon text,
    image text);
```

Insert data into book table

```sql
INSERT INTO book(author, title, amazon, image) VALUES ('Hai Tran', 'Deep Learning', '', 'hello.jpg');
```

## MariaDB as User

Create an user with password

```sql
CREATE USER IF NOT EXISTS 'dev'@'localhost' IDENTIFIED by 'Admin2024';
GRANT ALL PRIVILEGES ON * . * TO 'dev'@'localhost';
FLUSH PRIVILEGES;
```

Or granta access to only demo database

```sql
GRANT ALL PRIVILEGES on 'demo'.*to 'dev'@'localhost';
```

Remote connect

```bash
mysql -h localhost -P 3306 -u dev -p Admin2024;
```

## Photo Table

Create the photot table

```sql
CREATE TABLE IF NOT EXISTS photo(
    id int auto_increment primary key,
    title text,
    description text,
    date DATE DEFAULT CURRENT_DATE,
    image text,
    keyword text);
```

Insert a record in the photo table

```sql
INSERT INTO photo(title, description, date, image, keyword) VALUES ('Vim Book', 'Photo of the vim book', '2024-01-26', 'hello.jpg', 'Vim Book');
```

Select records from the photo table

```sql
select * from photo;
```

Update a record in the photo table

```sql
UPDATE photo
SET image = "golang-idiomatic.jpg"
WHERE id = 2;
```

Delete a record in the photo table

```sql
delete from photo where id=1;
```

## Photo Uploader

Let create a form for uploading photo

- Save photo to web server
- Create a record in database (photo table)

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
      cursor: pointer;
    }
  </style>
</head>

<body>
  <div class="container">
    <h1>Photo Uploader</h1>
    <form class="form-upload" action="./handle-form.php" method="post" enctype="multipart/form-data">
      <div class="form-grid">
        <div>
          <label for="title">Photo Title</label>
          <input type="text" id="title" name="title" class="input-title" />
        </div>
        <div>
          <label for="file">Select a photo</label>
          <input type="file" id="file" name="file" class="input-file" />
        </div>
        <div>
          <label for="description">Description</label>
          <input type="text" id="description" name="description" class="input-desc" />
        </div>
        <div>
          <label for="date">Date</label>
          <input type="date" id="date" name="date" class="input-date" />
        </div>
        <div>
          <label for="keyword">Keywords(comman-delimited, e.g. keyword1, keyword2, ...)</label>
          <input type="text" id="keyword" name="keyword" class="input-keyword" />
        </div>
        <button type="submit" class="button-submit">Upload</button>
      </div>
    </form>
    <a href="/photolookup.php">Photo Lookup</a>
  </div>
</body>

</html>
```

To save the uploaded photo and create a new record in database, create handle-form.php as below

```php
<html>

<head>
  <title>Handle Form</title>
  <style>
    :root {
      box-sizing: border-box;
    }

    *,
    ::before,
    ::after {
      box-sizing: inherit;
    }

    body {
      background-color: antiquewhite;
    }
  </style>
</head>

<body>
  <div>
    <?php

    // database connection
    $servername = "localhost";
    $username = "dev";
    $password = "Admin2024";
    $dbname = "demo";

    // create connection
    $conn = new mysqli($servername, $username, $password, $dbname);
    // check connection
    if ($conn->connect_error) {
      die("Connection failed: " . $conn->connect_error);
    }

    // extract form data
    $filename = $_FILES["file"]["name"];
    $title = $_POST["title"];
    $description = $_POST["description"];
    $date = $_POST["date"];
    $keyword = $_POST["keyword"];

    // save file to server
    move_uploaded_file($_FILES["file"]["tmp_name"], "./../static/" . basename($filename));

    // create a record in database
    $sql = "INSERT INTO photo(title, description, date, image, keyword) VALUES ('$title', '$description', '$date', '$filename', '$keyword')";
    $result = $conn->query($sql);

    // resposne to browser
    echo "<h1>" . $_POST["title"] . "</h1>";
    echo "<h1>" . $_FILES["file"]["name"] . "</h1>";
    echo "<h1>" . $_POST["description"] . "</h1>";
    echo "<h1>" . $_POST["date"] . "</h1>";
    echo "<h1>" . $_POST["keyword"] . "</h1>";
    echo "<a href='./photolookup.php'>Photo Lookup</a>";

    // close conn
    $conn->close();
    ?>
  </div>
</body>

</html>
```

## Photo Lookup

Let create a form for searching photos from the photo table.

- photolookup.php as form
- handle-lookup.php to search and return photos

Here is photolookup.php

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
      cursor: pointer;
    }
  </style>
</head>

<body>
  <div class="container">
    <h1>Photo Lookup</h1>
    <form class="form-upload" method="post" action="./handle-lookup.php">
      <div class="form-grid">
        <div>
          <label for="title">Photo Title</label>
          <input type="text" id="title" name="title" class="input-title" />
        </div>
        <div>
          <label for="keyword">Keywords</label>
          <input type="text" id="keyword" name="keyword" class="input-keyword" />
        </div>
        <div>
          <label for="from-date">From Date</label>
          <input type="date" id="from-date" name="from-date" class="input-date" />
        </div>
        <div>
          <label for="to-date">To Date</label>
          <input type="date" id="to-date" name="to-date" class="input-date" />
        </div>
        <button type="submit" class="button-submit">Search</button>
      </div>
    </form>
    <a href="/photouploader.php">Photo Uploader</a>
  </div>
</body>

</html>
```

We need to parse the form and query the photo table. Here is handle-lookup.php

```php
<html>

<head>
  <style>
    :root {
      box-sizing: border-box;
    }

    *,
    ::before,
    ::after {
      box-sizing: inherit;
    }

    .body {
      background-color: antiquewhite;
    }

    .container {
      max-width: 800px;
      margin-left: auto;
      margin-right: auto;
    }

    .grid {
      display: grid;
      row-gap: 10px;
      column-gap: 10px;
      grid-template-columns: repeat(1, minmax(0, 1fr));
    }

    .card {
      margin-left: 4px;
      margin-right: 4px;
      padding: 0.5em;
      background-color: white;
      width: 100%;
    }

    @media (min-width: 800px) {
      .grid {
        grid-template-columns: repeat(2, minmax(0, 1fr));
      }
    }

    .image {
      float: left;
      height: auto;
      width: 128px;
      margin-right: 6px;
    }

    .title {
      font: bold;
      margin-bottom: 8px;
    }
  </style>
</head>

<body class="body">
  <div class="container">
    <?php
    $servername = "localhost";
    $username = "dev";
    $password = "Admin2024";
    $dbname = "demo";

    // create connection
    $conn = new mysqli($servername, $username, $password, $dbname);

    // check connection
    if ($conn->connect_error) {
      die("Connection failed: " . $conn->connect_error);
    }

    // extract form data
    $title = $_POST["title"];
    $keyword = $_POST["keyword"];
    $from_date = $_POST["from-date"];
    $to_date = $_POST["to-date"];

    // query
    $sql = "SELECT title, description, date, image, keyword FROM photo where title like '%$title%'";
    $result = $conn->query($sql);

    if ($result->num_rows > 0) {
      echo "<div class='grid'>";
      // output data of each row
      while ($row = $result->fetch_assoc()) {
        echo "<div class='card'>"
          . "<h4 class='title'>" . $row["title"] . "</h4>"
          . "<h4 class='title'>" . $row["description"] . "</h4>"
          . "<h4 class='title'>" . $row["date"] . "</h4>"
          . "<img src= /static/" . $row["image"] . " class='image' />"
          . "<p>" . "Lorem ipsum, dolor sit amet consectetur adipisicing elit" . "</p>"
          . "</div>";
      }
      echo "</div>";
    } else {
      echo "0 results";
    }
    $conn->close();
    ?>
    <div>
</body>

</html>
```
