FROM golang:1.17 AS builder

WORKDIR /src
COPY . .
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o /app -a -ldflags '-linkmode external -extldflags "-static"' .
RUN ls

FROM alpine
COPY --from=builder /app /app
COPY frontend/ /frontend/
EXPOSE 8091

ENTRYPOINT ["/app"]