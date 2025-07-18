# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o deployview-backend

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/deployview-backend .
EXPOSE 8080
CMD ["./deployview-backend"]