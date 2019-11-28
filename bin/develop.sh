#!/bin/bash

name=""
httpServerAddr=""
logDir=""
freshName=""

echoFun(){
    str=$1
    color=$2
    case ${color} in
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
    echoFun "    status                                  查看服务状态" tip
    echoFun "    sync                                    同步服务vendor资源" tip
    echoFun "    restart                                 重启服务" tip
    echoFun "    stop                                    终止服务" tip
    echoFun "    help                                    查看命令的帮助信息" tip
    echoFun "有关某个操作的详细信息，请使用 help 命令查看" tip
    exit 0
}

initFun(){
    appfile="`pwd`/src/config/app.go"
    if [[ ! -f "$appfile" ]];then
        echoFun "file [$appfile] is not exist" err
        exit 1
    fi

    name=`cat ${appfile}|grep "Name"|awk -F '"' '{print $2}'`
    if [[ "$name" == "" ]];then
        echoFun "name is null" err
        exit 1
    fi
    echoFun "name: $name" tip

    httpServerAddr=`cat $appfile|grep "HttpServerDefaultAddr"|awk -F '"' '{print $2}'`
    if [[ "$httpServerAddr" == "" ]];then
        echoFun "httpServerAddr is empty" err
        exit 1
    fi
    echoFun "httpServerAddr: $httpServerAddr" tip

    logDir=`cat ${appfile}|grep "LogDir"|awk -F '"' '{print $2}'`
    if [[ "$logDir" == "" ]];then
        echoFun "logDir is null" err
        exit 1
    fi
    echoFun "logDir: $logDir" tip

    freshConfigFile="`pwd`/src/fresh.conf"
    if [[ ! -f "$freshConfigFile" ]];then
        echoFun "fresh config file [$freshConfigFile] is not exist" err
        exit 1
    fi
    freshName=fresh-${name}
    echoFun "freshName: ${freshName}" tip
}

statusFun(){
    initFun

    echoFun "ps process:" title
    if [[ `pgrep ${freshName}|wc -l` -gt 0 ]];then
        ps -p $(pgrep ${freshName}|sed ':t;N;s/\n/,/;b t') -o user,pid,ppid,%cpu,%mem,vsz,rss,tty,stat,start,time,command
    fi
    echoFun "lsof process:" title
    port=`echo ${httpServerAddr}|awk -F ':' '{print $2}'`
    lsof -i:${port}
}

syncFun(){
    initFun
    cd ./src
    export GOPROXY=https://goproxy.cn

    echoFun "get fresh:" title
    if [[ ! -f "../bin/$freshName" ]];then
        mkdir -p ./github.com/gravityblast
        cd ./github.com/gravityblast
        git clone -b master https://github.com/gravityblast/fresh.git
        cd ../../
        go build -x -o ../bin/${freshName} ./github.com/gravityblast/fresh/main.go
        if [[ ! -f "../bin/$freshName" ]];then
            echoFun "build fresh failed" err
            exit 1
        fi
        rm -rf ./github.com
        echoFun "get fresh finished" ok
    else
        echoFun "fresh is getted" tip
    fi

    echoFun "go mod vendor:" title
    if [[ ! -f "./go.mod" ]];then
        go mod init src
    fi
    go mod tidy
    rm -rf ./vendor
    go mod vendor
    echoFun "go mod vendor finished" ok
}

restartFun(){
    initFun

    echoFun "fresh restarting:" title

    if [[ ! -f "./bin/$freshName" ]];then
        echoFun "fresh [$freshName] is not exist" err
        exit 1
    fi

    if [[ ! -d "$logDir" ]];then
        echoFun "logDir [$logDir] is not exist" err
        exit 1
    fi

    logfile=${logDir}/${name}.log
    if [[ ! -f "$logfile" ]];then
        touch ${logfile}
    fi
    echoFun "log file: $logfile" tip

    if [[ ! -x "./bin/$freshName" ]];then
        chmod u+x ./bin/${freshName}
    fi

    stopFun ${freshName} ${httpServerAddr}

    cd ./src
    export GIN_MODE=release
    export GIN_DEBUG_STACK=on
    nohup ../bin/${freshName} -c fresh.conf >> ${logfile} 2>&1 &

    echoFun "fresh is restarted" ok
}

stopFun(){
    cmd=$1
    addr=$2
    port=`echo ${addr}|awk -F ':' '{print $2}'`

    # 删除之前启动的fresh进程
    if [[ `pgrep ${cmd}|wc -l` -gt 0 ]];then
        echoFun "$cmd already in running" tip
        pidSet=$(pgrep ${cmd})
        for pid in ${pidSet}
        do
            kill ${pid}
            echoFun "pid [$pid] is killed" ok
        done
    fi

    # 删除之前程序占用的端口进程
    if [[ `lsof -i tcp:${port}|grep LISTEN|wc -l` -gt 0 ]];then
        echoFun "tcp [$addr] address already in listening" tip
        pidSet=$(lsof -i tcp:${port}|grep LISTEN|awk '{print $2}')
        for pid in ${pidSet}
        do
            kill ${pid}
            echoFun "pid [$pid] is killed" ok
        done
    fi

    echoFun "stop finished" ok
}

case $1 in
        status)
            statusFun
        ;;
        sync)
            syncFun
        ;;
        stop)
            initFun
            stopFun ${freshName} ${httpServerAddr}
        ;;
        restart)
            restartFun
        ;;
        *)
            helpFun
        ;;
esac
