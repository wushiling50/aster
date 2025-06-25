package pack

import (
	"strconv"

	"github.com/wushiling50/aster/pkg/errno"
)

type Base struct {
	Code string `json:"code"`
	Msg  string `json:"message"`
}

type RespWithData struct {
	Code string `json:"code"`
	Msg  string `json:"message"`
	Data any    `json:"data"`
}

func RespSuccess() *Base {
	Errno := errno.Success
	return &Base{
		Code: strconv.FormatInt(Errno.ErrorCode, 10),
		Msg:  Errno.ErrorMsg,
	}
}

func RespData(data any) *RespWithData {
	return &RespWithData{
		Code: strconv.FormatInt(errno.SuccessCode, 10),
		Msg:  "Success",
		Data: data,
	}
}

func RespError(err error) *Base {
	Errno := errno.ConvertErr(err)
	return &Base{
		Code: strconv.FormatInt(Errno.ErrorCode, 10),
		Msg:  Errno.ErrorMsg,
	}
}
