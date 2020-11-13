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
	w.Header().Set("content-type", "application/json")

	code := http.StatusOK
	data, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !resp.Success {
		code = http.StatusBadRequest
	}
	w.WriteHeader(code)
	w.Write(data)
}
