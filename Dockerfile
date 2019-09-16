FROM golang:latest
WORKDIR $GOPATH/src/github.com/OrderSystem_WeiZhang
#COPY . $GOPATH/src/github.com/OrderSystem_WeiZhang
COPY main.go $GOPATH/src
#RUN  go get && go build .

#ENTRYPOINT ["/$GOPATH/OrderSystem_WeiZhang"]

EXPOSE 3000

