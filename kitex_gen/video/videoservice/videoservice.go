// Code generated by Kitex v0.6.1. DO NOT EDIT.

package videoservice

import (
	"context"
	video "github.com/benxinm/tiktok/kitex_gen/video"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return videoServiceServiceInfo
}

var videoServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "VideoService"
	handlerType := (*video.VideoService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Feed":                 kitex.NewMethodInfo(feedHandler, newVideoServiceFeedArgs, newVideoServiceFeedResult, false),
		"PutVideo":             kitex.NewMethodInfo(putVideoHandler, newVideoServicePutVideoArgs, newVideoServicePutVideoResult, false),
		"GetFavoriteVideoInfo": kitex.NewMethodInfo(getFavoriteVideoInfoHandler, newVideoServiceGetFavoriteVideoInfoArgs, newVideoServiceGetFavoriteVideoInfoResult, false),
		"GetPublishList":       kitex.NewMethodInfo(getPublishListHandler, newVideoServiceGetPublishListArgs, newVideoServiceGetPublishListResult, false),
		"GetWorkCount":         kitex.NewMethodInfo(getWorkCountHandler, newVideoServiceGetWorkCountArgs, newVideoServiceGetWorkCountResult, false),
		"GetVideoIDByUid":      kitex.NewMethodInfo(getVideoIDByUidHandler, newVideoServiceGetVideoIDByUidArgs, newVideoServiceGetVideoIDByUidResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "video",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.6.1",
		Extra:           extra,
	}
	return svcInfo
}

func feedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceFeedArgs)
	realResult := result.(*video.VideoServiceFeedResult)
	success, err := handler.(video.VideoService).Feed(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceFeedArgs() interface{} {
	return video.NewVideoServiceFeedArgs()
}

func newVideoServiceFeedResult() interface{} {
	return video.NewVideoServiceFeedResult()
}

func putVideoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServicePutVideoArgs)
	realResult := result.(*video.VideoServicePutVideoResult)
	success, err := handler.(video.VideoService).PutVideo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServicePutVideoArgs() interface{} {
	return video.NewVideoServicePutVideoArgs()
}

func newVideoServicePutVideoResult() interface{} {
	return video.NewVideoServicePutVideoResult()
}

func getFavoriteVideoInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceGetFavoriteVideoInfoArgs)
	realResult := result.(*video.VideoServiceGetFavoriteVideoInfoResult)
	success, err := handler.(video.VideoService).GetFavoriteVideoInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceGetFavoriteVideoInfoArgs() interface{} {
	return video.NewVideoServiceGetFavoriteVideoInfoArgs()
}

func newVideoServiceGetFavoriteVideoInfoResult() interface{} {
	return video.NewVideoServiceGetFavoriteVideoInfoResult()
}

func getPublishListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceGetPublishListArgs)
	realResult := result.(*video.VideoServiceGetPublishListResult)
	success, err := handler.(video.VideoService).GetPublishList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceGetPublishListArgs() interface{} {
	return video.NewVideoServiceGetPublishListArgs()
}

func newVideoServiceGetPublishListResult() interface{} {
	return video.NewVideoServiceGetPublishListResult()
}

func getWorkCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceGetWorkCountArgs)
	realResult := result.(*video.VideoServiceGetWorkCountResult)
	success, err := handler.(video.VideoService).GetWorkCount(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceGetWorkCountArgs() interface{} {
	return video.NewVideoServiceGetWorkCountArgs()
}

func newVideoServiceGetWorkCountResult() interface{} {
	return video.NewVideoServiceGetWorkCountResult()
}

func getVideoIDByUidHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceGetVideoIDByUidArgs)
	realResult := result.(*video.VideoServiceGetVideoIDByUidResult)
	success, err := handler.(video.VideoService).GetVideoIDByUid(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceGetVideoIDByUidArgs() interface{} {
	return video.NewVideoServiceGetVideoIDByUidArgs()
}

func newVideoServiceGetVideoIDByUidResult() interface{} {
	return video.NewVideoServiceGetVideoIDByUidResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Feed(ctx context.Context, req *video.FeedRequest) (r *video.FeedResponse, err error) {
	var _args video.VideoServiceFeedArgs
	_args.Req = req
	var _result video.VideoServiceFeedResult
	if err = p.c.Call(ctx, "Feed", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PutVideo(ctx context.Context, req *video.PutVideoRequest) (r *video.PutVideoResponse, err error) {
	var _args video.VideoServicePutVideoArgs
	_args.Req = req
	var _result video.VideoServicePutVideoResult
	if err = p.c.Call(ctx, "PutVideo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFavoriteVideoInfo(ctx context.Context, req *video.GetFavoriteVideoInfoRequest) (r *video.GetFavoriteVideoInfoResponse, err error) {
	var _args video.VideoServiceGetFavoriteVideoInfoArgs
	_args.Req = req
	var _result video.VideoServiceGetFavoriteVideoInfoResult
	if err = p.c.Call(ctx, "GetFavoriteVideoInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetPublishList(ctx context.Context, req *video.GetPublishListRequest) (r *video.GetPublishListResponse, err error) {
	var _args video.VideoServiceGetPublishListArgs
	_args.Req = req
	var _result video.VideoServiceGetPublishListResult
	if err = p.c.Call(ctx, "GetPublishList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetWorkCount(ctx context.Context, req *video.GetWorkCountRequest) (r *video.GetWorkCountResponse, err error) {
	var _args video.VideoServiceGetWorkCountArgs
	_args.Req = req
	var _result video.VideoServiceGetWorkCountResult
	if err = p.c.Call(ctx, "GetWorkCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetVideoIDByUid(ctx context.Context, req *video.GetVideoIDByUidRequest) (r *video.GetVideoIDByUidResponse, err error) {
	var _args video.VideoServiceGetVideoIDByUidArgs
	_args.Req = req
	var _result video.VideoServiceGetVideoIDByUidResult
	if err = p.c.Call(ctx, "GetVideoIDByUid", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
