# https://blog.golang.org/docker
# Start from a Debian image with the latest version of Go installed
############################
# STEP 1 build executable binary
############################
# golang alpine 1.17
FROM golang:1.17-alpine as builder
# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
# Else you will get error => local error: tls: bad record MAC
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates
# Create appuser
ENV USER=appuser
ENV UID=10001
# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
# RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates
RUN adduser \
   --disabled-password \
   --gecos "" \
   --home "/nonexistent" \
   --shell "/sbin/nologin" \
   --no-create-home \
   --uid "${UID}" \
   "${USER}"
WORKDIR /app
COPY . .
# Build the binary
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o app main.go
############################
# STEP 2 build a small image
############################
FROM scratch

WORKDIR /

# Import from builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
# Copy our static executable
COPY --from=builder /app/app .
# Use an unprivileged user.
USER appuser:appuser

# Run the outyet command by default when the container starts.
# ENTRYPOINT /go/bin/outyet
# Document that the service listens on port 8080.
EXPOSE 8080

# Run the hello binary.
ENTRYPOINT ["/app"]
