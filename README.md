---
title: build a web app using golang
description: build a web app using golang
author: haimtran
publishedDate: 01/05/2024
date: 2024-01-05
---

## Install Go

First, download go from [HERE](https://go.dev/dl/). For Linux, we can use below command to download a version of Go

```bash
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
```

Second, extract

```bash
tar -xvzf go1.21.5.linux-amd64.tar.gz
```

Next update the PATH environment

```bash
echo export PATH=/home/ec2-user/go/bin:$PATH >> ~/.bashrc
```

Finally, check go version

```bash
go version
```

## Hello World

Let create a new folder called hello

```bash
mkdir helloworld
```

Go into the folder helloworld and init a new go module

```bash
go module init hellomodule
```

Then create a main.go, the project structure look like this

```
|--hello
   |--go.mod
   |--go.sum
   |--main.go
```

Content of main.go

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello World")
}
```

Run the code

```bash
go run main.go
```

## Web App

Let create a web app

- Bedrock stream response
- Book static page
- Upload page

Project structure updated

```
|--go.mod
|--go.sum
|--main.go
|--static
   |--bedrock.html
   |--upload.html
   |--book.html
```

Install dependencies

```bash
go mod tidy
```

Run the web server

```bash
go run main.go
```

## UserData

UserData can be used when launching a new EC2 instance, so it will install GO and clone the repository for the web app

```bash
sudo su ec2-user
cd /home/ec2-user/
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
tar -xvf go1.21.5.linux-amd64.tar.gz
echo 'export PATH=/home/ec2-user/go/bin:$PATH' >> ~/.bashrc

wget https://github.com/cdk-entest/swinburne-dn-cos20019/archive/refs/heads/main.zip
unzip main
cd swinburne-dn-cos20019-main/
go mod tidy
go run main.go
```

## PostgreSQL

Let connect to a postgresql. Default username could be postgres or postgresql.

```bash
psql -h $HOST -p 5432 -U postgres -d demo
```

Create a database and table

```sql
create database demo;
\c demo;
```

Create a book table

```sql
CREATE TABLE IF NOT EXISTS book (
  id serial PRIMARY KEY,
  author TEXT,
  title TEXT,
  amazon TEXT,
  image TEXT
);
```

Insert data into table

```sql
INSERT INTO book (author, title, amazon, image)
values ('Hai Tran', 'Deep Learning', '', 'dog.jpg') RETURNING id;
```

## PostgreSQL and GORM

Create a database connection

```go
const HOST = "localhost"
const USER = "postgres"
const DBNAME = "dvdrental"
const PASS = "Mike@865525"

dns := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v", HOST, "5432", USER, PASS, DBNAME)
	db, _ := gorm.Open(postgres.Open(dns), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase:   false,
			SingularTable: true,
		},
	})

mux.HandleFunc("/postgresql", func(w http.ResponseWriter, r *http.Request) {

		// query a list of book []Book
		books := getBooks(db)

		// load template
		tmpl, error := template.ParseFiles("./static/book-template.html")

		if error != nil {
			fmt.Println(error)
		}

		// pass data to template and write to writer
		tmpl.Execute(w, books)
	})
```

Query database using an Object Relation Model (ORM) such as GORM

```go
func getBooks(db *gorm.DB) []Book {
	var books []Book

	db.Limit(10).Find(&books)

	for _, book := range books {
		fmt.Println(book.Title)
	}

	return books
}
```

Pass books to frontend template

```html
<div class="grid">
  {{range $book:= .}}
  <div class="card">
    <h4 class="title">{{ $book.Image}}</h4>
    <h4 class="title">{{ $book.Author }}</h4>
    <img src="/demo/{{ $book.Image }}" alt="book-image" class="image" />
    <p>
      Lorem ipsum dolor sit amet consectetur, adipisicing elit. Rem quaerat quas
      corrupti cum blanditiis, sint non officiis minus molestiae culpa
      consectetur ex voluptatibus distinctio ipsam. Possimus sint voluptatum at
      modi! Lorem ipsum, dolor sit amet consectetur adipisicing elit. Alias
      dolore soluta error adipisci eius pariatur laborum sed impedit. Placeat
      minus aut perspiciatis dolor veniam, dolores odio sint eveniet? Numquam,
      tenetur! Lorem ipsum dolor sit amet consectetur adipisicing elit. Earum
      suscipit porro animi! Ducimus maiores et non. Minima nostrum ipsa voluptas
      assumenda consequuntur dicta reprehenderit numquam similique, nesciunt
      officiis facere optio. {{ $book.Description}}
    </p>
  </div>
  {{end}}
</div>
```

Upload a file and create a record in the database

```go
// create a record in database
	db.Create(&Book{
		Title:       "Database Internals",
		Author:      "Hai Tran",
		Description: "Hello",
		Image:       handler.Filename,
	})
```

## Reference

- [golang http doc](https://go.dev/src/net/http/doc.go)

- [golang net/http package](https://pkg.go.dev/net/http)

- [download and install golang](https://go.dev/doc/install)
