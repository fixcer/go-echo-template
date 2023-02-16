# Golang Backend Template Service

A Go (Golang) Backend Template project with Echo, Postgres.

**You can use this project as a template to build your backend project in the Go language on top of this project.**

## Major Packages used in this project

- **echo**: Echo is a fast and unfancy HTTP server framework for Go.
- **oapi-codegen**: OpenAPI/Swagger code generator for Go.
- **sqlc**: Generate type-safe Go from SQL.
- **wire**: Compile-time Dependency Injection for Go.
- **viper**: For loading configuration from the `.yaml` file. Go configuration with fangs. Find, load, and unmarshal a configuration file in JSON, TOML, YAML, HCL, INI, envfile, or Java properties formats.
- Check more packages in `go.mod`.

## Preparing your environment
Install the following tools:
- [Go](https://golang.org/): we use Go to build the project.
- [Protocol Buffer](https://grpc.io/docs/protoc-installation/): we use Protocol Buffer to generate the gRPC code.

## Install dependencies
```
make install
```

## Development

To start your application in the dev profile, run:

```
go run main.go
```
