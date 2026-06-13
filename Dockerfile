# Step 1: Build stage
FROM golang:1.22-alpine AS builder
RUN apk add --no-cache git ca-certificates
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Step 2: Final stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/main .
COPY --from=builder /app /root/

EXPOSE 8080
CMD ["./main"]
