FROM --platform=linux/amd64 golang:1.23-alpine AS builder

# Working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY ./go.mod ./
COPY ./go.sum ./

# Copy the source code
COPY ./internal/ ./internal/
COPY ./cmd/ ./cmd/

# Copy HTML templates
COPY ./templates/ ./templates/

# Copy secrets
COPY ./secrets/ ./secrets/

# Build the project
RUN go build -o app ./cmd/oaiproxy/main.go

# Start a new stage
FROM alpine:latest

# Copy the executable from the builder stage
COPY --from=builder ./app ./

# Expose port 8080
EXPOSE 8080

# Run the executable
CMD [ "./app" ]
