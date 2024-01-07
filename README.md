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
echo PATH:PATH_TO_YOUR_GO/go/bin/go:PATH >> ~/.bashrc
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
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
tar -xvf go1.21.5.linux-amd64.tar.gz
```

## Reference

- [golang http doc](https://go.dev/src/net/http/doc.go)

- [golang net/http package](https://pkg.go.dev/net/http)

- [download and install golang](https://go.dev/doc/install)
