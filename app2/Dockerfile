FROM golang:alpine AS builder
WORKDIR /app2
COPY . .
RUN go build -o app ./cmd/main.go

FROM alpine:latest
WORKDIR /app2
COPY --from=builder /app2/ .
CMD ["/app2/app"]

# FROM golang:1.19-alpine as builder
# WORKDIR /user_service
# COPY . .
# RUN go build -o user_svc ./cmd/main.go

# FROM alpine:3.16
# WORKDIR /user_service
# COPY --from=builder /user_service/user_svc .
# EXPOSE 8888
# CMD ["/user_service/user_svc"]