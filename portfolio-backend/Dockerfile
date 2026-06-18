FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY portfolio-backend/go.mod portfolio-backend/go.sum ./
RUN go mod download

COPY portfolio-backend/ .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/api ./cmd/api

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /bin/api ./api
EXPOSE 8080
CMD ["./api"]
