namespace go chat

struct BaseResp {
    1: i64 code
    2: string msg
}

struct Message {
    1: i64 id,
    2: i64 to_user_id,
    3: i64 from_user_id,
    4: string content,
    5: string create_time,
}

struct MessagePostRequest {
    1: string token
    2: i64 to_user_id
    3: string content,
    4:i64 action_type,
}

struct MessagePostReponse {
    1: BaseResp base
}

struct MessageListRequest {
    1: string token
    2: i64 to_user_id    //  ->对方用户id
}

struct MessageListResponse {      
    1: BaseResp base
    2: list<Message> message_list //  ->按时间倒序排列
    3: i64 total
}

service MessageService {
    MessagePostReponse MessagePost(1:MessagePostRequest req)
    MessageListResponse MessageList(1:MessageListRequest req)
}
