FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download

COPY . .

RUN go build -o /bin/api ./cmd/api

ENTRYPOINT [ "/bin/api" ]