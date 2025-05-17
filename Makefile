# Binary name
BINARY_NAME=main
# Docker image name
IMAGE_NAME=realtime-chat-app-vue-go

# Build the Go binary
build:
	go build -o $(BINARY_NAME) .

# Build the Docker image (depends on build)
docker-build: build
	docker build -t $(IMAGE_NAME) .

# Run the Docker container
docker-run:
	docker run --rm -it -p 8080:8080 $(IMAGE_NAME)

# Clean up the binary
clean:
	rm -f $(BINARY_NAME)
