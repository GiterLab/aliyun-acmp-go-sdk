package acmp

import (
	"encoding/json"
	"errors"
	"strconv"
)

// BindTagResponse 返回响应体
type BindTagResponse struct {
	ErrorMessage
}

// BindTagRequest 绑定标签请求
type BindTagRequest struct {
	Request *Request
}

// String 序列化响应
func (l *BindTagResponse) String() string {
	body, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(body)
}

// DoActionWithException 发起http请求
func (b *BindTagRequest) DoActionWithException() (resp *BindTagResponse, err error) {
	if b != nil && b.Request != nil {
		resp := &BindTagResponse{}
		body, httpCode, err := b.Request.Do("BindTag")
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

// BindTag 绑定标签请求接口
func BindTag(appKey int, clientKey, keyType, tagName string) *BindTagRequest {
	req := newRequset()
	req.Put("Version", "2016-08-01")
	req.Put("Action", "BindTag")
	req.Put("AppKey", strconv.Itoa(appKey))
	req.Put("ClientKey", clientKey)
	req.Put("KeyType", keyType)
	req.Put("TagName", tagName)

	r := &BindTagRequest{Request: req}
	return r
}
