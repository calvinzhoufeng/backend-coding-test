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







***
# Xendit Coding Exercise

The goal of these exercises are to assess your proficiency in software engineering that is related to the daily work that we do at Xendit. Please follow the instructions below to complete the assessment.

## Setup

1. Create a new repository in your own github profile named `backend-coding-test` and commit the contents of this folder
2. Ensure `node (>8.6 and <= 10)` and `npm` are installed
3. Run `npm install`
4. Run `npm test`
5. Run `npm start`
6. Hit the server to test health `curl localhost:8010/health` and expect a `200` response 

## Tasks

Below will be your set of tasks to accomplish. Please work on each of these tasks in order. Success criteria will be defined clearly for each task

1. [Documentation](#documentation)
2. [Implement Tooling](#implement-tooling)
3. [Implement Pagination](#implement-pagination)
4. [Refactoring](#refactoring)
5. [Security](#security)
6. [Load Testing](#load-testing)

### Documentation

Please deliver documentation of the server that clearly explains the goals of this project and clarifies the API response that is expected.

#### Success Criteria

1. A pull request against `master` of your fork with a clear description of the change and purpose and merge it
3. **[BONUS]** Create an easy way to deploy and view the documentation in a web format and include instructions to do so

### Implement Tooling

Please implement the following tooling:

1. `eslint` - for linting
2. `nyc` - for code coverage
3. `pre-push` - for git pre push hook running tests
4. `winston` - for logging

#### Success Criteria

1. Create a pull request against `master` of your fork with the new tooling and merge it
    1. `eslint` should have an opinionated format
    2. `nyc` should aim for test coverage of `80%` across lines, statements, and branches
    3. `pre-push` should run the tests before allowing pushing using `git`
    4. `winston` should be used to replace console logs and all errors should be logged as well. Logs should go to disk.
2. Ensure that tooling is connected to `npm test`
3. Create a separate pull request against `master` of your fork with the linter fixes and merge it
4. Create a separate pull request against `master` of your fork to increase code coverage to acceptable thresholds and merge it
5. **[BONUS]** Add integration to CI such as Travis or Circle
6. **[BONUS]** Add Typescript support

### Implement Pagination

Please implement pagination to retrieve pages of the resource `rides`.

1. Create a pull request against `master` with your changes to the `GET /rides` endpoint to support pagination including:
    1. Code changes
    2. Tests
    3. Documentation
2. Merge the pull request

### Refactoring

Please implement the following refactors of the code:

1. Convert callback style code to use `async/await`
2. Reduce complexity at top level control flow logic and move logic down and test independently
3. **[BONUS]** Split between functional and imperative function and test independently

#### Success Criteria

1. A pull request against `master` of your fork for each of the refactors above with:
    1. Code changes
    2. Tests

### Security

Please implement the following security controls for your system:

1. Ensure the system is not vulnerable to [SQL injection](https://www.owasp.org/index.php/SQL_Injection)
2. **[BONUS]** Implement an additional security improvement of your choice

#### Success Criteria

1. A pull request against `master` of your fork with:
    1. Changes to the code
    2. Tests ensuring the vulnerability is addressed

### Load Testing

Please implement load testing to ensure your service can handle a high amount of traffic

#### Success Criteria

1. Implement load testing using `artillery`
    1. Create a PR against `master` of your fork including artillery
    2. Ensure that load testing is able to be run using `npm test:load`. You can consider using a tool like `forever` to spin up a daemon and kill it after the load test has completed.
    3. Test all endpoints under at least `100 rps` for `30s` and ensure that `p99` is under `50ms`
