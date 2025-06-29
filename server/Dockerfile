FROM golang:1.24 AS builder

# Set working directory inside container
WORKDIR /gRPC-context-cancellation

# Copy go.mod and go.sum from root
COPY go.mod go.sum ./
RUN go mod download

# Copy entire project (including pb, server, etc.)
COPY . .

# Build server from its folder
WORKDIR /gRPC-context-cancellation/server
ENV CGO_ENABLED=0
RUN go build -o /gRPC-context-cancellation/bin/server .

# Final stage
FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=builder /gRPC-context-cancellation/bin/server /server

EXPOSE 50051

# Run the server binary
ENTRYPOINT ["/server"]
