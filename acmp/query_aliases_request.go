package acmp

import (
	"encoding/json"
	"errors"
	"strconv"
)

// AliasInfo 别名信息结构体
type AliasInfo struct {
	AliasName string `json:"AliasName"`
}

// AliasInfos 别名信息
type AliasInfos struct {
	AliasInfos []AliasInfo `json:"AliasInfos"`
}

// QueryAliasesRespones 查询别名信息响应体
type QueryAliasesRespones struct {
	ErrorMessage
	AliasInfos *AliasInfos `json:"AliasInfos,omitempty"`
}

// GetAliasInfos 获取响应体的别名信息
func (q *QueryAliasesRespones) GetAliasInfos() *AliasInfos {
	if q != nil && q.AliasInfos != nil {
		return q.AliasInfos
	}
	return nil
}

// String 序列化响应
func (q *QueryAliasesRespones) String() string {
	body, err := json.Marshal(q)
	if err != nil {
		return ""
	}
	return string(body)
}

// QueryAliasesRequest http请求结构体
type QueryAliasesRequest struct {
	Request *Request
}

// DoActionWithException 发起http请求
func (q *QueryAliasesRequest) DoActionWithException() (resp *QueryAliasesRespones, err error) {
	if q != nil && q.Request != nil {
		resp := &QueryAliasesRespones{}
		body, httpCode, err := q.Request.Do("QueryAliases")
		resp.SetHTTPCode(httpCode)
		if err != nil {
			return resp, err
		}
		err = json.Unmarshal(body, resp)
		if err != nil {
			return resp, err
		}
		if httpCode != 200 {
			return resp, errors.New(resp.GetCode())
		}
		return resp, nil
	}
	return nil, errors.New("SendRequest is nil")
}

// QueryAliases 查询绑定别名请求接口
func QueryAliases(appKey int, deviceID string) *QueryAliasesRequest {
	req := newRequset()
	req.Put("Version", "2016-08-01")
	req.Put("Action", "QueryTags")
	req.Put("AppKey", strconv.Itoa(appKey))
	req.Put("DeviceId", deviceID)

	r := &QueryAliasesRequest{Request: req}
	return r
}
