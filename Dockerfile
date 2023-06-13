# Use the official Go image as the base image
FROM golang:1.16

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files to the working directory
COPY go.mod go.sum ./

# Download and install the Go dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o instant-messaging-system .

# Expose the port that the application listens on
EXPOSE 8080

# Run the Go application when the container starts
CMD ["./instant-messaging-system"]