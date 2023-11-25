package pack

import (
	"github.com/benxinm/tiktok/pkg/myerrors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type Response struct {
	Code int64  `json:"status_code"`
	Msg  string `json:"status_msg"`
}

func FailResponse(c *app.RequestContext, err error) {
	if err == nil {
		c.JSON(consts.StatusOK, Response{
			Code: myerrors.SuccessCode,
			Msg:  myerrors.SuccessMsg,
		})
		return
	}
	c.JSON(consts.StatusOK, Response{
		Code: -1,
		Msg:  err.Error(),
	})
}
