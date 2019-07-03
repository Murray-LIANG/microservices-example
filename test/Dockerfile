ARG GO_VERSION=1.12

# Build stage.
FROM golang:${GO_VERSION} AS builder

WORKDIR /src

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o ./test .


# Production stage.
FROM alpine as prod

WORKDIR /root

COPY --from=builder /src .

EXPOSE 8080

# Use ENTRYPOINT in production.
CMD ["./test"]