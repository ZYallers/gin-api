#!/bin/bash

echo "NowDir: `pwd`"

echo "Reset gopath..."
export GOPATH=`pwd`/../
echo "Now gopath: `echo $GOPATH`"
echo "Now path: `echo $PATH`"

if [ ! -f "../bin/govendor" ];then
# 先下载 govendor
echo "Go get govendor..."
go get github.com/kardianos/govendor
echo "Get govendor finished"
sleep 1s

# 初始化 govendor
echo "Govendor init..."
govendor init
govendor add +e
echo "Govendor init finished"
sleep 1s
else
echo "Govendor is inited"
fi

if [ ! -f "../bin/fresh" ];then
# 热启动 fresh 下载
echo "Go get fresh..."
go get github.com/pilu/fresh
echo "Go get fresh finished"
sleep 1s
else
echo "Fresh is getted"
fi

# govendor 同步
echo "Govendor sync..."
govendor sync
echo "Govendor sync finished"
sleep 1s

# fresh 启动
cd ./application
echo "NowDir: `pwd`"

if [ ! -f "fresh.conf" ];then
echo "fresh.conf not exists!"
exit 1
fi

echo "server starting... "
fresh -c fresh.conf > ./log/api.gin.com.log 2>&1