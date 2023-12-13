FROM golang:1.20.11-alpine3.18 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o ordersvc

FROM alpine:latest
WORKDIR /app
COPY --from=build-stage /app/ordersvc .
# COPY config.yml ./
# COPY db ./db

EXPOSE 8080

CMD ["./ordersvc", "api"]
