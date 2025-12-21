FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o backend

FROM scratch

WORKDIR /app

COPY --from=builder /app/backend ./

EXPOSE 8080

ENTRYPOINT ["./backend"]
