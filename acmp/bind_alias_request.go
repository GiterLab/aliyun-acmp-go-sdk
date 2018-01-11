package acmp

import (
	"encoding/json"
	"errors"
	"strconv"
)

// BindAliasResponse 返回响应体
type BindAliasResponse struct {
	ErrorMessage
}

// BindAliasRequest 绑定别名请求
type BindAliasRequest struct {
	Request *Request
}

// String 序列化响应
func (l *BindAliasResponse) String() string {
	body, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(body)
}

// DoActionWithException 发起http请求
func (b *BindAliasRequest) DoActionWithException() (resp *BindAliasResponse, err error) {
	if b != nil && b.Request != nil {
		resp := &BindAliasResponse{}
		body, httpCode, err := b.Request.Do("BindAlias")
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

// BindAlias 绑定别名请求接口
func BindAlias(appKey int, DeviceID, AliasName string) *BindAliasRequest {
	req := newRequset()
	req.Put("Version", "2016-08-01")
	req.Put("Action", "BindAlias")
	req.Put("AppKey", strconv.Itoa(appKey))
	req.Put("DeviceId", DeviceID)
	req.Put("AliasName", AliasName)

	r := &BindAliasRequest{Request: req}
	return r
}
