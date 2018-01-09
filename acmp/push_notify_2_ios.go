package acmp

import (
	"encoding/json"
	"errors"
)

type PushNotifyResponse struct {
	ErrorMessage
	MessageId string `json:"message_id"`
}

func (p *PushNotifyResponse) getMessageId() string {
	if p != nil && p.MessageId != "" {
		return p.MessageId
	}
	return ""
}

func (p *PushNotifyResponse) String() string {
	body, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(body)
}

type PushNotify2IosRequest struct {
	Request *Request
}

func (p *PushNotify2IosRequest) DoActionWithException() (resp *PushNotifyResponse, err error) {
	if p != nil && p.Request != nil {
		resp := &PushNotifyResponse{}
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

func PushNotify2Ios(target, targetValue, apnsEnv, body string) *PushMessge2IosRequest {
	if target == "" || targetValue == "" {
		return nil
	}
	req := newRequset()
	req.Put("Version", "2016-08-01")
	req.Put("Action", "PushMessageToIos")
	req.Put("Target", target)
	req.Put("TargetValue", targetValue)
	req.Put("ApnsEnv", apnsEnv)
	req.Put("Body", body)

	r := &PushMessge2IosRequest{Request: req}
	return r
}

func (p *PushMessge2IosRequest) SetPushTitle(title string) *PushMessge2IosRequest {
	p.Request.Put("Title", title)
	return p
}

func (p *PushMessge2IosRequest) SetPushExtParameters(extParameters map[string]interface{}) *PushMessge2IosRequest {
	if p == nil || p.Request == nil {
		return nil
	}
	body, _ := json.Marshal(extParameters)
	p.Request.Put("ExtParameters", string(body))
	return p
}
