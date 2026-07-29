# CI/CD Demo Application Dockerfile
# This creates a portable container with our Go app

# Step 1: Use official Go image as the base (includes Go compiler)
FROM golang:1.19-alpine

# Step 2: Set working directory inside the container
WORKDIR /app

# Step 3: Copy go.mod first for better Docker caching
# (If go.mod doesn't change, Docker reuses cached layer)
COPY go.mod ./

# Step 4: Download dependencies (none in this simple case)
RUN go mod download

# Step 5: Copy all source code into the container
COPY . .

# Step 6: Build the Go application into an executable
RUN go build -o myapp .

# Step 7: Define what runs when container starts
CMD ["./myapp"]
