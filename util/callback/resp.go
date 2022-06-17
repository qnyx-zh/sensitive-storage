package callback

import (
	"sensitive-storage/constant"
	"sensitive-storage/module/resp"
)

func SuccessData(data any) resp.Resp {
	result := &resp.Resp{
		Status: constant.RespSuccessStr,
		Code:   constant.RespSuccess,
		Data:   data,
	}
	return *result
}

func Success() resp.Resp {
	result := &resp.Resp{
		Status: constant.RespSuccessStr,
		Code:   constant.RespSuccess,
	}
	return *result
}

func BackFail(msg string) resp.Resp {
	result := &resp.Resp{
		Status: constant.RespFailStr,
		Code:   constant.RespFail,
		ErrMsg: msg,
	}
	return *result
}
