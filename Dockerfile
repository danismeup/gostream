FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum* ./
RUN go mod download
COPY . .
RUN go build -o gostream ./cmd/server

FROM alpine:latest
RUN apk add --no-cache \
    gstreamer \
    gst-plugins-base \
    gst-plugins-good \
    gst-plugins-bad \
    gst-plugins-ugly \
    gst-libav
WORKDIR /app
COPY --from=builder /app/gostream .
COPY web/ web/
EXPOSE 8080
CMD ["./gostream"]
