FROM golang:1.22-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main ./cmd/server
EXPOSE 8080
CMD ["./main"]