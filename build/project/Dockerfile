FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN GOPROXY=https://goproxy.cn go get github.com/klauspost/compress@v1.17.0
RUN go mod download

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -d cmd -o docs/api
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/app cmd/main.go

FROM alpine:latest

WORKDIR /usr/local/bin
COPY --from=builder /go/bin/app /usr/local/bin/app
COPY --from=builder /app/config.json /usr/local/bin/config.json
COPY --from=builder /app/docs/api/swagger.json /usr/local/bin/docs/api/swagger.json

EXPOSE 8080

CMD ["/usr/local/bin/app", "--config", "/usr/local/bin/config.json"]