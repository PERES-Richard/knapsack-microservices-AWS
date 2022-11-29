# Security best pratices from Snyk https://snyk.io/blog/containerizing-go-applications-with-docker/
# syntax=docker/dockerfile:1
FROM golang:1.17-alpine AS build

# Use a custom non root user
RUN adduser -D -h /app -u 1001 iamuser
USER iamuser

WORKDIR /app/src

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" -o /app/server .

FROM scratch

WORKDIR /app
COPY --from=build /etc/passwd /etc/passwd

COPY --from=build /app/server .

USER iamuser

EXPOSE 8080

ENTRYPOINT ["./server"]
HEALTHCHECK --interval=5s --timeout=10s --start-period=2s --retries=3 CMD curl -f http://localhost:8080/health
