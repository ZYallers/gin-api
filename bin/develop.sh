#!/bin/bash

echo -e "\033[33m NowDir: `pwd` \033[0m"
sleep 1s

echo -e "\033[42;34m Reset GOPATH: \033[0m"
export GOPATH=`pwd`
export PATH="$GOROOT/bin:$GOPATH/bin:$PATH"
echo -e "\033[33m Now GOPATH: `echo $GOPATH` \033[0m"
sleep 2s

cd ./src
echo -e "\033[33m NowDir: `pwd` \033[0m"
sleep 1s

echo -e "\033[42;34m Govendor Get And Init: \033[0m"
if [ ! -f "../bin/govendor" ];then
    # 先下载 govendor
    go get github.com/kardianos/govendor
    echo -e "\033[32m  Govendor Getted \033[0m"
    sleep 2s

    # 初始化 govendor
    govendor init
    govendor add +e
    echo -e "\033[32m Govendor Inited \033[0m"
    sleep 2s
else
    echo -e "\033[32m Govendor Inited \033[0m"
    sleep 2s
fi

echo -e "\033[42;34m Fresh Get: \033[0m"
if [ ! -f "../bin/fresh" ];then
    # 热启动 fresh 下载
    go get github.com/pilu/fresh
    echo -e "\033[32m Fresh Getted \033[0m"
    sleep 2s
else
    echo -e "\033[32m Fresh Getted \033[0m"
    sleep 2s
fi

# govendor 同步
echo -e "\033[42;34m Govendor Sync: \033[0m"
govendor sync
echo -e "\033[32m Govendor Sync Finished \033[0m"
sleep 2s

# fresh 启动
cd ./application
echo -e "\033[33m NowDir: `pwd` \033[0m"
sleep 1s

### 需要配置 Begin ###
confile=fresh_ginapi.conf
logfile=./log/api.gin.com.log
listenAddr="127.0.0.1:8087"
### 需要配置 End ###

if [ ! -f "$confile" ];then
    echo -e "\033[31m 'Fresh.conf' Not Exists \033[0m"
    exit 1
fi

echo -e "\033[42;34m Server Running: \033[0m"
# 删除之前启动的fresh进程
if [ `ps aux|grep "fresh -c $confile"|grep -v grep|wc -l` -gt 0 ];then
    pid=`ps aux|grep "fresh -c $confile"|grep -v grep|awk '{print $2}'`
    echo -e "\033[33m Fresh -c $confile already in runned, runPid: $pid \033[0m"
    kill -9 $pid
    echo -e "\033[32m Pid: $pid killed \033[0m"
fi
sleep 2s

# 删除之前程序占用的端口进程
if [ `lsof -n -P|grep "$listenAddr"|grep LISTEN|wc -l` -gt 0 ];then
    pid=`lsof -n -P|grep "$listenAddr"|grep LISTEN|awk '{print $2}'`
    echo -e "\033[33m Listen tcp $listenAddr address already in use, usePid: $pid \033[0m"
    kill -9 $pid
    echo -e "\033[32m Pid: $pid killed \033[0m"
fi
sleep 2s

# 重新启动fresh守护进程
nohup fresh -c $confile > $logfile 2>&1 &
echo -e "\033[32m Server Started \033[0m"
### End ###
