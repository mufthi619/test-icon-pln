FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

#RUN go test -v ./...

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w -s' -o main ./cmd/api/

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8001

CMD ["./main"]