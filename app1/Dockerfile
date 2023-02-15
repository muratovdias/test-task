FROM golang:alpine AS builder
WORKDIR /app1
COPY . .
RUN go build -o app main.go

FROM alpine:latest
WORKDIR /app1
COPY --from=builder /app1/ .
CMD ["/app1/app"]


# FROM golang:1.19-alpine as builder
# WORKDIR /salt_service
# COPY . .
# RUN go build -o salt_svc main.go

# FROM alpine:3.16
# WORKDIR /salt_service
# COPY --from=builder /salt_service/salt_svc .
# EXPOSE 8887
# CMD ["/salt_service/salt_svc"]