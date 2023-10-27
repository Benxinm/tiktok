package pack

import (
	"github.com/benxinm/tiktok/cmd/video/dal/db"
	"github.com/benxinm/tiktok/kitex_gen/user"
	"github.com/benxinm/tiktok/kitex_gen/video"
)

func Video(data *db.Video, user *user.User, favourites int64, comments int64, isFavourite bool) *video.Video {
	if data == nil {
		return nil
	}
	return &video.Video{
		Id: data.Id,
		Author: &video.User{
			Id:              user.Id,
			Name:            user.Name,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			IsFollow:        user.IsFollow,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			TotalFavorited:  user.TotalFavorited,
			WorkCount:       user.WorkCount,
			FavoritedCount:  user.FavoritedCount,
		},
		PlayUrl:       data.PlayUrl,
		CoverUrl:      data.CoverUrl,
		FavoriteCount: favourites,
		CommentCount:  comments,
		IsFavourite:   isFavourite,
		Title:         data.Title,
	}
}

func VideoList(data []db.Video, userList []*user.User, favoriteCountList []int64, commentCountList []int64, isFavoriteList []bool) []*video.Video {
	videoList := make([]*video.Video, 0, len(data))
	for i := 0; i < len(data); i++ {
		videoList = append(videoList, Video(&data[i], userList[i], favoriteCountList[i], commentCountList[i], isFavoriteList[i]))
	}
	return videoList
}
