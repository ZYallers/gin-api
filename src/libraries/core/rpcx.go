// @Title rpcx
// @Description rpcx
// @Author cloud 2021/11/10 上午11:04
// @Software GoLand
package core

import (
	"context"
	client2 "github.com/rpcxio/rpcx-etcd/client"
	"github.com/smallnest/rpcx/client"
	xlog "github.com/smallnest/rpcx/log"
	"github.com/smallnest/rpcx/protocol"
	"github.com/smallnest/rpcx/share"
	"os"
	"src/config/env"
	"src/libraries/helper"
	"src/libraries/logger"
	"strings"
	"sync"
	"time"
)

var (
	etcdBasePath = "/hxs/rpcx/" + env.App.Env
	etcdAddr     []string
	xClientMap   sync.Map
	failMode     = client.Failover
	selectMode   = client.RoundRobin
	clientOption = client.Option{
		Retries:            3,
		RPCPath:            share.DefaultRPCPath,
		ConnectTimeout:     time.Second,
		SerializeType:      protocol.MsgPack,
		CompressType:       protocol.None,
		BackupLatency:      10 * time.Millisecond,
		TCPKeepAlivePeriod: time.Minute, // if it is zero we don't set keepalive
		IdleTimeout:        time.Minute, // ReadTimeout sets max idle time for underlying net.Conns
		GenBreaker: func() client.Breaker {
			// if failed 10 times, return error immediately, and will try to connect after 60 seconds
			return client.NewConsecCircuitBreaker(10, 60*time.Second)
		},
	}
)

func init() {
	defer helper.SafeDefer()
	switch env.App.Env {
	case env.DevMode:
		failMode = client.Failfast
		selectMode = client.RandomSelect
		etcdAddr = []string{"127.0.0.1:2379"}
		hostname, _ := os.Hostname()
		hostname = strings.ToLower(hostname)
		if hostname != "ali-pre-001" {
			etcdBasePath = strings.Replace(etcdBasePath, env.App.Env, "developer@"+hostname, 1)
			etcdAddr = []string{"121.41.83.91:2379"}
		}
	case env.GrayMode:
		etcdAddr = []string{"127.0.0.1:2379"}
	case env.ProdMode:
		etcdAddr = []string{"10.81.68.208:2379", "10.81.69.65:2379", "10.81.164.174:2379"}
	}
	if len(etcdAddr) == 0 {
		panic("etcd address is empty")
	}
	xlog.SetLogger(logger.NewRPCXLogger())
}

func ForwardRPCXService(service, method string, args map[string]interface{}) (interface{}, error) {
	share.Trace = false
	if val, ok := args["trace"]; ok && val.(string) == "on" {
		share.Trace = true
		xlog.Infof("env: %s, etcd: %s->%+v", env.App.Env, etcdBasePath, etcdAddr)
	}

	var xClient client.XClient
	if val, ok := xClientMap.Load(service); ok {
		xClient = val.(client.XClient)
	} else {
		d, _ := client2.NewEtcdV3Discovery(etcdBasePath, service, etcdAddr, false, nil)
		xClient = client.NewXClient(service, failMode, selectMode, d, clientOption)
		xClientMap.Store(service, xClient)
	}

	var reply interface{}
	ctx, cancel := context.WithTimeout(context.Background(), helper.DefaultHttpClientTimeout)
	defer cancel()
	return reply, xClient.Call(ctx, method, args, &reply)
}
