# Reviews Service

A reviews service that provides reviews capabilities to products.

>## Build

### Using Go natively

```bash
go build -o reviews
```

>## Run

### Natively

```bash
docker-compose up -d reviews-db
./bin/user -port=8080 -database=mongodb -mongo-host=localhost:27017
```

### Using Docker Compose

```bash
docker-compose up
```

>## Check

```bash
curl http://localhost:8080/health
```

### Reviews

```bash
curl http://localhost:8080/reviews
```
