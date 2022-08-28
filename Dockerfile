FROM docker.io/oxio/golang-protoc
COPY go_scaffold /go/bin
RUN rm -f ~/.gitconfig
RUN go env -w GO111MODULE=on
RUN go get -u github.com/golang/protobuf/protoc-gen-go@v1.3
WORKDIR /opt/app
