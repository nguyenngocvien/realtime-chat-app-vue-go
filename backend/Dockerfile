# Use Go official image for building
FROM golang:1.23 as builder

WORKDIR /app

# Copy your source code
COPY . .

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Build the binary
# RUN go build -o main .

# Use a small image for running the app
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/main .

COPY .env .

# Expose the port if needed (optional)
EXPOSE 8080

RUN chmod +x main

# Run the binary
CMD ["./main"]
