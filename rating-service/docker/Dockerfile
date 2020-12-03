# Service dev image
FROM golang:alpine as dev
# Name service
LABEL Name=middleware-rating-service
# Enable gomod & other env
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
# Setup working directory /build
WORKDIR /build
# Copy source code into the container
ADD . /build
# Copy & download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download
# Build the service
RUN go build -o main .
# Move to /dist directory - place for resulting binaries
WORKDIR /dist
# Copy binary from build to main folder
RUN cp /build/main .
# Export necessary port
EXPOSE 8080
# Command to run service container when the container starts
CMD ["/dist/main"]


# Service prod image
FROM alpine:3.9.2 as prod
# Setup working directory /root/
WORKDIR /root/
# Copy service resulting binaries
COPY --from=dev /dist .
# Service listens on port 8080.
EXPOSE 8080
# Run product service container when the container starts
ENTRYPOINT ["./main"]



