#!/bin/bash

### 需要配置 ###
appName=""
appHttpServerAddr=""
appLogDir=""
produceRunnerName=""
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
    echoFun "    build                                   生成ProduceRunner" tip
    echoFun "    reload                                  平滑重启服务" tip
    echoFun "    quit                                    停止服务" tip
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
    appfile="`pwd`/src/code/app/constant/app.go"
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
        echoFun "appHttpServerAddr is null" err
        exit 1
    fi
    echoFun "AppHttpServerAddr: $appHttpServerAddr" tip

    appLogDir=`cat $appfile|grep "LogDir"|awk -F '"' '{print $2}'`
    if [ "$appLogDir" == "" ];then
        echoFun "AppLogDir is null" err
        exit 1
    fi
    echoFun "AppLogDir: $appLogDir" tip

    produceRunnerName="produceRunner_`echo "$appName"|awk -F '.' '{print $1}'`"
    if [ "$produceRunnerName" == "" ];then
        echoFun "ProduceRunnerName is null" err
        exit 1
    fi
    echoFun "ProduceRunnerName: $produceRunnerName" tip
}

statusFun(){
    initFun
    sleep 2s

    echoFun "ProduceRunner process:" title
    ps auxw|head -1;ps auxw|grep "$produceRunnerName"|grep -v grep

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
    rm -rf ./github.com
    sleep 2s

    cd ./code

    echoFun "Glide install:" title
    if [[ ! -f "./glide.lock" && ! -f "./glide.yaml" ]];then
        echoFun "Glide lock or yaml file is not exist" err
        exit 1
    fi
    glide install -v
    echoFun "Glide is synced" ok
}

buildFun(){
    initFun
    sleep 2s

    resetPathFun
    sleep 2s

    echoFun "Build produceRunner:" title
    go build -v -x -o ./bin/${produceRunnerName}_tmp ./src/code/main.go
    sleep 2s
    if [ ! -f "./bin/${produceRunnerName}_tmp" ];then
        echoFun "Build produceRunner failed" err
        exit 1
    fi

    /bin/cp -rf ./bin/${produceRunnerName}_tmp ./bin/$produceRunnerName
    rm ./bin/${produceRunnerName}_tmp
    echoFun "Build produceRunner [`pwd`/bin/$produceRunnerName] is successful" ok
}

reloadFun(){
    initFun
    sleep 2s

    if [ ! -f "./bin/$produceRunnerName" ];then
        echoFun "ProduceRunner [`pwd`/bin/$produceRunnerName] is not exist" err
        exit 1
    fi

    quitFun $produceRunnerName

    echoFun "ProduceRunner running:" title
    if [ ! -x "./bin/$produceRunnerName" ];then
        chmod u+x ./bin/$produceRunnerName
    fi

    cd ./src/code
    if [ ! -d "$appLogDir" ];then
        echoFun "AppLogDir [$appLogDir] is not exist" err
        exit
    fi

    export GIN_MODE=release
    nohup ../../bin/$produceRunnerName > ${appLogDir}/${appName}.log 2>&1 &
    echoFun "ProduceRunner is reloaded" ok
}

quitFun(){
    if [ `ps aux|grep "$1"|grep -v grep|wc -l` -gt 0 ];then
        pid=`ps aux|grep "$1"|grep -v grep|awk '{print $2}'`
        echoFun "ProduceRunner [$1] already in runned, Pid: $pid" tip
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
        build)
            buildFun
        ;;
        quit)
            initFun
            sleep 2s
            quitFun $produceRunnerName
            echoFun "ProduceRunner is quited" ok
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
