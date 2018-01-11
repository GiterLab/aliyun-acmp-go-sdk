package acmp

import (
	"encoding/json"
	"errors"
)

// CancelPushResponse 取消定时推送任务结构体
type CancelPushResponse struct {
	ErrorMessage
}

// String 响应序列化
func (c *CancelPushResponse) String() string {
	body, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(body)
}

// CancelPushRequest 取消定时推送任务请求结构体
type CancelPushRequest struct {
	Request *Request
}

// DoActionWithException 取消定时推送任务http请求
func (c *CancelPushRequest) DoActionWithException() (resp *CancelPushResponse, err error) {
	if c != nil && c.Request != nil {
		resp := &CancelPushResponse{}
		body, httpCode, err := c.Request.Do("CancelPush")
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

// CancelPush 取消定时推送任务接口
func CancelPush(messageID string) *CancelPushRequest {
	if messageID == "" {
		return nil
	}
	req := newRequset()
	req.Put("Version", "2016-08-01")
	req.Put("Action", "CancelPush")
	req.Put("MessageId", messageID)

	r := &CancelPushRequest{Request: req}
	return r
}
