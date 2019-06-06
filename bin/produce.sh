#!/bin/bash

echo -e "\033[33m NowDir: `pwd` \033[0m"
sleep 2s

echo -e "\033[42;34m Reset GOPATH: \033[0m"
export GOPATH=`pwd`
export PATH="$GOROOT/bin:$GOPATH/bin:$PATH"
echo -e "\033[33m Now GOPATH: `echo $GOPATH` \033[0m"
sleep 2s

### 需要配置 Begin ###
listenAddr="127.0.0.1:8087"
runnerName=produce_runner_ginapi
### 需要配置 End ###

# 先bulid生成线上跑的可执行文件
echo -e "\033[42;34m Build ProduceRunner: \033[0m"
go build -o ./bin/${runnerName}_tmp ./src/application/main.go
if [ ! -f "./bin/${runnerName}_tmp" ];then
    echo -e "\033[31m ProduceRunner Build Failed \033[0m"
    exit 1
fi

/bin/cp -rf ./bin/${runnerName}_tmp ./bin/$runnerName
rm ./bin/${runnerName}_tmp

echo -e "\033[32m Build ProduceRunner Finished \033[0m"
sleep 2s

# 关闭已在运行的Server，如果存在
if [ `ps aux|grep "$runnerName"|grep -v grep|wc -l` -gt 0 ];then
    pid=`ps aux|grep "$runnerName"|grep -v grep|awk '{print $2}'`
    echo -e "\033[33m ProduceRunner already in runned, runPid: $pid \033[0m"
    kill $pid
    echo -e "\033[32m Pid: $pid killed \033[0m"
fi

# 重新启动Server
echo -e "\033[42;34m ProduceRunner Running: \033[0m"
chmod +x ./bin/$runnerName
cd ./src/application
echo -e "\033[33m NowDir: `pwd` \033[0m"

export GIN_MODE=release
nohup ../../bin/$runnerName > /dev/null 2>&1 &
echo -e "\033[32m ProduceRunner Is Started \033[0m"
### End ###
