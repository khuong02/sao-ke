FROM golang:1.23.2 AS builder

COPY go.mod go.sum /modules/

WORKDIR /modules
COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -tags migrate -o /bin/app ./cmd/server/main.go

FROM alpine:3.12

COPY --from=builder /bin/app /app

EXPOSE 8080

ENTRYPOINT ["/app"]