#!/bin/bash

appConfigFile="$(pwd)/src/config/app/base.go"
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
    echoFun "  status               查看服务状态" tip
    echoFun "  sync                 同步服务vendor资源" tip
    echoFun "  build                编译生成服务程序" tip
    echoFun "  reload               平滑重启服务" tip
    echoFun "  quit                 停止服务" tip
    echoFun "  help                 查看命令的帮助信息" tip
    echoFun "有关某个操作的详细信息，请使用 help 命令查看" tip
    exit 0
}

syncFun(){
    cd ./src

    echoFun "go mod vendor:" title
    if [[ ! -f "./go.mod" ]];then
        go mod init src
    fi
    go mod tidy
    rm -rf ./vendor
    go mod vendor
    echoFun "go mod vendor finished" ok
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
    echoFun "httpServerAddr: $httpServerAddr" tip
    if [[ "$httpServerAddr" == "" ]];then
        echoFun "httpServerAddr is empty" err
        exit 1
    fi

    logDir=`cat ${appConfigFile}|grep "LogDir"|awk -F '"' '{print $2}'`
    if [[ "$logDir" == "" ]];then
        echoFun "logDir is null" err
        exit 1
    fi
    echoFun "logDir: $logDir" tip
}

statusFun(){
    echoFun "ps process:" title
    if [[ `pgrep ${name}|wc -l` -gt 0 ]];then
        ps -p $(pgrep ${name}|sed ':t;N;s/\n/,/;b t'|sed -n '1h;1!H;${g;s/\n/,/g;p;}') -o user,pid,ppid,%cpu,%mem,vsz,rss,tty,stat,start,time,command
    fi

    echoFun "lsof process:" title
    port=`echo ${httpServerAddr}|awk -F ':' '{print $2}'`
    lsof -i:${port}
}

buildFun(){
    echoFun "git pull:" title
    branch=$1
    env=$2
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

    echoFun "build runner:" title
    cd ./src
    tmpName="${name}_tmp_$(date +'%Y-%m-%d-%H-%M-%S')"
    if [[ "$dlv" ]];then
        echoFun "build with -gcflags 'all=-N -l'" tip
        if [[ "$env" == "dev" ]];then
            echoFun 'build in develop environment' tip
            CGO_ENABLED=0 go build -v -installsuffix cgo -gcflags 'all=-N -l' -i -o ../bin/${tmpName} -tags=jsoniter ./main.go
        else
            CGO_ENABLED=0 go build -a -installsuffix cgo -gcflags 'all=-N -l' -i -o ../bin/${tmpName} -tags=jsoniter ./main.go
        fi
    else
        echoFun "no extra build options" tip
        if [[ "$env" == "dev" ]];then
            echoFun 'build in develop environment' tip
            CGO_ENABLED=0 go build -v -installsuffix cgo -ldflags '-w' -i -o ../bin/${tmpName} -tags=jsoniter ./main.go
        else
            ### build编译参数参考资料：
            # 无依赖编译：https://blog.csdn.net/weixin_42506905/article/details/93135684
            # build参数详解：https://blog.csdn.net/zl1zl2zl3/article/details/83374131
            # ldflags参数：https://blog.csdn.net/javaxflinux/article/details/89177863
            CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-w' -i -o ../bin/${tmpName} -tags=jsoniter ./main.go
        fi
    fi

    cd ../
    if [[ ! -f "./bin/${tmpName}" ]];then
        echoFun "build tmp runner ($(pwd)/bin/${tmpName}) failed" err
        exit 1
    fi

    mv -f ./bin/${tmpName} ./bin/${name}
    if [[ ! -f "./bin/${name}" ]];then
        echoFun "mv tmp runner failed" err
        exit 1
    fi

    echoFun "build runner ($(pwd)/bin/${name}) finished" ok
}

sendMsg(){
    env="Env: $(echo $hxsenv)"
    app="App: $name"
    addr="Addr: $httpServerAddr"
    hostName="HostName: $(hostname)"
    sip="SystemIP: $(/sbin/ifconfig -a|grep inet|grep -v 127.0.0.1|grep -v inet6|awk '{print $2}'|tr -d "addr:"|head -1)"
    time="Time: $(date "+%Y/%m/%d %H:%M:%S")"
    token="0d10a39901249e43fe66065f55449e17df7bc788617edd5dbcaf77396668be7b"
    url="https://oapi.dingtalk.com/robot/send?access_token=$token"
    content="$1\n---------------------------\n$env\n$app\n$addr\n$hostName\n$time\n$sip"
    cnt=$(echo ${content//\"/\\\"})
    header="Content-Type: application/json"
    curl -o /dev/null -m 3 -s "$url" -H "$header" -d "{\"msgtype\": \"text\",\"text\": {\"content\":\"$cnt\"}}"
}

quitFun(){
    port=`echo ${httpServerAddr}|awk -F ':' '{print $2}'`
    counter=0
    while true;
    do
        pid=`lsof -i tcp:${port}|grep LISTEN|awk '{print $2}'`
        if [[ ${pid} -gt 0 ]];then
            if [[ ${counter} -ge 30 ]];then
                kill -9 ${pid}
                echoFun "service $name has been killed for 30s and is ready to be forcibly killed" tip
                sendMsg "service $name has been killed for 30s and is ready to be forcibly killed"
                break
            else
                kill ${pid}
                counter=$(($counter+1))
                echoFun "killing service $name($port), pid($pid), $counter tried" tip
                sleep 1s
            fi
        else
            echoFun "service $name($port) is stopped" ok
            break
        fi
    done
}

reloadFun(){
    echoFun "runner reloading:" title

    if [[ ! -f "./bin/$name" ]];then
        echoFun "runner [`pwd`/bin/$name] is not exist" err
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
    echoFun "logfile: $logfile" tip

    if [[ ! -x "./bin/$name" ]];then
        chmod u+x ./bin/${name}
    fi

    quitFun

    # 防止Jenkins默认会在Build结束后Kill掉所有的衍生进程
    export BUILD_ID=dontKillMe

    nohup ./bin/${name} -http.addr=${httpServerAddr} >> ${logfile} 2>&1 &
    echoFun "service $name($httpServerAddr) is reloaded, pid: `echo $!`" ok

    # 检查健康接口是否访问正常
    sleep 3s
    resp=`curl -m 3 -s "http://$httpServerAddr/health"`
    echoFun "curl \"http://$httpServerAddr/health\": $resp" tip
    sendMsg "curl \"http://$httpServerAddr/health\": $resp"
}

while getopts ':d' OPT; do
    case ${OPT} in
        d)
            # 配合 delve 使用, http://wiki.sys.hxsapp.net/pages/viewpage.action?pageId=21349181
            dlv=d
            shift 1;
        ;;
        ?)  #当有不认识的选项的时候arg为?
            echo "unknown argument"
            exit 1
        ;;
    esac
done

case $1 in
    status)
        initFun
        statusFun
    ;;
    build)
        initFun
        buildFun $2 $3
    ;;
    quit)
        initFun
        quitFun
    ;;
    reload)
        initFun
        reloadFun
    ;;
    sync)
        syncFun
    ;;
    *)
        helpFun
    ;;
esac
