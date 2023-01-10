# go-api-server-echo
[![](https://img.shields.io/github/go-mod/go-version/golang/go/release-branch.go1.18?filename=src%2Fgo.mod&label=GO%20VERSION&style=for-the-badge&logo=appveyor)](https://github.com/golang/go/releases/tag/go1.18)

The main objective of this repository is to compose a generic API server architecture using the Echo.

## Dependencies

- Go Framework - [Echo](https://github.com/labstack/echo)
- RESTful API doc generater - [echo-swagger](https://github.com/swaggo/echo-swagger)

Detail : go.mod

## Architecture

### Directory
|key|value|
|---|-----|
|configs|environments|
|handlers|controller, routing|
|services|usecase, Application Logic|
|repository|The persistent data store|
|stores|data store|
|docs|API doc to generated with Swagger|

## Installation

Create go.sum:
```sh
$ go mod tidy
```

## Usage

1. run `go run .` to start server.
2. Open http://localhost:3000 with your browser to see the result.
3. Open http://localhost:3000/swagger/index.html, you can see Swagger Api documents.

## Generate API doc using Swagger

1. Install Swag

```sh
$ go install github.com/swaggo/swag/cmd/swag@lates
```

2. Run `swag init` and generate API document in `/docs` directory.

```sh
$ swag init
```

3. Restart Server.

## ToDo
More Code Design.
