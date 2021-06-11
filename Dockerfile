# Start from golang base image
FROM golang:alpine as builder

# Enable go modules
ENV GO111MODULE=on

# Install git.
RUN apk update && apk add --no-cache git

# Set current working directory
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

# Download all dependencies.
RUN go mod download

# Now, copy the source code
COPY . .

# Build the application.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main .

# Start a new stage from scratch to build a small image
FROM scratch

# Copy the Pre-built binary file
COPY --from=builder /app/bin/main .
COPY --from=builder /app/.env .
COPY --from=builder /app/db /db

# Run executable
CMD ["./main"]