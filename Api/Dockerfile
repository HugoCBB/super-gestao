FROM golang:1.24

WORKDIR /super-gestao-api

COPY go.mod go.sum ./

RUN go mod download

RUN go install github.com/air-verse/air@latest

COPY . .

EXPOSE 8080

CMD ["air"]