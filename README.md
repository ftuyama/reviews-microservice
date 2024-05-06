# Reviews Service

A reviews service that provides reviews capabilities to products.

>## Build

### Using Go natively

```bash
go mod init reviews
go mod download
go build -o reviews
```

>## Run

### Natively

```bash
docker-compose up -d reviews-db
./reviews -port=8080 -database=mongodb -mongo-host=localhost:27016
```

### Using Docker Compose

```bash
docker-compose up
```

## Reviews

```bash
curl http://localhost:8080/reviews
```
