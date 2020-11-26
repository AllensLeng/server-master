package common

import (
	"encoding/json"
	"net/http"
)

// RespData 响应数据
type RespData struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}

// StdResp 标准响应 standard response
func StdResp(w http.ResponseWriter, resp *RespData) {

	// 响应格式为 JSON
	w.Header().Set("content-type", "application/json")

	// 默认 OK
	code := http.StatusOK

	// 序列化结构数据
	data, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 失败返回
	if !resp.Success {
		code = http.StatusBadRequest
	}

	// 响应
	w.WriteHeader(code)
	w.Write(data)
}
