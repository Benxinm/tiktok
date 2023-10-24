package myerrors

var (
	Success = NewMyError(SuccessCode, "Success")
	//Common
	ServiceError = NewMyError(InternalErrorCode, "something wrong in the service")
	ParamError   = NewMyError(ParamErrorCode, "parameter error")
	// User
	UserExistedError  = NewMyError(ParamErrorCode, "user already existed")
	UserNotFoundError = NewMyError(InternalErrorCode, "user not found")
)
