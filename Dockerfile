FROM golang:1.20-buster AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./

RUN go build -o yaml2env

FROM gcr.io/distroless/base-debian11 AS app

WORKDIR /app

COPY --from=builder /app/yaml2env /app/yaml2env

ENTRYPOINT ["/app/yaml2env"]
