FROM golang:1.16-alpine3.14 as go-builder

WORKDIR /src
COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o server cmd/server/main.go
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o request cmd/request/main.go


# ================================

FROM busybox:stable AS final

COPY --from=go-builder /src/server /server
COPY --from=go-builder /src/request /bin/request

USER 0:0

ENTRYPOINT ["/server"]