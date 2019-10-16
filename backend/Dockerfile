FROM golang:1.12
ENV GO111MODULE on

WORKDIR $GOPATH/src/backend
COPY . .
# RUN bin/goose -dir migrations/ mysql "root:123456@tcp(10.1.7.239:3306)/blog" up
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./backend"]
