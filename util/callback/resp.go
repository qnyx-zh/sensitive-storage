package callback

import (
	"sensitive-storage/constant"
	"sensitive-storage/module/resp"
)

func CallBackSuccess(data interface{}) resp.Resp {
	result := &resp.Resp{
		Status: constant.RespSuccessStr,
		Code:   constant.RespSuccess,
		Data:   data,
	}
	return *result
}

func CallBackFail(msg string) resp.Resp {
	result := &resp.Resp{
		Status: constant.RespFailStr,
		Code:   constant.RespFail,
		ErrMsg: msg,
	}
	return *result
}
