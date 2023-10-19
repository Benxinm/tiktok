package myerrors

var (
	Success = NewMyError(SuccessCode, "Success")
	//Common
	ServiceError = NewMyError(InternalErrorCode, "something wrong in the service")
	// User
	UserExistedError  = NewMyError(ParamErrorCode, "user already existed")
	UserNotFoundError = NewMyError(InternalErrorCode, "user not found")
)
