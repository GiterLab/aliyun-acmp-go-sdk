package acmp

import (
	"encoding/json"
	"errors"
)

type PushMessageResponse struct {
	ErrorMessage
	MessageId string `json:"message_id"`
}

func (p *PushMessageResponse) getMessageId() string {
	if p != nil && p.MessageId != "" {
		return p.MessageId
	}
	return ""
}

func (p *PushMessageResponse) String() string {
	body, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(body)
}

type PushMessge2IosRequest struct {
	Request *Request
}

func (p *PushMessge2IosRequest) DoActionWithException() (resp *PushMessageResponse, err error) {
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

func PushMessage2Ios(target, targetValue, body string) *PushMessge2IosRequest {
	if target == "" || targetValue == "" {
		return nil
	}
	req := newRequset()
	req.Put("Version", "2016-08-01")
	req.Put("Action", "PushMessageToIos")
	req.Put("Target", target)
	req.Put("TargetValue", targetValue)
	req.Put("Body", body)

	r := &PushMessge2IosRequest{Request: req}
	return r
}
