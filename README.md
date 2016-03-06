# PIFLab Store API
[![Circle CI](https://circleci.com/gh/o0khoiclub0o/piflab-store-api-go.svg?style=svg&circle-token=b62eec2adc4baa81e1e0d75b704de98d94b49be6)](https://circleci.com/gh/o0khoiclub0o/piflab-store-api-go)

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

## Development

``docker run -p 80:80 -v `pwd`:/go/src/github.com/o0khoiclub0o/piflab-store-api-go piflab-store-api-go gin -p 80 run``

## Testing

``docker run -v `pwd`:/go/src/github.com/o0khoiclub0o/piflab-store-api-go piflab-store-api-go ginkgo -r``

## Production

`docker run -p 80:80 piflab-store-api-go`

## Deployment
