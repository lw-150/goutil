package goutil

import "github.com/kataras/iris/v12"

/**
json结构
*/
type BackJson struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

/**
响应Json
*/
func RepoJson(code int, data interface{}, message string, ctx iris.Context) {
	backJson := BackJson{}
	backJson.Code = code
	backJson.Data = data
	backJson.Message = message
	_, _ = ctx.JSON(backJson)
}
