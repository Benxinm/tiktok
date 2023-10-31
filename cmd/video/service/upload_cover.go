package service

import (
	"bytes"
	"fmt"
	"github.com/benxinm/tiktok/kitex_gen/video"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/grpc/syscall"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func (s *VideoService) UploadCover(req *video.PutVideoRequest, coverName string) (err error) {
	var imageBuffer bytes.Buffer
	pipePath := filepath.Join(os.TempDir(), fmt.Sprintf("input_pipe_%d", time.Now().UnixMilli()))

	err = syscall.Mkfifo(pipePath, 0666)
	if err != nil && !os.IsExist(err) {
		klog.Errorf("failed to create named pipe:%v\n", err)
		return err
	}
	defer os.Remove(pipePath)

	cmd := exec.Command("ffmpeg", "-i", pipePath, "-vframes", "1", "-f", "image2pipe", "-vcodec", "jpg", "-")
	cmd.Stdout = &imageBuffer
	cmd.Stderr = os.Stderr

	go func() {
		pipeWiter, err := os.OpenFile(pipePath, os.O_WRONLY, os.ModeNamedPipe)
		if err != nil {
			klog.Errorf("failed to open pipe :%v", err)
			return
		}
		defer pipeWiter.Close()
		_, err = pipeWiter.Write(req.VideoFile)
		if err != nil {
			klog.Errorf("failed to write to pipe: %v", err)
			return
		}
	}()
	err = cmd.Run()

	if err != nil {
		klog.Errorf("FFmpeg exec error: %v", err)
		return err
	}

	imageReader := bytes.NewReader(imageBuffer.Bytes())
	err = s.bucket.PutObject(fmt.Sprintf("/%s", coverName), imageReader)
	if err != nil {
		klog.Errorf("failed to upload cover: %v", err)
	}
	return err
}
