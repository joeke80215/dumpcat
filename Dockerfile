# compile
#  docker run --rm -v $PWD:/home/GOPATH/src/github.com/joeke80215/dumpcat -e GOPATH=/home/GOPATH dumpcat-compile

FROM golang:latest

MAINTAINER joeke80215 <aejaejaej80215@gmail.com>

RUN apt-get update
RUN apt-get install -y libpcap-dev

WORKDIR /home/GOPATH/src/github.com/joeke80215/dumpcat

ENTRYPOINT [ "bash","-c" ]

CMD [ "go build -v" ]