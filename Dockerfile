FROM golang:1.25.2-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN GOOS=windows GOARCH=amd64 go build -o main ./cmd/app/.

RUN chmod +x main

EXPOSE 8000

CMD ["./main"]