#!/bin/bash

### 需要配置 Begin ###
listenPort=":8087"
confile=fresh_ginapi.conf
logfile=./log/api.gin.com.log
### 需要配置 End ###

echoFun(){
    str=$1
    color=$2
    case $color in
        ok)
            echo -e "\033[32m $str \033[0m"
        ;;
        err)
            echo -e "\033[31m $str \033[0m"
        ;;
        tip)
            echo -e "\033[33m $str \033[0m"
        ;;
        title)
            echo -e "\033[42;34m $str \033[0m"
        ;;
        *)
            echo "$str"
        ;;
    esac
}

helpFun(){
    echoFun "操作:" title
    echoFun "    status                                  查看Server状态" tip
    echoFun "    restart                                 重载服务" tip
    echoFun "    stop                                    热重载服务" tip
    echoFun "    help                                    查看命令的帮助信息" tip
    echoFun "有关某个操作的详细信息，请使用 help 命令查看" tip
    exit 0
}

statusFun(){
    echoFun "Fresh process:" title
    ps auxw|head -1;ps auxw|grep "fresh -c $confile"|grep -v grep
    sleep 2s
    echoFun "Fresh runner process:" title
    lsof -i$listenPort
}

restartFun(){
    echoFun "Now dir: `pwd`" tip
    sleep 1s

    echoFun "Reset GOPATH:" title
    export GOPATH=`pwd`
    export PATH="$GOROOT/bin:$GOPATH/bin:$PATH"
    echoFun "Now GOPATH: `echo $GOPATH`" tip
    sleep 1s

    cd ./src
    echoFun "Now dir: `pwd`" tip
    sleep 1s

    echoFun "Govendor get and init:" title
    if [ ! -f "../bin/govendor" ];then
        # 先下载 govendor
        go get github.com/kardianos/govendor
        echoFun "Govendor is getted" ok
        sleep 1s

        # 初始化 govendor
        govendor init
        govendor add +e
        echoFun "Govendor is inited" ok
        sleep 1s
    else
        echoFun "Govendor is inited" tip
        sleep 1s
    fi

    echoFun "Fresh get:" title
    if [ ! -f "../bin/fresh" ];then
        # 热启动 fresh 下载
        go get github.com/pilu/fresh
        echoFun "Fresh is getted" ok
        sleep 1s
    else
        echoFun "Fresh is getted" tip
        sleep 1s
    fi

    # govendor 同步
    echoFun "Govendor sync:" title
    govendor sync
    echoFun "Govendor is synced" ok
    sleep 1s

    # fresh 启动
    cd ./application
    echoFun "Now dir: `pwd`" tip
    sleep 1s

    if [ ! -f "$confile" ];then
        echo " File [$confile] is not exists" err
        exit 1
    fi

    echoFun "Server running:" title
    # 删除之前启动的fresh进程
    if [ `ps aux|grep "fresh -c $confile"|grep -v grep|wc -l` -gt 0 ];then
        pid=`ps aux|grep "fresh -c $confile"|grep -v grep|awk '{print $2}'`
        echoFun "Fresh [$confile] already in runned, Pid: $pid" tip
        kill $pid
        echo "Pid [$pid] is killed" ok
    fi
    sleep 1s

    # 删除之前程序占用的端口进程
    if [ `lsof -n -P|grep "$listenPort"|grep LISTEN|wc -l` -gt 0 ];then
        pid=`lsof -n -P|grep "$listenPort"|grep LISTEN|awk '{print $2}'`
        echoFun "Listen tcp [$listenPort] port already in use, Pid: $pid" tip
        kill $pid
        echoFun "Pid [$pid] is killed" ok
    fi
    sleep 1s

    # 重新启动fresh守护进程
    nohup fresh -c $confile > $logfile 2>&1 &
    echo "Server is started" ok
}

stopFun(){
    # 删除之前启动的fresh进程
    if [ `ps aux|grep "fresh -c $confile"|grep -v grep|wc -l` -gt 0 ];then
        pid=`ps aux|grep "fresh -c $confile"|grep -v grep|awk '{print $2}'`
        echoFun "Fresh [$confile] already in runned, Pid: $pid" tip
        kill $pid
        echoFun "Pid [$pid] is killed" ok
        sleep 1s
    fi
    # 删除之前程序占用的端口进程
    if [ `lsof -n -P|grep "$listenPort"|grep LISTEN|wc -l` -gt 0 ];then
        pid=`lsof -n -P|grep "$listenPort"|grep LISTEN|awk '{print $2}'`
        echoFun "Listen tcp [$listenPort] port already in use, Pid: $pid" tip
        kill $pid
        echoFun "Pid [$pid] is killed" ok
        sleep 1s
    fi
    echoFun "Server is stoped" ok
}

case $1 in
        status)
            statusFun
        ;;
        stop)
            stopFun
        ;;
        restart)
            restartFun
        ;;
        *)
            helpFun
        ;;
esac
exit 0
### End ###
