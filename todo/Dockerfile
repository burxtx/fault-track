FROM golang

RUN mkdir -p /go/src/github.com/burxtx/fault

ADD . /go/src/github.com/burxtx/fault

RUN curl https://glide.sh/get | sh
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

RUN cd /go/src/github.com/burxtx/fault && glide install

ENTRYPOINT  watcher -run github.com/burxtx/fault/todo/cmd -watch github.com/burxtx/fault/todo
