package services

import (
	"fmt"
	"net/http"

	"server-master/common"
)

// ExampleHandle 示例握手函数
func ExampleHandle(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		exampleGet(w, req)
	case http.MethodPost:
		examplePost(w, req)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not match method"))
	}
}

// exampleGet 示例 Get 请求
func exampleGet(w http.ResponseWriter, _ *http.Request) {
	respData := new(common.RespData)
	defer common.StdResp(w, respData)

	respData.Success = true
	respData.Data = "Completed Get Request"
}

// examplePostReq 示例 Post 请求结构
type examplePostReq struct {
	OK string `json:"ok"`
}

// examplePost 示例 Post 请求
func examplePost(w http.ResponseWriter, req *http.Request) {
	respData := new(common.RespData)
	defer common.StdResp(w, respData)

	reqData := new(examplePostReq)
	err := common.GetRequestJSONBody(req, reqData)
	if err != nil {
		respData.Error = err.Error()
		return
	}

	respData.Success = true
	respData.Data = fmt.Sprintf("Completed Get Request , req.OK : %v", reqData.OK == "true")
}
