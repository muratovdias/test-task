FROM golang:alpine AS builder
WORKDIR /app1
COPY . .
RUN go build -o app main.go

FROM alpine:latest
WORKDIR /app1
COPY --from=builder /app1/ .
CMD ["/app1/app"]