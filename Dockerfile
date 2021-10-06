FROM golang:1.16-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go get .
RUN go build -o ./preguntados-game ./main.go

FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/preguntados-game .
EXPOSE 8080
ENTRYPOINT ["./preguntados-game"]