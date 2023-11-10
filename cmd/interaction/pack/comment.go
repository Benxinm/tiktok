package pack

import (
	"github.com/benxinm/tiktok/cmd/interaction/dal/db"
	"github.com/benxinm/tiktok/kitex_gen/interaction"
	"github.com/benxinm/tiktok/kitex_gen/user"
)

func Comment(data *db.Comment, user *user.User) *interaction.Comment {
	if data == nil {
		return nil
	}
	return &interaction.Comment{
		Id:         data.Id,
		User:       user,
		Content:    data.Content,
		CreateDate: data.CreatedAt.Format("2023-11-8 15:24:00"),
	}
}
