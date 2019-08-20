#!/bin/bash

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
    echoFun "    status [httpAddr(监听IP:Port)]          查看服务状态" tip
    echoFun "    sync                                    同步服务vendor资源" tip
    echoFun "    build [branch(分支)]                    编译生成服务程序" tip
    echoFun "    reload [httpAddr(监听IP:Port)]          平滑重启服务" tip
    echoFun "    quit [httpAddr(监听IP:Port)]            停止服务" tip
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

    if [[ "$1" == "" ]];then
        httpServerAddr=`cat ${appfile}|grep "HttpServerDefaultAddr"|awk -F '"' '{print $2}'`
        echoFun "httpServerAddr[httpServerDefaultAddr]: $httpServerAddr" tip
    else
        httpServerAddr=$1
        echoFun "httpServerAddr[shellArgs]: $httpServerAddr" tip
    fi
    if [[ "$httpServerAddr" == "" ]];then
        echoFun "httpServerAddr is empty" err
        exit 1
    fi

    logDir=`cat ${appfile}|grep "LogDir"|awk -F '"' '{print $2}'`
    if [[ "$logDir" == "" ]];then
        echoFun "logDir is null" err
        exit 1
    fi
    echoFun "logDir: $logDir" tip
}

statusFun(){
    initFun $1
    sleep 2s

    echoFun "ps process:" title
    if [[ `pgrep ${name}|wc -l` -gt 0 ]];then
        ps -p $(pgrep ${name}|sed ':t;N;s/\n/,/;b t') -o user,pid,ppid,%cpu,%mem,vsz,rss,tty,stat,start,time,command
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
        rm -rf ./github.com
        rm -rf ./golang.org
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

buildFun(){
    echoFun "git pull:" title
    branch=$1
    if [[ "$branch" == "" ]];then
        echoFun "branch of the build is empty" err
        exit 1
    fi
    if [[ "$branch" == "local" ]];then
        echoFun "ignore git pull, direct build by local" tip
    else
        git remote update origin --prune # 更新远程分支列表
        git checkout ${branch} # 切换分支
        git pull # 拉取最新版本
        echoFun "git pull [$branch] finish" ok
    fi
    sleep 2s

    initFun
    sleep 2s

    resetPathFun
    sleep 2s

    echoFun "build runner:" title

    go build -v -x -o ./bin/${name}_tmp ./src/code/main.go
    sleep 2s

    if [[ ! -f "./bin/${name}_tmp" ]];then
        echoFun "build runner failed" err
        exit 1
    fi

    /bin/cp -rf ./bin/${name}_tmp ./bin/${name}
    rm ./bin/${name}_tmp
    echoFun "build runner [`pwd`/bin/$name] succeed" ok
}

reloadFun(){
    initFun $1
    sleep 2s

    echoFun "runner reloading:" title

    if [[ ! -f "./bin/$name" ]];then
        echoFun "runner [`pwd`/bin/$name] is not exist" err
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
    echoFun "logfile: $logfile" tip

    if [[ ! -x "./bin/$name" ]];then
        chmod u+x ./bin/${name}
    fi

    quitFun ${httpServerAddr}

    export GIN_MODE=release
    nohup ./bin/${name} --http.addr=${httpServerAddr} >> ${logfile} 2>&1 &
    echoFun "runner [$httpServerAddr] is reloaded, pid: `echo $!`" ok
}

quitFun(){
    addr=$1
    port=`echo ${addr}|awk -F ':' '{print $2}'`

    pid=""
    if [[ `lsof -i tcp:${port}|grep LISTEN|wc -l` -gt 0 ]];then
        pid=`lsof -i tcp:${port}|grep LISTEN|awk '{print $2}'`
        echoFun "tcp [$addr] address already in listening, pid: $pid" tip
        kill ${pid}
        echoFun "pid[$pid] killing" ok
    fi

    if [[ "$pid" != "" ]];then
        while true;
        do
            if [[ `lsof -i tcp:${port}|grep LISTEN|wc -l` -gt 0 ]];then
                echoFun "killing pid[$pid] ..." tip
                sleep 1s;
            else
                echoFun "quit finish" ok
                break;
            fi
        done
    else
        echoFun "no pid in running" tip
    fi
}

cmd=$1
arg1=$2
case ${cmd} in
        status)
            statusFun ${arg1}
        ;;
        sync)
            syncFun
        ;;
        update)
            glideUpdateFun
        ;;
        build)
            buildFun ${arg1}
        ;;
        quit)
            initFun ${arg1}
            sleep 2s
            quitFun ${httpServerAddr}
        ;;
        reload)
            reloadFun ${arg1}
        ;;
        *)
            helpFun
        ;;
esac
exit 0
