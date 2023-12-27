// Code generated by hertz generator.

package main

import (
	"context"
	"fmt"
	"github.com/benxinm/tiktok/cmd/api/biz/rpc"
	"github.com/benxinm/tiktok/pkg/myerrors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzUtils "github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	sentinel "github.com/hertz-contrib/opensergo/sentinel/adapter"
)

func Init() {
	rpc.Init()

}

func main() {
	Init()

	r := server.New(
		server.WithHostPorts("127.0.0.1:8080"),
		server.WithHandleMethodNotAllowed(true),
		server.WithMaxRequestBodySize(1<<31),
	)
	r.Use(recovery.Recovery(recovery.WithRecoveryHandler(recoveryHandler)))
	// Sentinel 流量治理
	r.Use(sentinel.SentinelServerMiddleware(
		sentinel.WithServerResourceExtractor(func(c context.Context, ctx *app.RequestContext) string {
			return "server_test"
		}),
		sentinel.WithServerBlockFallback(func(ctx context.Context, c *app.RequestContext) {
			hlog.CtxInfof(ctx, "frequent requests have been rejected by the gateway. clientIP: %v\n", c.ClientIP())
			c.AbortWithStatusJSON(400, hertzUtils.H{
				"status_msg":  "too many request; the quota used up",
				"status_code": -1,
			})
		}),
	))

	register(r)
	r.Spin()
}

func recoveryHandler(ctx context.Context, c *app.RequestContext, err interface{}, stack []byte) {
	hlog.CtxInfof(ctx, "[Recovery] InternalServiceError err=%v\n stack=%s\n", err, stack)
	c.JSON(consts.StatusInternalServerError, map[string]interface{}{
		"code":    myerrors.UnknownErrorCode,
		"message": fmt.Sprintf("[Recovery] err=%v\nstack=%s", err, stack),
	})
}