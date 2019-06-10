#!/bin/bash

### 需要配置 Begin ###
listenPort=":8087"
runnerName=produce_ginapi_runner
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
    echoFun "    build                                   生成服务Runner" tip
    echoFun "    reload                                  平滑重启服务" tip
    echoFun "    quit                                    停止服务" tip
    echoFun "    help                                    查看命令的帮助信息" tip
    echoFun "有关某个操作的详细信息，请使用 help 命令查看" tip
    exit 0
}

statusFun(){
    echoFun "Runner process:" title
    ps auxw|head -1;ps auxw|grep "$runnerName"|grep -v grep
    sleep 2s
    echoFun "Lsof process:" title
    lsof -i$listenPort
}

buildFun(){
    echoFun "Now dir: `pwd`" tip
    sleep 1s

    echoFun "Reset GOPATH:" title
    export GOPATH=`pwd`
    export PATH="$GOROOT/bin:$GOPATH/bin:$PATH"
    echoFun "Now GOPATH: `echo $GOPATH`" tip
    sleep 1s

    # 先bulid生成线上跑的可执行文件
    echoFun "Build runner:" title
    go build -o ./bin/${runnerName}_tmp ./src/application/main.go
    sleep 1s
    if [ ! -f "./bin/${runnerName}_tmp" ];then
        echoFun "Build runner failed" err
        exit 1
    fi

    /bin/cp -rf ./bin/${runnerName}_tmp ./bin/$runnerName
    rm ./bin/${runnerName}_tmp
    sleep 1s

    echoFun "Build runner [`pwd`/bin/$runnerName] is successful" ok
}

reloadFun(){
    if [ ! -f "./bin/$runnerName" ];then
        echoFun "Runner [`pwd`/bin/$runnerName] is not exist" err
        exit 1
    fi

    # 关闭已在运行的Server，如果存在
    if [ `ps aux|grep "$runnerName"|grep -v grep|wc -l` -gt 0 ];then
        pid=`ps aux|grep "$runnerName"|grep -v grep|awk '{print $2}'`
        echoFun "Runner [$runnerName] already in runned, Pid: $pid" tip
        kill $pid
        echoFun "Pid [$pid] is killed" ok
    fi

    # 重新启动Server
    echoFun "Runner running:" title
    chmod +x ./bin/$runnerName
    cd ./src/application
    echoFun "Now dir: `pwd`" tip
    export GIN_MODE=release
    nohup ../../bin/$runnerName > /dev/null 2>&1 &
    echoFun "Runner is started" ok
}

quitFun(){
    if [ `ps aux|grep "$runnerName"|grep -v grep|wc -l` -gt 0 ];then
        pid=`ps aux|grep "$runnerName"|grep -v grep|awk '{print $2}'`
        echoFun "Runner [$runnerName] already in runned, Pid: $pid" tip
        kill $pid
        echoFun "Pid [$pid] is killed" ok
    fi
    echoFun "Runner is quited" ok
}

case $1 in
        status)
            statusFun
        ;;
        build)
            buildFun
        ;;
        quit)
            quitFun
        ;;
        reload)
            reloadFun
        ;;
        *)
            helpFun
        ;;
esac
exit 0
### End ###
