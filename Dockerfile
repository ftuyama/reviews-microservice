FROM golang:1.22-alpine
ENV sourcesdir /go/src/github.com/ftuyama/reviews-microservice/
ENV MONGO_HOST mytestdb:27016
ENV HATEAOS reviews
ENV USER_DATABASE mongodb

COPY . ${sourcesdir}
RUN apk update
RUN apk add git
RUN go mod init reviews
RUN go mod download && go install

ENTRYPOINT reviews
EXPOSE 8084
