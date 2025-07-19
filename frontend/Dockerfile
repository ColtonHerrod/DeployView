# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o deployview-frontend

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/deployview-frontend .
COPY static/ static/
COPY templates/ templates/
EXPOSE 8080
CMD ["./deployview-frontend"]