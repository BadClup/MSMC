FROM golang:alpine3.20

WORKDIR /app

ARG DEV_MODE
RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN if [ "$DEV_MODE" = "true" ]; then \
    go install github.com/mitranim/gow@latest; \
fi

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN swag init
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

EXPOSE 3000

CMD ["./main"]
