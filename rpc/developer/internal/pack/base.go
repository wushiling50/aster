package pack

import (
	"github.com/wushiling50/aster/gen/base"
	"github.com/wushiling50/aster/pkg/errno"
)

func BuildBaseResp(err error) *base.Base {
	if err == nil {
		return &base.Base{
			Code:    errno.SuccessCode,
			Message: errno.Success.ErrorMsg,
		}
	}
	Errno := errno.ConvertErr(err)
	return &base.Base{
		Code:    Errno.ErrorCode,
		Message: Errno.ErrorMsg,
	}
}

func BuildSuccessResp() *base.Base {
	return BuildBaseResp(nil) // 直接调用原始函数，传入 nil 表示无错误
}
