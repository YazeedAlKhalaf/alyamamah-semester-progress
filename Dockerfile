FROM golang:1.19 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main /app/main
COPY app.env.example .

EXPOSE 8080
CMD [ "/app/main" ]