package myerrors

var (
	Success = NewMyError(SuccessCode, "Success")
	//Common
	ServiceError = NewMyError(InternalErrorCode, "something wrong in the service")
	ParamError   = NewMyError(ParamErrorCode, "parameter error")
	// User
	UserExistedError  = NewMyError(ParamErrorCode, "user already existed")
	UserNotFoundError = NewMyError(InternalErrorCode, "user not found")
	AuthFailedError   = NewMyError(AuthFailedErrorCode, "auth failed")
	//Follow
	FollowNotFoundError = NewMyError(InternalErrorCode, "follow not found but exec unfollow")
	//Interaction
	CommentNotFoundError = NewMyError(InternalErrorCode, "comments not found")
	//Video
	VideoUploadError = NewMyError(NetworkFailedErrorCode, "Upload Failed")
)
