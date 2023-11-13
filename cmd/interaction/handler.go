package main

import (
	"context"
	"github.com/benxinm/tiktok/cmd/interaction/pack"
	"github.com/benxinm/tiktok/cmd/interaction/service"
	interaction "github.com/benxinm/tiktok/kitex_gen/interaction"
	"github.com/benxinm/tiktok/pkg/myerrors"
	"github.com/benxinm/tiktok/pkg/utils"
	"github.com/cloudwego/kitex/pkg/klog"
)

// InteractionServiceImpl implements the last service interface defined in the IDL.
type InteractionServiceImpl struct{}

// FavoriteAction implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) FavoriteAction(ctx context.Context, req *interaction.FavoriteActionRequest) (resp *interaction.FavoriteActionResponse, err error) {
	resp = new(interaction.FavoriteActionResponse)
	claims, err := utils.VerifyToken(req.Token)
	if err != nil {
		resp.Base = pack.MakeBaseResp(myerrors.AuthFailedError)
		return resp, nil
	}
	switch req.ActionType {
	case 1:
		if err := service.NewInteractionService(ctx).Favorite(req, claims.UserId); err != nil {
			klog.Errorf("interaction err: %v", err)
			resp.Base = pack.MakeBaseResp(err)
			return resp, nil
		}
	case 2:
		if err := service.NewInteractionService(ctx).UnFavorite(req, claims.UserId); err != nil {
			klog.Errorf("interaction err: %v", err)
			resp.Base = pack.MakeBaseResp(err)
			return resp, nil
		}
	default:
		resp.Base = pack.MakeBaseResp(myerrors.ParamError)
		return resp, nil
	}
	resp.Base = pack.MakeBaseResp(nil)
	return resp, nil
}

// FavoriteList implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) FavoriteList(ctx context.Context, req *interaction.FavoriteListRequest) (resp *interaction.FavoriteListResponse, err error) {
	resp = new(interaction.FavoriteListResponse)
	_, err = utils.VerifyToken(req.Token)
	if err != nil {
		resp.Base = pack.MakeBaseResp(myerrors.AuthFailedError)
		return resp, nil
	}
	vids, err := service.NewInteractionService(ctx).FavoriteList(req)
	if err != nil {
		klog.Errorf("interaction err: %v", err)
		resp.Base = pack.MakeBaseResp(err)
		return resp, nil
	}
	resp.Base = pack.MakeBaseResp(nil)
	resp.VideoList = pack.Videos(vids)
	return
}

// VideoFavoritedCount implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) VideoFavoritedCount(ctx context.Context, req *interaction.VideoFavoritedCountRequest) (resp *interaction.VideoFavoritedCountResponse, err error) {
	resp = new(interaction.VideoFavoritedCountResponse)
	_, err = utils.VerifyToken(req.Token)
	if err != nil {
		resp.Base = pack.MakeBaseResp(myerrors.AuthFailedError)
		return resp, nil
	}
	count, err := service.NewInteractionService(ctx).GetVideoFavoriteCount(req)
	if err != nil {
		klog.Errorf("interaction err: %v", err)
		resp.Base = pack.MakeBaseResp(err)
		return resp, nil
	}
	resp.Base = pack.MakeBaseResp(nil)
	resp.LikeCount = count
	return
}

// UserFavoriteCount implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) UserFavoriteCount(ctx context.Context, req *interaction.UserFavoriteCountRequest) (resp *interaction.UserFavoriteCountResponse, err error) {
	resp = new(interaction.UserFavoriteCountResponse)
	_, err = utils.VerifyToken(req.Token)
	if err != nil {
		resp.Base = pack.MakeBaseResp(myerrors.AuthFailedError)
		return resp, nil
	}
	count, err := service.NewInteractionService(ctx).GetUserFavoriteCount(req)
	if err != nil {
		klog.Errorf("interaction err: %v", err)
		resp.Base = pack.MakeBaseResp(err)
		return resp, nil
	}
	resp.Base = pack.MakeBaseResp(nil)
	resp.LikeCount = count
	return
}

// UserTotalFavorited implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) UserTotalFavorited(ctx context.Context, req *interaction.UserTotalFavoritedRequest) (resp *interaction.UserTotalFavoritedResponse, err error) {
	resp = new(interaction.UserTotalFavoritedResponse)
	_, err = utils.VerifyToken(req.Token)
	if err != nil {
		resp.Base = pack.MakeBaseResp(myerrors.AuthFailedError)
		return resp, nil
	}
	count, err := service.NewInteractionService(ctx).GetUserTotalFavorited(req)
	if err != nil {
		klog.Errorf("interaction err: %v", err)
		resp.Base = pack.MakeBaseResp(err)
		return resp, nil
	}
	resp.Base = pack.MakeBaseResp(nil)
	resp.TotalFavorited = count
	return resp, nil
}

// IsFavorite implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) IsFavorite(ctx context.Context, req *interaction.IsFavoriteRequest) (resp *interaction.IsFavoriteResponse, err error) {
	resp = new(interaction.IsFavoriteResponse)
	_, err = utils.VerifyToken(req.Token)
	if err != nil {
		resp.Base = pack.MakeBaseResp(myerrors.AuthFailedError)
		return resp, nil
	}
	exist, err := service.NewInteractionService(ctx).IsFavorite(req)
	if err != nil {
		klog.Errorf("interaction err: %v", err)
		resp.Base = pack.MakeBaseResp(err)
		return resp, nil
	}
	resp.Base = pack.MakeBaseResp(nil)
	resp.IsFavorite = exist
	return
}

// CommentAction implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) CommentAction(ctx context.Context, req *interaction.CommentActionRequest) (resp *interaction.CommentActionResponse, err error) {
	resp = new(interaction.CommentActionResponse)
	claims, err := utils.VerifyToken(req.Token)
	if err != nil {
		resp.Base = pack.MakeBaseResp(myerrors.AuthFailedError)
		return resp, nil
	}
	curService := service.NewInteractionService(ctx)
	switch req.ActionType {
	case 1:
		if req.CommentText == "" || len(req.CommentText) > 255 {
			resp.Base = pack.MakeBaseResp(myerrors.ParamError)
			return resp, nil
		}
		comment, err := curService.CreateComment(req, claims.UserId)
		if err != nil {
			resp.Base = pack.MakeBaseResp(err)
			return resp, nil
		}
		resp.Comment = comment
	case 2:
		if &req.CommentId == nil {
			resp.Base = pack.MakeBaseResp(myerrors.ParamError)
			return resp, nil
		}
		comment, err := curService.DeleteComment(req, claims.UserId)
		if err != nil {
			resp.Base = pack.MakeBaseResp(err)
			return resp, nil
		}
		resp.Comment = comment
	default:
		resp.Base = pack.MakeBaseResp(myerrors.ParamError)
		return resp, nil
	}
	resp.Base = pack.MakeBaseResp(nil)
	return
}

// CommentList implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) CommentList(ctx context.Context, req *interaction.CommentListRequest) (resp *interaction.CommentListResponse, err error) {
	resp = new(interaction.CommentListResponse)
	_, err = utils.VerifyToken(req.Token)
	if err != nil {
		resp.Base = pack.MakeBaseResp(myerrors.AuthFailedError)
		return resp, nil
	}
	comments, err := service.NewInteractionService(ctx).GetComments(req)
	if err != nil {
		klog.Errorf("interaction err: %v", err)
		resp.Base = pack.MakeBaseResp(err)
		return resp, nil
	}
	resp.Base = pack.MakeBaseResp(nil)
	resp.CommentList = comments
	return
}

// CommentCount implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) CommentCount(ctx context.Context, req *interaction.CommentCountRequest) (resp *interaction.CommentCountResponse, err error) {
	resp = new(interaction.CommentCountResponse)
	_, err = utils.VerifyToken(req.Token)
	if err != nil {
		resp.Base = pack.MakeBaseResp(myerrors.AuthFailedError)
		return resp, nil
	}
	count, err := service.NewInteractionService(ctx).GetCommentCount(req)
	if err != nil {
		resp.Base = pack.MakeBaseResp(err)
		return resp, nil
	}
	resp.Base = pack.MakeBaseResp(nil)
	resp.CommentCount = count
	return
}
