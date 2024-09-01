FROM golang:1.22.2-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy
RUN go mod download

COPY . .

RUN go build -o main .

RUN chmod +x /app/main

EXPOSE 8081
в
CMD ["./main"]