# Reviews Microservice

A microservices-demo service that provides reviews capabilities to products.
This service is built, tested and released by travis.

## API Spec

Checkout the API Spec [here](https://github.com/ftuyama/reviews-microservice/master/api-spec/reviews.json)

## To build this service

### Dependencies

```bash
go get -u github.com/FiloSottile/gvt
gvt restore
```

### Go tools

In order to build the project locally you need to make sure that the repository directory is located in the correct
$GOPATH directory: $GOPATH/src/github.com/ftuyama/reviews/. Once that is in place you can build by running:

```bash
cd $GOPATH/src/github.com/ftuyama/reviews/cmd/reviewssvc/
go build -o reviews
```

The result is a binary named `reviews`, in the current directory.

### Docker

`docker-compose build`

### To run the service on port 8080

#### Go native

If you followed to Go build instructions, you should have a "reviews" binary in $GOPATH/src/github.com/ftuyama/reviews/cmd/reviewssvc/.
To run it use:

```bash
./reviews
```

#### Docker
`docker-compose up`

### Run tests before submitting PRs
`make test`

### Check whether the service is alive
`curl http://localhost:8080/health`

### Use the service endpoints
`curl http://localhost:8080/reviews`

### Push the service to Docker Container Registry
`GROUP=weaveworksdemos COMMIT=test ./scripts/push.sh`
