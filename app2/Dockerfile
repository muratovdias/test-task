FROM golang:alpine AS builder
WORKDIR /app2
COPY . .
RUN go build -o app ./cmd/main.go

FROM alpine:latest
WORKDIR /app2
COPY --from=builder /app2/ .
CMD ["/app2/app"]