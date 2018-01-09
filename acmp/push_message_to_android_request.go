package acmp

import (
	"encoding/json"
	"errors"
	"strconv"
)

// PushMessageToAndroidRequest 推送消息结构体
type PushMessageToAndroidRequest struct {
	Request *Request
}

// DoActionWithException 发起请求
func (p *PushMessageToAndroidRequest) DoActionWithException() (resp *PushMessageResponse, err error) {
	if p != nil && p.Request != nil {
		resp := &PushMessageResponse{}
		body, httpCode, err := p.Request.Do("PushMessageToAndroid")
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

// PushMessage2Android 推送安卓消息接口
func PushMessage2Android(appKey int, target, targetValue, title, body string) *PushMessageToAndroidRequest {
	if target == "" || targetValue == "" {
		return nil
	}
	req := newRequset()
	req.Put("Version", "2016-08-01")
	req.Put("Action", "PushMessageToAndroid")
	req.Put("AppKey", strconv.Itoa(appKey))
	req.Put("Target", target)
	req.Put("TargetValue", targetValue)
	req.Put("Title", title)
	req.Put("Body", body)

	r := &PushMessageToAndroidRequest{Request: req}
	return r
}
