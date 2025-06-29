FROM golang:1.24 AS builder

WORKDIR /gRPC-context-cancellation

# Copy go.mod and go.sum from root
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build server from its folder
WORKDIR /gRPC-context-cancellation/client
ENV CGO_ENABLED=0
RUN go build -o /gRPC-context-cancellation/bin/client .

FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=builder /gRPC-context-cancellation/bin/client /client

ENTRYPOINT ["/client"]
