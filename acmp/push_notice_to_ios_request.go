package acmp

import (
	"encoding/json"
	"errors"
	"strconv"
)

// PushNoticeResponse 通知响应结构体
type PushNoticeResponse struct {
	ErrorMessage
	MessageID string `json:"MessageId,omitempty"`
}

// GetMessageID 获取通知响应的messageId
func (p *PushNoticeResponse) GetMessageID() string {
	if p != nil && p.MessageID != "" {
		return p.MessageID
	}
	return ""
}

// String 序列化响应
func (p *PushNoticeResponse) String() string {
	body, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(body)
}

// PushNoticeToiOSRequest 通知请求结构体
type PushNoticeToiOSRequest struct {
	Request *Request
}

// DoActionWithException 发起http请求
func (p *PushNoticeToiOSRequest) DoActionWithException() (resp *PushNoticeResponse, err error) {
	if p != nil && p.Request != nil {
		resp := &PushNoticeResponse{}
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

// PushNoticeToiOS 推送通知到iOS平台接口
func PushNoticeToiOS(appKey int, target, targetValue, apnsEnv, body string) *PushMessageToiOSRequest {
	if target == "" || targetValue == "" {
		return nil
	}
	req := newRequset()
	req.Put("Version", "2016-08-01")
	req.Put("Action", "PushMessageToIos")
	req.Put("AppKey", strconv.Itoa(appKey))
	req.Put("Target", target)
	req.Put("TargetValue", targetValue)
	req.Put("ApnsEnv", apnsEnv)
	req.Put("Body", body)

	r := &PushMessageToiOSRequest{Request: req}
	return r
}

// SetPushTitle 设置通知title，可选
func (p *PushMessageToiOSRequest) SetPushTitle(title string) *PushMessageToiOSRequest {
	if p == nil || p.Request == nil {
		return nil
	}
	p.Request.Put("Title", title)
	return p
}

// SetPushExtParameters 设置通知额外的参数，可选
func (p *PushMessageToiOSRequest) SetPushExtParameters(extParameters map[string]interface{}) *PushMessageToiOSRequest {
	if p == nil || p.Request == nil {
		return nil
	}
	body, _ := json.Marshal(extParameters)
	p.Request.Put("ExtParameters", string(body))
	return p
}
