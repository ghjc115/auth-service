# syntax=docker/dockerfile:1

FROM golang:1.21rc3

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY . .

RUN go mod download

EXPOSE 7500

# Build
RUN go build auth-service

# Run
CMD ["./auth-service"]
