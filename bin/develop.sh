#!/bin/bash

### 需要配置 Begin ###
appName=""
appHttpServerAddr=""
appLogDir=""
freshConfigFile=""
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
    echoFun "    sync                                    初始化&同步服务" tip
    echoFun "    restart                                 重载服务" tip
    echoFun "    stop                                    热重载服务" tip
    echoFun "    help                                    查看命令的帮助信息" tip
    echoFun "有关某个操作的详细信息，请使用 help 命令查看" tip
    exit 0
}

resetPathFun(){
    echoFun "Reset GOPATH:" title
    export GOPATH=`pwd`
    echoFun "Now GOPATH: `echo $GOPATH`" tip

    echoFun "Add PATH:" title
    export PATH="$GOROOT/bin:$GOPATH/bin:$PATH"
    echoFun "Now PATH: `echo $PATH`" tip
}

initFun(){
    appfile="`pwd`/src/application/app/constant/app.go"
    if [ ! -f "$appfile" ];then
        echoFun "File [$appfile] is not exist" err
        exit 1
    fi

    appName=`cat $appfile|grep "Name"|awk -F '"' '{print $2}'`
    if [ "$appName" == "" ];then
        echoFun "AppName is null" err
        exit 1
    fi
    echoFun "AppName: $appName" tip

    appHttpServerAddr=`cat $appfile|grep "HttpServerAddr"|awk -F '"' '{print $2}'`
    if [ "$appHttpServerAddr" == "" ];then
        echoFun "AppHttpServerAddr is null" err
        exit 1
    fi
    echoFun "AppHttpServerAddr: $appHttpServerAddr" tip

    appLogDir=`cat $appfile|grep "LogDir"|awk -F '"' '{print $2}'`
    if [ "$appLogDir" == "" ];then
        echoFun "AppLogDir is null" err
        exit 1
    fi
    echoFun "AppLogDir: $appLogDir" tip

    freshConfigFile="fresh_`echo "$appName"|awk -F '.' '{print $1}'`.conf"
    filePath="`pwd`/src/application/$freshConfigFile"
    if [ ! -f "$filePath" ];then
        echoFun "FreshConfigFile [$filePath] is not exist" err
        exit 1
    fi
    echoFun "FreshConfigFileDir: $filePath" tip
}

statusFun(){
    initFun
    sleep 1s

    echoFun "Fresh process:" title
    ps auxw|head -1;ps auxw|grep "fresh -c $freshConfigFile"|grep -v grep
    sleep 1s

    echoFun "Lsof process:" title
    port=`echo $appHttpServerAddr|awk -F ':' '{print $2}'`
    lsof -i tcp:$port
}

syncFun(){
    initFun
    sleep 1s

    resetPathFun
    sleep 1s

    cd ./src
    echoFun "Now dir: `pwd`" tip
    sleep 1s

    echoFun "Govendor get and init:" title
    if [ ! -f "../bin/govendor" ];then
        go get github.com/kardianos/govendor
        if [ ! -f "../bin/govendor" ];then
            echoFun "Govendor get failed" err
            exit 1
        fi
        if [ ! -x "../bin/govendor" ];then
            chmod +x "../bin/govendor"
        fi
        echoFun "Govendor is getted" ok
        sleep 1s

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
        go get github.com/pilu/fresh
        if [ ! -f "../bin/fresh" ];then
            echoFun "Fresh get failed" err
            exit 1
        fi
        if [ ! -x "../bin/fresh" ];then
            chmod +x "../bin/fresh"
        fi
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
}

restartFun(){
    initFun
    sleep 1s

    resetPathFun
    sleep 1s

    cd ./src/application
    if [ ! -d "$appLogDir" ];then
        echoFun "AppLogDir [$appLogDir] is not exist" err
        exit
    fi

    stopFun $freshConfigFile $appHttpServerAddr

    # 重新启动fresh守护进程
    nohup fresh -c $freshConfigFile > ${appLogDir}/${appName}.log 2>&1 &
}

stopFun(){
    # 删除之前启动的fresh进程
    if [ `ps aux|grep "fresh -c $1"|grep -v grep|wc -l` -gt 0 ];then
        pid=`ps aux|grep "fresh -c $1"|grep -v grep|awk '{print $2}'`
        echoFun "Fresh [$1] already in runned, Pid: $pid" tip
        kill $pid
        echoFun "Pid [$pid] is killed" ok
    fi

    # 删除之前程序占用的端口进程
    port=`echo $2|awk -F ':' '{print $2}'`
    if [ `lsof -i tcp:$port|grep LISTEN|wc -l` -gt 0 ];then
        pid=`lsof -i tcp:$port|grep LISTEN|awk '{print $2}'`
        echoFun "Listen tcp [$2] addr already in use, Pid: $pid" tip
        kill $pid
        echoFun "Pid [$pid] is killed" ok
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
            stopFun $freshConfigFile $appHttpServerAddr
            echoFun "Server is stoped" ok
        ;;
        restart)
            restartFun
            echoFun "Server is restarted" ok
        ;;
        *)
            helpFun
        ;;
esac
exit 0
### End ###
