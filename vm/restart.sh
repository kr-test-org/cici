#!/bin/bash

set -e

backend_port=20219

check_port() {
        echo "正在检测端口。。。"
        netstat -tlpn | grep $backend_port
}

if [ -f "cici" ];then
    rm cici
fi

if check_port;then
    kill -9 $(lsof -i tcp:$backend_port -t)
fi

tar xvf $1

nohup ./cici > /dev/null 2> /dev/null  &

rm $1
