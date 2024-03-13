# Building Stage
FROM golang:alpine3.19 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o main .

WORKDIR /dist
RUN cp /build/main .

# Final Stage
FROM scratch

LABEL maintainer="Your Name <your.email@example.com>"
LABEL org.opencontainers.image.source="https://github.com/yourusername/yourrepository"

COPY --from=builder /dist/main /

EXPOSE 3000

ENTRYPOINT ["/main"]
