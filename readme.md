# go-api-server-echo
[![](https://img.shields.io/github/go-mod/go-version/golang/go/release-branch.go1.18?filename=src%2Fgo.mod&label=GO%20VERSION&style=for-the-badge&logo=appveyor)](https://github.com/golang/go/releases/tag/go1.18)

The main objective of this repository is to compose a generic API server architecture using the Echo.

## Dependencies

- Go Framework - [Echo](https://github.com/labstack/echo)
- RESTful API doc generater - [echo-swagger](https://github.com/swaggo/echo-swagger)
- Environment Variables - [GoDotEnv](https://github.com/joho/godotenv)

Detail : go.mod

## Architecture

### Layered Architecture
Note : The dependency of repository on domain should be fixed. 

```mermaid
flowchart LR
    subgraph Presentation
    id1["handlers/"]
    end
    subgraph Application
    services/
    end
    subgraph Domain
    domains/
    end
    subgraph Repository
    repositorys/
    end
    Presentation --> Application
    Application --> Domain
    Application --> Repository
    Repository --> Domain
```

### Directory
|key|value|
|---|-----|
|configs|environments|
|handlers|controller, routing|
|services|usecase, Application Logic|
|repository|data store|
|docs|API doc to generated with Swagger|

## Installation

Create go.sum:
```sh
$ go mod tidy
```

## Usage

1. Run `go run .` to start server.
2. Open http://localhost:3000 with your browser to see the result.
3. Open http://localhost:3000/swagger/index.html, you can see Swagger Api documents.

## Generate API doc using Swagger

1. Install Swag

```sh
$ go install github.com/swaggo/swag/cmd/swag@lates
```

2. Write API annotations to go file. ([Swagger Document](https://github.com/swaggo/swag#declarative-comments-format))

3. Run `swag init` and generate API documents in `/docs` directory.

```sh
$ swag init
```

4. Restart Server.

## Create Containers for MySQL & Adminer

```sh
$ docker-compose up -d
```

|key|host|port|view|
|---|---|---|---|
|db|127.0.0.1|3306||
|adminer|127.0.0.1|8080|http://localhost:8080/|

## ToDo
More Code Design.
