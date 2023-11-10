package pack

import (
	"github.com/benxinm/tiktok/kitex_gen/interaction"
	"github.com/benxinm/tiktok/kitex_gen/user"
	"github.com/benxinm/tiktok/kitex_gen/video"
)

func User(data *video.User) *user.User {
	if data == nil {
		return nil
	}
	return &user.User{
		Id:              data.Id,
		Name:            data.Name,
		Avatar:          data.Avatar,
		FollowCount:     data.FollowCount,
		FollowerCount:   data.FollowerCount,
		IsFollow:        data.IsFollow,
		BackgroundImage: data.BackgroundImage,
		Signature:       data.Signature,
		WorkCount:       data.WorkCount,
		FavoritedCount:  data.FavoritedCount,
		TotalFavorited:  data.TotalFavorited,
	}
}

func Video(data *video.Video) *interaction.Video {
	if data == nil {
		return nil
	}
	return &interaction.Video{
		Id:            data.Id,
		Author:        User(data.Author),
		PlayUrl:       data.PlayUrl,
		CoverUrl:      data.CoverUrl,
		FavoriteCount: data.FavoriteCount,
		CommentCount:  data.CommentCount,
		IsFavorite:    data.IsFavourite,
		Title:         data.Title,
	}
}

func Videos(data []*video.Video) []*interaction.Video {
	if data == nil {
		return nil
	}
	videos := make([]*interaction.Video, 0, len(data))
	for _, data := range data {
		videos = append(videos, Video(data))
	}
	return videos
}
