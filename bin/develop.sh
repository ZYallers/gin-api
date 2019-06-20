#!/bin/bash

### 需要配置 ###
appName=""
appHttpServerAddr=""
appLogDir=""
freshConfigFile=""
### 需要配置 ###

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
    echoFun "    sync                                    同步vendor资源" tip
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
    appfile="`pwd`/src/code/app/cons/app.go"
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
    filePath="`pwd`/src/code/$freshConfigFile"
    if [ ! -f "$filePath" ];then
        echoFun "FreshConfigFile [$filePath] is not exist" err
        exit 1
    fi
    echoFun "FreshConfigFileDir: $filePath" tip
}

statusFun(){
    initFun
    sleep 2s

    echoFun "Fresh process:" title
    ps auxw|head -1;ps auxw|grep "fresh -c $freshConfigFile"|grep -v grep

    echoFun "Tcp process:" title
    port=`echo $appHttpServerAddr|awk -F ':' '{print $2}'`
    lsof -i tcp:$port
}

syncFun(){
    initFun
    sleep 2s

    resetPathFun
    sleep 2s

    cd ./src

    echoFun "Go get glide:" title
    if [ ! -f "../bin/glide" ];then
        mkdir -p ./github.com/Masterminds
        cd ./github.com/Masterminds
        git clone -b v0.13.2 https://github.com/Masterminds/glide.git
        cd ../../
        go install -v -x github.com/Masterminds/glide
        if [ ! -f "../bin/glide" ];then
            echoFun "Go get glide failed" err
            exit 1
        fi
        if [ ! -x "../bin/glide" ];then
            chmod u+x "../bin/glide"
        fi
        echoFun "Go get glide succeed" ok
    else
        echoFun "Glide is go getted" tip
    fi
    sleep 2s

    cd ./code

    echoFun "Glide install:" title
    if [[ ! -f "./glide.lock" && ! -f "./glide.yaml" ]];then
        echoFun "Glide lock or yaml file is not exist" err
        exit 1
    fi
    glide install
    echoFun "Glide is installed" ok
    sleep 2s

    cd ../
    echoFun "Go get fresh:" title
    if [ ! -f "../bin/fresh" ];then
        if [ ! -d "./golang.org/x/sys" ];then
            mkdir -p ./golang.org/x/sys
        fi
        /bin/cp -rf ./code/vendor/golang.org/x/sys ./golang.org/x
        go get -v -x github.com/pilu/fresh
        if [ ! -f "../bin/fresh" ];then
            echoFun "Go get fresh failed" err
            exit 1
        fi
        if [ ! -x "../bin/fresh" ];then
            chmod u+x "../bin/fresh"
        fi
        echoFun "Go get fresh succeed" ok
    else
        echoFun "Fresh is go getted" tip
    fi
    rm -rf ./github.com
    rm -rf ./golang.org
    echoFun "Glide is synced" ok
}

restartFun(){
    initFun
    sleep 2s

    resetPathFun
    sleep 2s

    cd ./src/code
    if [ ! -d "$appLogDir" ];then
        echoFun "AppLogDir [$appLogDir] is not exist" err
        exit
    fi

    stopFun $freshConfigFile $appHttpServerAddr

    # 重新启动fresh守护进程
    logfile=${appLogDir}/${appName}.log
    if [ ! -f "$logfile" ];then
        touch $logfile
    fi
    echoFun "Log file: $logfile" tip
    nohup fresh -c $freshConfigFile >> $logfile 2>&1 &
    echoFun "Server is restarted" ok
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
            sleep 2s
            stopFun $freshConfigFile $appHttpServerAddr
            echoFun "Server is stoped" ok
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
