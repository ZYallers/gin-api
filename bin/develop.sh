#!/bin/bash

appConfigFile="$(pwd)/src/config/app/base.go"
freshName=""
name=""
httpServerAddr=""
logDir=""

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
    echoFun "    swag_init                               swag init(同步生成接口文档)" tip
    echoFun "    help                                    查看命令的帮助信息" tip
    echoFun "有关某个操作的详细信息，请使用 help 命令查看" tip
    exit 0
}

initFun(){
    if [[ ! -f "$appConfigFile" ]];then
        echoFun "file [$appConfigFile] is not exist" err
        exit 1
    fi

    name=`cat ${appConfigFile}|grep "Name"|awk -F '"' '{print $2}'`
    if [[ "$name" == "" ]];then
        echoFun "name is null" err
        exit 1
    fi
    echoFun "name: $name" tip

    httpServerAddr=`cat ${appConfigFile}|grep "HttpServerDefaultAddr"|awk -F '"' '{print $2}'`
    if [[ "$httpServerAddr" == "" ]];then
        echoFun "httpServerAddr is empty" err
        exit 1
    fi
    echoFun "httpServerAddr: $httpServerAddr" tip

    logDir=`cat ${appConfigFile}|grep "LogDir"|awk -F '"' '{print $2}'`
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
        ps -p $(pgrep ${name}|sed ':t;N;s/\n/,/;b t'|sed -n '1h;1!H;${g;s/\n/,/g;p;}') -o user,pid,ppid,%cpu,%mem,vsz,rss,tty,stat,start,time,command
    fi
    echoFun "lsof process:" title
    port=`echo ${httpServerAddr}|awk -F ':' '{print $2}'`
    lsof -i:${port}
}

syncFun(){
    initFun

    echoFun "go get fresh:" title
    if [[ ! -f "../bin/$freshName" ]];then
        if [[ ! -f "$GOPATH/bin/fresh" ]];then
            nowPwd="$(pwd)"
            cd $GOPATH
            go get -v github.com/pilu/fresh
            if [[ ! -f "./bin/fresh" ]];then
                echoFun "go get fresh command failed" err
                exit 1
            fi
            cd ${nowPwd}
        fi
        cp -f $GOPATH/bin/fresh ./bin/${freshName}
        if [[ ! -f "./bin/${freshName}" ]];then
            echoFun "cp fresh failed" err
            exit 1
        fi
        echoFun "go get fresh finished" ok
    else
        echoFun "fresh have got" tip
    fi

    echoFun "go mod vendor:" title
    cd ./src
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

    # 日志目录
    if [[ ! -d "$logDir" ]];then
        mkdir -p ${logDir}
    fi
    if [[ ! -d "$logDir" ]];then
        echoFun "logDir [$logDir] is not exist" err
        exit 1
    fi

    # 日志文件
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

    # 防止Jenkins默认会在Build结束后Kill掉所有的衍生进程
    export BUILD_ID=dontKillMe

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

swagInitFun() {
    cd ./src
    echoFun "pwd: $(pwd)" tip

    $GOPATH/bin/swag init -d ./controller -g ../main.go -o ../doc

    if [[ -f "../doc/docs.go" ]];then
        rm -f ../doc/docs.go
    fi

    if [[ -f "../doc/swagger.yaml" ]];then
        rm -f ../doc/swagger.yaml
    fi
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
        swag_init)
           swagInitFun
        ;;
        *)
            helpFun
        ;;
esac
