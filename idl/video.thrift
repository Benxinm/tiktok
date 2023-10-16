namespace go video

struct BaseResp{
    1: i64 code,
    2: string msg,
}

struct User{
    1: i64 id,
    2: string name,
    3: i64 follow_count,
    4: i64 follower_count,
    5: bool is_follow
    6: string avatar,
    7: string background_image,
    8: string signature,
    9: i64 total_favorited,
    10: i64 work_count,
    11: i64 favorited_count,
}

struct Video{
    1: i64 id,
    2: User author,
    3: string play_url,
    4: string cover_url,
    5: i64 favorite_count,
    6: i64 comment_count,
    7: bool is_favourite,
    8: string title,
}

struct FeedRequest{
    1: i64 latest_time,
    2: string token,
}

struct FeedResponse{
    1: BaseResp base,
    2: i64 next_time,
    3: list<Video> video_list,
}

struct PutVideoRequest{
    1: binary video_file,
    2: string title,
    3: string token,
}

struct PutVideoResponse{
    1: BaseResp base,
}

struct GetFavoriteVideoInfoRequest{
    1: list<i64> video_id,
    2: string token,
}

struct GetFavoriteVideoInfoResponse{
    1: BaseResp base,
    2: list<Video> video_list,
}

struct GetPublishListRequest{
    1: string token,
    2: i64 user_id,
}

struct GetPublishListResponse{
    1: BaseResp base,
    2: list<Video> list_video,
}

struct GetWorkCountRequest{
    1: string token,
    2: i64 user_id,
}

struct GetWorkCountResponse{
    1: BaseResp base,
    2: i64 work_count,
}

struct GetVideoIDByUidRequest{
    1: string token,
    2: i64 user_id,
}
struct GetVideoIDByUidResponse{
    1: BaseResp base,
    2: list<i64> video_id,
}

service VideoService{
    FeedResponse Feed(1: FeedRequest req)
    PutVideoResponse PutVideo(1: PutVideoRequest req)
    GetFavoriteVideoInfoResponse GetFavoriteVideoInfo(1: GetFavoriteVideoInfoRequest req)
    GetPublishListResponse GetPublishList(1: GetPublishListRequest req)
    GetWorkCountResponse GetWorkCount(1: GetWorkCountRequest req)
    GetVideoIDByUidResponse GetVideoIDByUid(1: GetVideoIDByUidRequest req)
}
