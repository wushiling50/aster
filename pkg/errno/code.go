package errno

const (
	// For success
	SuccessCode = 20000
	SuccessMsg  = "ok"

	// For service
	ServiceErrorCode = 50000

	// Error
	ParamErrorCode          = 50001 // 参数错误
	UnexpectedTypeErrorCode = 50002 // 未知类型
)
