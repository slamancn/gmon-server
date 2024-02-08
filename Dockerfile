FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/slamancn/gmon-server
COPY . $GOPATH/src/github.com/slamancn/gmon-server
RUN go build .

EXPOSE 20006
ENTRYPOINT ["./go-gin-example"]
