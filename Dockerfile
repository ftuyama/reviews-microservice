FROM golang:1.22-alpine
ENV HATEAOS reviews
ENV USER_DATABASE mongodb

RUN apk update
RUN apk add git

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o app .

RUN	chmod +x /app

EXPOSE 80
CMD ["./app"]
