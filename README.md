# GGTIME API
![code coverage badge](https://github.com/Alastair7/ggtime-api/actions/workflows/ci.yml/badge.svg)

## How to run
1. Clone the project.
2. `cd ggtime-api`
3. go run `cmd/api-server/main.go` or `go run ./...`

### Run with docker
1. `docker build -t ggtime-api:latest .`
2. `docker run --name ggtime-container -p 8080:8080 ggtime-api:latest`

## How to run tests
### Run all tests
1. cd to project root `cd ggtime-api`
2. `go test ./...`

### Run specific test
1. `go test [folder-path]/`
2. `go test` to test only the files existing in the current directory.
