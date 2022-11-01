FROM golang:1.16-alpine3.14 as go-builder

WORKDIR /src
COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o zkapp main.go


# ================================

FROM alpine:3.14 AS final

COPY --from=go-builder /src/zkapp /zkapp

USER 0:0

ENTRYPOINT ["/zkapp"]