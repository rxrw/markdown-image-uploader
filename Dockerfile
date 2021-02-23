FROM golang:1.15.7-alpine AS BUILDER

WORKDIR /app

ADD . ./

RUN go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.io,direct && \
    go build main.go

COPY /app/main /bin/uploader
RUN chmod +x /bin/uploader
