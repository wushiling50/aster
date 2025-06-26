package errno

var (
	Success = NewErrNo(SuccessCode, "ok")

	ParamError = NewErrNo(ParamErrorCode, "参数错误") // 参数校验失败，可能是参数为空、参数类型错误等

	BizError = NewErrNo(BizErrorCode, "请求业务出现问题")

	InternalServiceError   = NewErrNo(InternalServiceErrorCode, "内部服务错误")
	InternalDatabaseError  = NewErrNo(InternalDatabaseErrorCode, "数据库错误")
	InternalRedisError     = NewErrNo(InternalRedisErrorCode, "Redis 错误")
	InternalJSONError      = NewErrNo(InternalJSONErrorCode, "JSON 错误")
	InternalLanguagesError = NewErrNo(InternalLanguagesErrorCode, "获取语言数据错误")
	InternalGithubError    = NewErrNo(InternalGithubErrorCode, "Github 错误")
)
