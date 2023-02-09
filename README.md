# Golang Backend Template Service

## Preparing your environment
Install the following tools:
- [Go](https://golang.org/): we use Go to build the project.
- [Docker](https://www.docker.com/): we use Docker to using OpenAPI Generator.
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

## Building for production

To build the final binary and optimize the senme chatbot service application for production, run:

```
go build -o senme-chatbot-service main.go
```

To load the application in the production config from Spring Cloud Config, run:

```
./senme-chatbot-service --appName senme-chatbot --profile prod --configServer http://localhost:8888
```

Default config server is http://localhost:8888, app name is senme-chatbot and profile is dev. You can change it in config/cloud.go
