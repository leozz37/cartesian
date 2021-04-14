# Cartesian

[![Build](https://github.com/leozz37/cartesian/actions/workflows/build.yml/badge.svg)](https://github.com/leozz37/cartesian/actions/workflows/build.yml)
[![Unit Tests](https://github.com/leozz37/cartesian/actions/workflows/unit_tests.yml/badge.svg)](https://github.com/leozz37/cartesian/actions/workflows/unit_tests.yml)
[![Functional Tests](https://github.com/leozz37/cartesian/actions/workflows/functional_tests.yml/badge.svg)](https://github.com/leozz37/cartesian/actions/workflows/functional_tests.yml)
[![Docker](https://github.com/leozz37/cartesian/actions/workflows/docker.yml/badge.svg)](https://github.com/leozz37/cartesian/actions/workflows/docker.yml)
[![Docker Compose](https://github.com/leozz37/cartesian/actions/workflows/docker_compose.yml/badge.svg)](https://github.com/leozz37/cartesian/actions/workflows/docker_compose.yml)
[![Terraform](https://github.com/leozz37/cartesian/actions/workflows/terraform.yml/badge.svg)](https://github.com/leozz37/cartesian/actions/workflows/terraform.yml)

This is an API for Civi test.

## Contents

- [Architecture](#architecture)
  - [Directory Structure](#directory-structure)
- [Quick-start](#quick-start)
- [Dependencies](#dependencies)
- [Building](#building)
  - [Binary](#binary)
  - [Makefile](#makefile)
  - [Docker](#docker)
  - [Docker-compose](#docker-compose)
  - [Kubernetes](#kubernetes)
- [Testing](#testing)
  - [Unit Tests](#unit-tests)

### Architecture

This is the actual state of the project architecture.

![Architecture](./resources/images/architecture.png)

### Directory Structure

The project is divided by the following strucure:

```txt
.github
    |__ workflows
controllers
data
handler
k8s
models
resources
routes
services
    |__ db
    |__ metrics
terraform
tests
```

These following directory has some misc files, not directly related to the API code:

- [.github/](./github) holds the GitHub templates.

- [.github/worflows](.github/worflows) has the CI pipelines for GitHub Actions.

- [k8s](./k8s) Kubernetes services and pods config files.

- [data](./data) has configuration files for some services.

- [resources](./resources) has configuration files for some services and misc files.

- [terraform](./terraform) has Terraform infra as code files.

- [tests](./tests) has the functional tests written in Robot Framework.

For the API, this is the direcoty rules:

- [handler](./handler) has the endpoint handlers.

- [controllers](./controllers) has the API tools.

- [models](./models) has the most used models across the API.

- [routes](./routes) has the Gin definitions of the endpoints routes.

- [services](./services) holds the services used across the code like the database.

- [services/db](./services/db) holds the databases integrations (MySQL and Redis).

- [services/metrics](./services/metrics) holds the metrics integrations (Prometheus).

## Quick-start

This is an REST API, made with [Golang](https://golang.org/) and [Gin](https://github.com/gin-gonic/gin). You can manually run it or use [docker-compose](https://docs.docker.com/compose/install/) (recommended) or [Kubernetes](https://kubernetes.io/docs/setup/) to get everthing up.

The majoraty of the sensible data is hardcoded on environment variables. Is not a good practice
but for the sake of simplicity, we'll be declaring them on the `.env.example` file. Change it if
needed for your environment.

## Dependencies

In this project we need the following dependencies:

- MySQL
- Prometheus
- Grafana
- Robot Framework

You can install them on your system or run them on containers (besides robot framework). If you're running them manually, set them up on the `.env.example` file to point to the right address or set the environment variables.

To install the API dependencies run:

```shell
$ go mod download
```

## Building

By far, the easiest way to get everything running is with `docker-compose`. See the [docker-compose](#docker-compose) section.

### Binary

To build the binary, run the following:

```shell
$ go build -o cartesian
```

To run the binary, run the following:

```shell
$ ./cartesian
```

Or simply:

```shell
$ go run main.go
```

### Makefile

To run the through the Makefile, run the following:

```shell
$ make run
```

### Docker

Make sure you have [Docker](https://www.docker.com/get-started) installed on your machine.

To build the Docker image, run the following:

```shell
$ docker build . -t cartesian
```

To run the Docker image, set the `$PORT` variable and run:

```shell
$ export PORT=8080

$ docker run --env "PORT=$PORT" -p $PORT:$PORT cartesian
```

### Docker-compose

To run the docker-compose, set the `$PORT` variable and run:

```shell
$ export PORT=8080

$ docker-compose up
```

### Kubernetes

To create the Kubernetes pods, make sure you have [MinuKube](https://minikube.sigs.k8s.io/docs/start/) or [Kind](https://kind.sigs.k8s.io/) installed. In this example, we'll be using `kind`.

Then run the following command:

Create a cluster:

```shell
$ kind create cluster --name=cartesian
```

Select the cluster you just created:

```shell
$ kubectl cluster-info --context kind-cartesian
```

Create the application pods:

```shell
$ kubectl apply -f k8s/
```

Verify pods status:

```shell
$ kubectl get pods
```

Wait until all the pods STATUS are `running`.

You can expose the API service port, running:

```shell
$ kubectl port-forward service/cartesian 8080:8080
```

## Testing

The unit testes are written with the default testing tool of Golang.

### Unit Tests

To run the unit tests, do the following:

```shell
$ go test -v ./...
```

To run the tests with coverage, do the following:

```shell
$ go test -v -covermode=count ./...
````
