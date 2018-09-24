# dumpcat
[![Go Report Card](https://goreportcard.com/badge/github.com/joeke80215/dumpcat)](https://goreportcard.com/report/github.com/joeke80215/dumpcat)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/7bd35b95e8664bd5b8636efa7673c729)](https://www.codacy.com/app/joeekee/dumpcat?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=joeke80215/dumpcat&amp;utm_campaign=Badge_Grade)
[![GitHub](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/joeke80215/dumpcat/blob/master/LICENSE)
[![Build Status](https://travis-ci.org/joeke80215/dumpcat.svg?branch=master)](https://travis-ci.org/joeke80215/dumpcat)

dump network packet to database
## overview
* dump packet network to database
* monitor network packet
* analysis network packet


## prerequest
- linux: install libpcap 
  - centos 
    ```
    sudo yum install -y libpcap-devel && sufo yum -y install libpcap
    ```
  - ubuntu 
    ```
    apt-get install -y libpcap-devel && apt-get install -y libpcap
    ```
- windows: install winpcap (https://www.winpcap.org/)

## install
```
go get github.com/joeke80215/dumpcat
```
## build
```
cd $GOPATH/src/github.com/joeke80215/dumpcat
go build -v
```

## filter
BPF format (http://biot.com/capstats/bpf.html)

## database support
- elasticsearch

## logic layer support
- latency

## config
config.yaml
```
dumplist:
  {dump name}:
    device: {device name}
    bpf: {BPF filter string}
    .
    .
    .

logics:
- timeoffset
output:
- elasticsearch:
  host: {elasticsearch server host}
```

## usage example
#### create config.yaml
```
dumpList: 
  http:
    device: "enp2s0"
    bpf: "tcp port 80"
  ftp:
    device: "enp2s0"
    bpf: "tcp port 21"
  sftp:
    device: "enp2s0"
    bpf: "tcp port 22"
logics:
- latency
output:
  elasticsearch:
    host: "http://192.168.0.100:9200"
```
#### execute
```
./dumpcat -f config.yaml
```
