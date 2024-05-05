FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy && go mod vendor

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o pokemon_go ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/pokemon_go .

EXPOSE 8080

CMD ["./pokemon_go"]
