namespace go interaction

include "user.thrift"

struct BaseResp {
    1: i64 code,
    2: string msg,
}

struct Video {
    1: i64 id,
    2: user.User author,
    3: string play_url,
    4: string cover_url,
    5: i64 favorite_count,
    6: i64 comment_count,
    7: bool is_favorite,
    8: string title,
}

struct Comment {
    1: i64 id,
    2: user.User user,
    3: string content,
    4: string create_date,
}

struct FavoriteActionRequest {
    1: i64 video_id,
    2: i64 action_type,
    3: string token,
}

struct FavoriteActionResponse {
    1: BaseResp base,
}

struct FavoriteListRequest {
    1: i64 user_id,
    2: string token,
}

struct FavoriteListResponse {
    1: BaseResp base,
    2: list<Video> video_list,
}

struct VideoFavoritedCountRequest {
    1: i64 video_id,
    2: string token,
}

struct VideoFavoritedCountResponse {
    1: BaseResp base,
    2: i64 like_count,
}

struct UserTotalFavoritedRequest {
    1: string token,
    2: i64 user_id,
}

struct UserTotalFavoritedResponse {
    1: BaseResp base,
    2: i64 total_favorited,
}

struct UserFavoriteCountRequest {
    1: i64 user_id,
    2: string token,
}

struct UserFavoriteCountResponse {
    1: BaseResp base,
    2: i64 like_count,
}

struct IsFavoriteRequest {
    1: i64 user_id,
    2: i64 video_id,
    3: string token,
}

struct IsFavoriteResponse {
    1: BaseResp base,
    2: bool is_favorite,
}

struct CommentActionRequest {
    1: i64 video_id,
    2: i64 action_type,
    3: string token,
    4: string comment_text,
    5: i64 comment_id,
}

struct CommentActionResponse {
    1: BaseResp base,
    2: Comment comment,
}

struct CommentListRequest {
    1: i64 video_id,
    2: string token,
}

struct CommentListResponse {
    1: BaseResp base,
    2: list<Comment> comment_list,
}

struct CommentCountRequest {
    1: i64 video_id,
    2: string token,
}

struct CommentCountResponse {
    1: BaseResp base,
    2: i64 comment_count,
}

service InteractionService {
    FavoriteActionResponse FavoriteAction(1: FavoriteActionRequest req)
    FavoriteListResponse FavoriteList(1 : FavoriteListRequest req)
    VideoFavoritedCountResponse VideoFavoritedCount(1 : VideoFavoritedCountRequest req)
    UserFavoriteCountResponse UserFavoriteCount(1 : UserFavoriteCountRequest req)
    UserTotalFavoritedResponse UserTotalFavorited(1 : UserTotalFavoritedRequest req)
    IsFavoriteResponse IsFavorite(1 : IsFavoriteRequest req)
    CommentActionResponse CommentAction(1 : CommentActionRequest req)
    CommentListResponse CommentList(1 : CommentListRequest req)
    CommentCountResponse CommentCount(1 : CommentCountRequest req)
}