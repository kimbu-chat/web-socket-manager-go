FROM golang:1.17 AS builder
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app cmd/websocketmanager/main.go

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app/app ./app
CMD ["./app"]  
