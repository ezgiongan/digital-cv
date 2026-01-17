# ------------------- Builder Stage -------------------
FROM golang:1.22-alpine AS builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod ./
RUN go mod download || true

COPY . .

RUN CGO_ENABLED=0 \
    GOOS=linux \
    go build -trimpath \
    -ldflags="-s -w" \
    -o /app/api \
    ./cmd/api

# ------------------- Final Stage -------------------
FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/api .

ENV TZ=Europe/Istanbul

EXPOSE 8080

CMD ["./api"]

