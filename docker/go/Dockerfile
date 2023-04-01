# Use an official Golang runtime as a parent image
FROM golang:1.18-alpine

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY backend/ .

# Build the main package
RUN go build -o main ./cmd/main.go

# Set the startup command to run the executable
CMD ["/app/main"]
