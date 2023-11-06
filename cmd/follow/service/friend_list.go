package service

//TODO Other module needed
//func (s *FollowService) FriendList(req *follow.FriendListRequest) ([]*follow.FriendUser, error) {
//	friendList := make([]*follow.FriendUser,20)
//	var eg errgroup.Group
//	userList, err := cache.FriendList(s.ctx, req.UserId)
//	if err != nil {
//		return nil, err
//	}
//
//	for index,userId := range userList{
//		eg.Go(func() error {
//			user, err := rpc.GetUser(s.ctx, &user.InfoRequest{UserId: userId, Token: req.Token})
//			if err != nil {
//				return err
//			}
//
//			return nil
//		})
//
//	}
//}
