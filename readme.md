# Go Project - Icon PLN Test

Project for ICON PLN Test.

## Prerequisites

- [Go](https://golang.org/dl/) (version 1.20)
- [Docker](https://www.docker.com/get-started)

## Makefile Commands

### Run the Application

This command will tidy the `go.mod` file, vendor dependencies, and run the main application.

```sh
make run
```

### Test the Application
This command will run the tests and generate a coverage report.

```sh 
make test
```

### Clean Up
This command will remove the coverage report file.

```sh
make clean
```

### Dokerize App
It's fine if you want use it as container (Optional)

```sh
make docker-build
make docker-run
```