# PIFLab Store API
[![CircleCI](https://circleci.com/gh/zealotnt/piflab-store-api-go.svg?style=svg)](https://circleci.com/gh/zealotnt/piflab-store-api-go)  
[![Coverage Status](https://coveralls.io/repos/github/zealotnt/piflab-store-api-go/badge.svg)](https://coveralls.io/github/zealotnt/piflab-store-api-go)  

## API Docs
http://docs.piflabstore.apiary.io/

## Dependencies

- **GO 1.5**

## 3rd parties

## Framework

- **Dependency**: [Godep](https://github.com/tools/godep)
- **Router**: [Gorilla Mux](https://github.com/gorilla/mux)

## Build Docker image

`docker build -t piflab-store-api-go .`

## Run a command inside container

``docker run -v `pwd`:/go/src/github.com/o0khoiclub0o/piflab-store-api-go piflab-store-api-go <command you want to run>``

## Add package

- ``docker run -it -v `pwd`:/go/src/github.com/o0khoiclub0o/piflab-store-api-go piflab-store-api-go bash``
- `go get <package>`
- `import "<package>"`
- `godep save ./...`

## Testing

`./testcoverage.sh`

## Migration

### Migrate
`goose up`

### Rollback
`goose down`

### Seed
`go run db/seeds/main.go`

## PIFLab store API services structure
![services-structure](https://www.lucidchart.com/publicSegments/view/9921c064-ed07-4a0f-868c-f9f37def2443/image.png)
