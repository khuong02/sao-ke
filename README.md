## Start server
```bash
go run cmd/server/main.go
```

## Build docker
```bash
# migrate
docker build -t sao-ke-migrate:latest . -f dockerfile.migrate
# server
docker build -t sao-ke-be:latest .
# fe
cd ./fe && docker build -t sao-ke-fe:latest .
```

## Run migrate
```bash
docker run -v $PWD/your_file.csv:/chuyen_khoan.csv --rm sao-ke-migrate:latest
```

## Run app
```
docker-compose up -d
```

## Down app
```bash
docker-compose down
```
