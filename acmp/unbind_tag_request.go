package acmp

import (
	"encoding/json"
	"errors"
	"strconv"
)

// UnbindTagResponse 返回响应体
type UnbindTagResponse struct {
	ErrorMessage
}

// String 序列化响应
func (l *UnbindTagResponse) String() string {
	body, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(body)
}

// UnbindTagRequest 绑定标签请求
type UnbindTagRequest struct {
	Request *Request
}

// DoActionWithException 发起http请求
func (u *UnbindTagRequest) DoActionWithException() (resp *UnbindTagResponse, err error) {
	if u != nil && u.Request != nil {
		resp := &UnbindTagResponse{}
		body, httpCode, err := u.Request.Do("UnbindTag")
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

// UnbindTag 解除绑定标签请求接口
func UnbindTag(appKey int, clientKey, keyType, tagName string) *UnbindTagRequest {
	req := newRequset()
	req.Put("Version", "2016-08-01")
	req.Put("Action", "UnbindTag")
	req.Put("AppKey", strconv.Itoa(appKey))
	req.Put("ClientKey", clientKey)
	req.Put("KeyType", keyType)
	req.Put("TagName", tagName)

	r := &UnbindTagRequest{Request: req}
	return r
}
