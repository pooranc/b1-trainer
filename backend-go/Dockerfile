FROM golang:1.26.2-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o b1trainer .

EXPOSE 8081

CMD ["./b1trainer"]