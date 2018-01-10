package acmp

import (
	"encoding/json"
	"errors"
	"strconv"
)

// PushMessageResponse 推送消息响应结构体
type PushMessageResponse struct {
	ErrorMessage
	MessageId string `json:"message_id"`
}

// getMessageId 获取messageId
func (p *PushMessageResponse) GetMessageId() string {
	if p != nil && p.MessageId != "" {
		return p.MessageId
	}
	return ""
}

// String 序列化响应
func (p *PushMessageResponse) String() string {
	body, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(body)
}

// PushMessageToiOSRequest http请求
type PushMessageToiOSRequest struct {
	Request *Request
}

// DoActionWithException 发送http请求
func (p *PushMessageToiOSRequest) DoActionWithException() (resp *PushMessageResponse, err error) {
	if p != nil && p.Request != nil {
		resp := &PushMessageResponse{}
		body, httpCode, err := p.Request.Do("PushMessageToIos")
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

// PushMessageToiOSRequest 推送iOS消息接口
func PushMessageToiOS(appKey int, target, targetValue, body string) *PushMessageToiOSRequest {
	if target == "" || targetValue == "" {
		return nil
	}
	req := newRequset()
	req.Put("Version", "2016-08-01")
	req.Put("Action", "PushMessageToIos")
	req.Put("AppKey", strconv.Itoa(appKey))
	req.Put("Target", target)
	req.Put("TargetValue", targetValue)
	req.Put("Body", body)

	r := &PushMessageToiOSRequest{Request: req}
	return r
}
