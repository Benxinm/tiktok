package service

import (
	"bytes"
	"fmt"
	"github.com/benxinm/tiktok/kitex_gen/video"
	"github.com/cloudwego/kitex/pkg/klog"
)

func (s *VideoService) UploadVideo(req *video.PutVideoRequest, videoName string) (err error) {
	fileReader := bytes.NewReader(req.VideoFile)
	err = s.bucket.PutObject(fmt.Sprintf("/%s", videoName), fileReader)
	if err != nil {
		klog.Errorf("upload file error: %v\n", err)
	}
	return
}
