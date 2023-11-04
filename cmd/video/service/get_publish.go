package service

import (
	"github.com/benxinm/tiktok/cmd/video/dal/db"
	"github.com/benxinm/tiktok/cmd/video/pack"
	"github.com/benxinm/tiktok/cmd/video/rpc"
	"github.com/benxinm/tiktok/kitex_gen/interaction"
	"github.com/benxinm/tiktok/kitex_gen/user"
	"github.com/benxinm/tiktok/kitex_gen/video"
	"golang.org/x/sync/errgroup"
)

func (s *VideoService) GetPublishVideoInfo(req *video.GetPublishListRequest) ([]*video.Video, error) {
	videoList, err := db.GetVideoByUid(s.ctx, req.UserId)
	var eg errgroup.Group

	type result struct {
		user     *user.User
		likes    int64
		comments int64
		isLiked  bool
	}
	results := make([]result, len(videoList))
	for i := 0; i < len(videoList); i++ {
		index := i
		eg.Go(func() error {
			user, err := rpc.GetUser(s.ctx, &user.InfoRequest{
				UserId: videoList[index].UserID,
				Token:  req.Token,
			})
			if err != nil {
				return err
			}
			likes, err := rpc.GetVideoLikeCount(s.ctx, &interaction.VideoFavoritedCountRequest{
				VideoId: videoList[index].Id,
				Token:   req.Token,
			})
			if err != nil {
				return err
			}
			comments, err := rpc.GetCommentCount(s.ctx, &interaction.CommentCountRequest{
				VideoId: videoList[index].Id,
				Token:   req.Token,
			})
			if err != nil {
				return err
			}
			isLiked, err := rpc.GetVideoIsLiked(s.ctx, &interaction.InteractionServiceIsFavoriteArgs{Req: &interaction.IsFavoriteRequest{
				UserId:  videoList[index].UserID,
				VideoId: videoList[index].Id,
				Token:   req.Token,
			}})
			if err != nil {
				return err
			}
			results[index] = result{
				user:     user,
				likes:    likes,
				comments: comments,
				isLiked:  isLiked,
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return nil, err
	}
	var userList []*user.User
	var likesList []int64
	var commentsList []int64
	var isLikedList []bool
	for _, result := range results {
		userList = append(userList, result.user)
		likesList = append(likesList, result.likes)
		commentsList = append(commentsList, result.comments)
		isLikedList = append(isLikedList, result.isLiked)
	}
	return pack.VideoList(videoList, userList, likesList, commentsList, isLikedList), err
}
