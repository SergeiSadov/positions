FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main cmd/positions/main.go
CMD ["/app/main"]
EXPOSE 6060