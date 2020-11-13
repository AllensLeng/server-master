package services

import (
	"net/http"

	"server-master/common"
)

// ExampleGet 示例 Get 请求
func ExampleGet(w http.ResponseWriter, _ *http.Request) {
	respData := new(common.RespData)
	defer common.StdResp(w, respData)

	respData.Success = true
	respData.Data = "Request Succeeded"
}

// examplePostReq 示例 Post 请求结构
type examplePostReq struct {
	OK string `json:"ok"`
}

// ExamplePost 示例 Post 请求
func ExamplePost(w http.ResponseWriter, req *http.Request) {
	respData := new(common.RespData)
	defer common.StdResp(w, respData)

	reqData := new(examplePostReq)
	err := common.GetRequestJSONBody(req, reqData)
	if err != nil {
		respData.Error = err.Error()
		return
	}

	if reqData.OK == "true" {
		respData.Success = true
		respData.Data = "Request Succeeded"
	}
}
