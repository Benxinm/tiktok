package pack

import (
	"errors"
	"github.com/benxinm/tiktok/kitex_gen/video"
	"github.com/benxinm/tiktok/pkg/myerrors"
)

func MakeBaseResp(err error) *video.BaseResp {
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

func baseResp(err myerrors.MyError) *video.BaseResp {
	return &video.BaseResp{
		Code: err.ErrorCode,
		Msg:  err.ErrorMsg,
	}
}
