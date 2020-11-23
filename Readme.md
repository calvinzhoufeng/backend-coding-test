# Golang verison 

The goal of this project is to setup and complete all tasks assigned using a pure golang solution
<em>Due to language differences, certain tasks are tweaked into different solutions</em>

## Setup

1. The github repo is https://github.com/calvinzhoufeng/backend-coding-test
2. Ensure Golang `go (>1.13)` is installed
3. `GO111MODULE=on go run src/cmd/main.go` To start the server 
4. `go test -v src/ride/*.go` To test all testcases
5. `GO111MODULE=on go build src/cmd/main.go` To build a binary
6. Hit the server to test health `curl localhost:8010/health` and expect a `200` response 

## Tasks

### Documentation

There are 3 types of documents:

1. This README file as a guide including the general info and setup
2. `godoc -http=localhost:6060` To check the go code documentation in HTML, go to http://localhost:6060/pkg/go/src/ride/ to see details
3. `go test -coverprofile cover.out -v src/ride/*.go` To generate raw coverage report for unit testing
   `go tool cover -html=cover.out -o cover.html` To convert the raw coverage report into html format

### Implement Tooling

1. Golang (Go Core/Module/Test) Programming language with testing, module management and documentation tools
    - Ensure GOMODULE is turned on for compile and build, and module will be automatically downloaded
    - Go test will be triggered for every push
    - Coverage report is generated for each go test executed 
    - Go lint is triggered manually `golint src/ride/`
2. gorm (https://gorm.io/)  The ORM library for Golang
3. iris (https://www.iris-go.com/) high-performance web applications and APIs framework
4. zerolog (https://github.com/rs/zerolog)  A fast and simple logger dedicated to JSON output(Aggregator friendly)
5. Golint (https://godoc.org/golang.org/x/lint/golint) Lints the Go source files 
6. CirclesCI integration
    - Please refer to PR #3 and #4 from branch circleci
    - Each push will trigger a unit test & build into docker, which is configurable, e.g. only main branch is pushed or scheduler

### Pagination 

1. Please refer to PR #2 from branch loadtest

2. Pagination sample code can be found in src/ride/reposiory.go in function Paginate

### Refactoring

This is ignored since Golang implementation is different from NodeJS. If neccessary, certain heavy computation can be moved into go routine functions

### Security

1. Please refer to PR #2 from branch loadtest

2. SQL injection is handled by Golang ORM framework GORM, regarding to the details how SQL injection is handled, please refer to https://gorm.io/docs/security.html.
The rule of thumb is that the Security control should be handled in a separate layer from logic layer, e.g. a middleware

### Load Testing

1. Please refer to PR #2 from branch loadtest
2. `artillery run load-test.yaml` to trigger the load test

