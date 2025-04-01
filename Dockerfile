# Use lightweight Alpine base image
FROM golang:1.24-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy all files from the project directory to the container
COPY . .

# Install necessary dependencies (if any)
RUN go mod tidy

RUN go build -o server .

# Expose the port your app runs on
EXPOSE 3000

# Command to run the Go app
CMD ["./server"]
