FROM golang:1.24-alpine

WORKDIR /app

## Initialize dependencies
COPY . .
RUN go mod download

## Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o /out/main ./

ENTRYPOINT ["/out/main"]
