package handler

import (
	"github.com/arl/statsviz"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"time"
)

// Work loops forever, generating a bunch of allocations of various sizes
// in order to force the garbage collector to work.
func work() {
	m := map[string][]byte{}
	for {
		b := make([]byte, 512+rand.Intn(16*1024))
		m[strconv.Itoa(len(m)%(10*100))] = b

		if len(m)%(10*100) == 0 {
			m = make(map[string][]byte)
		}

		time.Sleep(10 * time.Millisecond)
	}
}

func StatsHandler(ctx *gin.Context) {
	if ctx.Param("filepath") == "/ws" {
		statsviz.Ws(ctx.Writer, ctx.Request)
		return
	}
	statsviz.IndexAtRoot("/statsviz").ServeHTTP(ctx.Writer, ctx.Request)
}
