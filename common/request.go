package common

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// GetRequestJSONBody 获取请求体
func GetRequestJSONBody(req *http.Request, recipient interface{}) (err error) {
	var data []byte

	// 读取请求体
	data, err = ioutil.ReadAll(req.Body)
	if err != nil {
		err = errors.Wrap(err, "ioutil.ReadAll")
		return
	}

	// 反序列化请求结构
	err = json.Unmarshal(data, &recipient)
	return errors.Wrap(err, "json.Unmarshal")
}
