# Use the official Golang image to create a build artifact.
FROM golang:1.22-alpine as builder

# Copy local code to the container image.
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

# Build the binary.
RUN go build -v -o server

# Use the official Alpine image for a lean production container.
FROM alpine:3
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/server /server

# Run the web service on container startup.
CMD ["/server"]
