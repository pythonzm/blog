FROM golang:1.12
ENV GO111MODULE on

WORKDIR $GOPATH/src/backend
COPY . .
# RUN bin/goose -dir migrations/ mysql "root:123456@tcp(10.1.7.239:3306)/blog" up
RUN apt-get update && apt-get install netcat -y
RUN go build .


#EXPOSE 8000
ENTRYPOINT ["./wait.sh", "mysql:3306", "--", "./backend"]
