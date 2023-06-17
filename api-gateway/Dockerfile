FROM golang:latest as builder

WORKDIR /app

COPY ./api-gateway/go.mod ./api-gateway/go.sum ./

COPY /user-service ../user-service

COPY /apartment-service ../apartment-service

RUN go mod download

COPY ./api-gateway/ .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .



FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8000

CMD ["./main"]