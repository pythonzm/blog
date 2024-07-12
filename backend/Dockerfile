FROM golang:1.22 as builder
ENV GO111MODULE on

WORKDIR $GOPATH/src/backend
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# RUN bin/goose -dir migrations/ mysql "root:123456@tcp(10.1.7.239:3306)/blog" up
RUN go build

FROM alpine:latest
RUN apk --no-cache add netcat-openbsd
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

WORKDIR /go/src
COPY --from=builder /go/src/backend .
#EXPOSE 8000
ENTRYPOINT ["./wait.sh", "mysql:3306", "--", "./backend"]
