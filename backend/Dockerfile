FROM golang:alpine3.20

WORKDIR /app

RUN docker pull itzg/minecraft-server
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN swag init
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

EXPOSE 3000

CMD ["./main"]
