package main

import (
	"context"
	interaction "github.com/benxinm/tiktok/cmd/interaction/kitex_gen/interaction"
)

// InteractionServiceImpl implements the last service interface defined in the IDL.
type InteractionServiceImpl struct{}

// FavoriteAction implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) FavoriteAction(ctx context.Context, req *interaction.FavoriteActionRequest) (resp *interaction.FavoriteActionResponse, err error) {
	// TODO: Your code here...
	return
}

// FavoriteList implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) FavoriteList(ctx context.Context, req *interaction.FavoriteListRequest) (resp *interaction.FavoriteListResponse, err error) {
	// TODO: Your code here...
	return
}

// VideoFavoritedCount implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) VideoFavoritedCount(ctx context.Context, req *interaction.VideoFavoritedCountRequest) (resp *interaction.VideoFavoritedCountResponse, err error) {
	// TODO: Your code here...
	return
}

// UserFavoriteCount implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) UserFavoriteCount(ctx context.Context, req *interaction.UserFavoriteCountRequest) (resp *interaction.UserFavoriteCountResponse, err error) {
	// TODO: Your code here...
	return
}

// UserTotalFavorited implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) UserTotalFavorited(ctx context.Context, req *interaction.UserTotalFavoritedRequest) (resp *interaction.UserTotalFavoritedResponse, err error) {
	// TODO: Your code here...
	return
}

// IsFavorite implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) IsFavorite(ctx context.Context, req *interaction.IsFavoriteRequest) (resp *interaction.IsFavoriteResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentAction implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) CommentAction(ctx context.Context, req *interaction.CommentActionRequest) (resp *interaction.CommentActionResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentList implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) CommentList(ctx context.Context, req *interaction.CommentListRequest) (resp *interaction.CommentListResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentCount implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) CommentCount(ctx context.Context, req *interaction.CommentCountRequest) (resp *interaction.CommentCountResponse, err error) {
	// TODO: Your code here...
	return
}
