FROM golang:1.15.6-alpine AS BUILDER

WORKDIR /app

ADD . ./

RUN go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.io,direct && \
    go build main.go

FROM alpine:latest

COPY --from=BUILDER /app/main /bin/uploader

RUN chmod +x /bin/uploader

CMD ["/bin/main"]
