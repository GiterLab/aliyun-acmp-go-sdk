package acmp

import (
	"encoding/json"
	"errors"
)

type CancelPushResponse struct {
	ErrorMessage
}

func (c *CancelPushResponse) String() string {
	body, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(body)
}

type CancelPushRequest struct {
	Request *Request
}

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

func CancelPush(messageId string) *CancelPushRequest {
	if messageId == "" {
		return nil
	}
	req := newRequset()
	req.Put("Version", "2016-08-01")
	req.Put("Action", "CancelPush")
	req.Put("MessageId", messageId)

	r := &CancelPushRequest{Request: req}
	return r
}
