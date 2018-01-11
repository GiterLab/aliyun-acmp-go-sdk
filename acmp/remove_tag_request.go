package acmp

import (
	"encoding/json"
	"errors"
	"strconv"
)

// RemoveTagResponse 返回响应体
type RemoveTagResponse struct {
	ErrorMessage
}

// RemoveTagRequest 绑定标签请求
type RemoveTagRequest struct {
	Request *Request
}

// String 序列化响应
func (l *RemoveTagResponse) String() string {
	body, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(body)
}

// DoActionWithException 发起http请求
func (r *RemoveTagRequest) DoActionWithException() (resp *RemoveTagResponse, err error) {
	if r != nil && r.Request != nil {
		resp := &RemoveTagResponse{}
		body, httpCode, err := r.Request.Do("RemoveTag")
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

// RemoveTag 删除标签请求接口
func RemoveTag(appKey int, tagName string) *RemoveTagRequest {
	req := newRequset()
	req.Put("Version", "2016-08-01")
	req.Put("Action", "RemoveTag")
	req.Put("AppKey", strconv.Itoa(appKey))
	req.Put("TagName", tagName)

	r := &RemoveTagRequest{Request: req}
	return r
}
