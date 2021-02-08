# Firemark assignment

The goal of this project is to setup and complete all tasks assigned using a pure golang solution

## Setup

1. The github repo is will be https://github.com/MarioAriasC/calvin
2. Ensure Golang `go (>1.14)` and docker Community Edition `docker > 19` installed
3. To build, go to the source folder and run `docker build .` Run `docker images` and see if firemark-app is created
4. To start the service, go the source folder and run `docker-compose up -d`
5. The docker-compose launches 2 service containers: MySQL DB and firemark-app
6. Hit the server to test health `curl localhost:8000/health` and expect a `200` response 

## Tasks

### Code

There are 2 types of code:

1. Golang application code: is RESTful app is layed with controller, model and repository
2. Unit test, to run unit test, run `go test -v src/note/*.go` To test all testcases

### Documentation and configuration

There are 3 types of documents:

1. This README file as a guide including the general info and setup
2. The dockerfile to build the web application image
3. The docker-compose file to setup and maintain the services

### Implement Tooling

1. Golang (Go Core/Module/Testing) Programming language with testing, module management and documentation tools
2. gorm (https://gorm.io/)  The ORM library for Golang
3. iris (https://www.iris-go.com/) high-performance web applications and APIs framework
4. zerolog (https://github.com/rs/zerolog)  A fast and simple logger dedicated to JSON output(Aggregator friendly)

## Future works
1. Unit testing is for demostration purpose, all the functions in controller are isolated and can be easily tested without dependencies, but the code coverage is pretty low due to time constraints 

2. There are a few TODOs in the code, e.g.
- The configurations are not taken out
- Validations are not sufficient 
- There is an known issue in GROM many2many Preload(repository.go line 77) so the tag model is not returned in the join query

3. Please note there are some setup used to testing purpose only and should not be used for PROD
- The MySQL root account is used for development and testing purpose
- The docker-compose is used for local env
- The MySQL server is hosted in Docker for testing purpose