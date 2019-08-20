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

resetPathFun(){
    echoFun "reset GOPATH:" title
    export GOPATH=`pwd`
    echoFun "now GOPATH: `echo $GOPATH`" tip

    echoFun "add PATH:" title
    export PATH="$GOROOT/bin:$GOPATH/bin:$PATH"
    echoFun "now PATH: `echo $PATH`" tip
}

initFun(){
    appfile="`pwd`/src/code/app/cons/app.go"
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

    freshConfigFile="`pwd`/src/code/fresh.conf"
    if [[ ! -f "$freshConfigFile" ]];then
        echoFun "fresh config file [$freshConfigFile] is not exist" err
        exit 1
    fi
    freshName=fresh-${name}
    echoFun "freshName: ${freshName}" tip
}

statusFun(){
    initFun
    sleep 2s

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
    sleep 2s

    resetPathFun
    sleep 2s

    echoFun "get glide:" title
    cd ./src
    if [[ ! -f "../bin/glide" ]];then
        mkdir -p ./github.com/Masterminds
        cd ./github.com/Masterminds
        git clone -b v0.13.2 https://github.com/Masterminds/glide.git
        cd ../../
        go install -v -x github.com/Masterminds/glide
        if [[ ! -f "../bin/glide" ]];then
            echoFun "get glide failed" err
            exit 1
        fi
        if [[ ! -x "../bin/glide" ]];then
            chmod u+x "../bin/glide"
        fi
        echoFun "get glide succeed" ok
    else
        echoFun "glide already getted" tip
    fi
    sleep 2s

    echoFun "glide install:" title
    cd ./code
    if [[ ! -f "./glide.lock" && ! -f "./glide.yaml" ]];then
        echoFun "glide lock or yaml file is not exist" err
        exit 1
    fi
    glide install
    echoFun "glide is installed" ok
    sleep 2s

    cd ../
    echoFun "get fresh:" title
    if [[ ! -f "../bin/$freshName" ]];then
        if [[ ! -d "./golang.org/x/sys" ]];then
            mkdir -p ./golang.org/x/sys
        fi
        /bin/cp -rf ./code/vendor/golang.org/x/sys ./golang.org/x
        go get -v -x github.com/pilu/fresh
        if [[ ! -f "../bin/fresh" ]];then
            echoFun "get fresh failed" err
            exit 1
        fi

        /bin/cp -rf ../bin/fresh ../bin/${freshName}
        rm -rf ../bin/fresh

        if [[ ! -x "../bin/$freshName" ]];then
            chmod u+x "../bin/$freshName"
        fi
        echoFun "get fresh succeed" ok
    else
        echoFun "fresh already getted" tip
    fi
    rm -rf ./github.com
    rm -rf ./golang.org
}

restartFun(){
    initFun
    sleep 2s

    resetPathFun
    sleep 2s

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

    cd ./src/code
    nohup ../../bin/${freshName} -c fresh.conf >> ${logfile} 2>&1 &

    echoFun "fresh is restarted" ok
}

glideUpdateFun(){
    initFun
    sleep 2s

    resetPathFun
    sleep 2s

    echo "glide update:" title
    cd ./src
    if [[ ! -f "../bin/glide" ]];then
        echoFun "glide is not exist" err
        exit 1
    fi
    cd ./code
    if [[ ! -f "./glide.lock" && ! -f "./glide.yaml" ]];then
        echoFun "glide lock or yaml file is not exist" err
        exit 1
    fi
    glide update
    echoFun "glide update finish" ok
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
        update)
            glideUpdateFun
        ;;
        stop)
            initFun
            sleep 2s
            stopFun ${freshName} ${httpServerAddr}
        ;;
        restart)
            restartFun
        ;;
        *)
            helpFun
        ;;
esac
exit 0
