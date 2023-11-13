FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN touch .env
RUN CGO_ENABLED=0 GOOS=linux go build -o ./checklist ./src

EXPOSE 8080

CMD ["./checklist"]