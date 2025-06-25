package utils

import (
	"github.com/wushiling50/aster/gen/base"
	"github.com/wushiling50/aster/pkg/errno"
)

func IsSuccess(baseResp *base.Base) bool {
	return baseResp.Code == errno.SuccessCode
}
