FROM golang:1.23.4-alpine AS builder

WORKDIR /shoppe_be_go

COPY . .

RUN go mod download

RUN go build -o shoppe_be_go ./cmd/server

FROM scratch

COPY ./config /config

COPY --from=builder /shoppe_be_go /

EXPOSE 8080

ENTRYPOINT [ "/shoppe_be_go", "config/local.yaml" ]