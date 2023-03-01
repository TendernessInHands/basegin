package log

import "net/http"

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	data interface{}
}

func ResultSuccess() Result {
	return Result{Code: http.StatusOK, Msg: "请求成功!"}
}

func ResultCodeMsg(code int, msg string) Result {
	return Result{Code: code, Msg: msg}
}

func ResultData(data string) Result {
	return Result{http.StatusOK, "请求成功!", data}
}

func ResultMsgData(msg string, data string) Result {
	return Result{http.StatusOK, msg, data}
}

func ResultMap(data map[string]interface{}) Result {
	return Result{http.StatusOK, "请求成功!", data}
}

func ResultMsgMap(msg string, data map[string]interface{}) Result {
	return Result{http.StatusOK, msg, data}
}

func ResultSuccess201() Result {
	code := http.StatusCreated
	return Result{Code: code, Msg: http.StatusText(code)}
}

func ResultFail401(msg string) Result {
	return Result{Code: http.StatusUnauthorized, Msg: "暂无权限，请联系管理员进行授权！"}
}

func ResultFail404() Result {
	return Result{Code: http.StatusNotFound, Msg: "请求失败，请联系管理进行处理！"}
}

func ResultFail500() Result {
	return Result{Code: http.StatusInternalServerError, Msg: "当前服务器请求错误，请联系管理员进行处理！"}
}
