package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JsonResult 返回结构
type JsonResultTest struct {
	Code     int         `json:"code"`
	ErrorMsg string      `json:"errorMsg,omitempty"`
	Data     interface{} `json:"data"`
}

// test
func Receive(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	res.Code = -1
	res.Data = "hello"

	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}
