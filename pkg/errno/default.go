package errno

var (
	// Success
	Success = NewErrNo(SuccessCode, "Success")

	// service
	ServiceError = NewErrNo(ServiceErrorCode, "service is unable to work successfully")

	// Error
	ParamError          = NewErrNo(ParamErrorCode, "parameter error")
	UnexpectedTypeError = NewErrNo(UnexpectedTypeErrorCode, "unexpected type error")
)
