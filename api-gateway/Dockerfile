FROM golang:1.15-alpine3.13

WORKDIR api-gateway

COPY go.mod .

#COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o api-gateway

EXPOSE 8080

CMD ["./api-gateway"]