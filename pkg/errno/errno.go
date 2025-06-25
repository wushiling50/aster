package errno

const (
	// For microservices
	SuccessCode = 10000
	SuccessMsg  = "ok"

	// Error
	/*
		200xx: 参数错误，Param 打头
		400xx: 业务错误，Biz 打头
		500xx: 内部错误，Internal 打头
	*/

	ParamErrorCode = 20001 // 参数错误

	BizErrorCode = 40001 // 业务错误

	InternalServiceErrorCode  = 50001 // 未知服务错误
	InternalDatabaseErrorCode = 50002 // 数据库错误
	InternalRedisErrorCode    = 50003 // Redis错误
)
