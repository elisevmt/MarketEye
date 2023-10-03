FROM golang:1.17.2-alpine3.14

WORKDIR /app

RUN ls -la

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

CMD ["go", "run", "/app/cmd/api/main.go"]
