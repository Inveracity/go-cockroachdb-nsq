# go-cockroachdb-nsq
exploring distributing tasks with go, cockroachdb and nsq

# Run

Start the message queue and database

```sh
docker-compose up -d
```

Start distributor and worker

```sh
go run ./cmd/main.go distributor
go run ./cmd/main.go worker
```
