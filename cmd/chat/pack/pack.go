package pack

import (
	"errors"
	"github.com/benxinm/tiktok/kitex_gen/chat"
	"github.com/benxinm/tiktok/pkg/myerrors"
)

func MakeBaseResp(err error) *chat.BaseResp {
	if err == nil {
		return baseResp(myerrors.Success)
	}
	e := myerrors.MyError{}
	if errors.As(err, &e) {
		return baseResp(e)
	}
	s := myerrors.ServiceError.AddMessage(err.Error())
	return baseResp(s)
}

func baseResp(err myerrors.MyError) *chat.BaseResp {
	return &chat.BaseResp{
		Code: err.ErrorCode,
		Msg:  err.ErrorMsg,
	}
}
