FROM golang:1.25-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o backend

FROM scratch AS release
COPY --from=build /app/backend /backend
EXPOSE 8080
ENTRYPOINT ["/backend"]
