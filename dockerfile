# Stage 1: Build stage
FROM golang:1.22-alpine3.19 AS build

# Set the working directory
WORKDIR /

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o emission-factors-api api/main.go

# Stage 2: Final stage
FROM alpine:edge

# Set the working directory
WORKDIR /

# Copy the binary from the build stage
COPY --from=build /emission-factors-api .

# Set the timezone and install CA certificates
RUN apk --no-cache add ca-certificates tzdata

ENV PORT 8080
EXPOSE $PORT

# Set the entrypoint command
CMD /emission-factors-api --port $(echo $PORT)