FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /build

EXPOSE 8080

CMD ["/build"]