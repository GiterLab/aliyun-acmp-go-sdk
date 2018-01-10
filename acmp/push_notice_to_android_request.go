package acmp

import (
	"encoding/json"
	"errors"
	"strconv"
)

// PushNoticeToAndroidRequest 推送通知到android平台请求结构体
type PushNoticeToAndroidRequest struct {
	Request *Request
}

// DoActionWithException 发起http请求
func (p *PushNoticeToAndroidRequest) DoActionWithException() (resp *PushNoticeResponse, err error) {
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

// PushNoticeToAndroid 推送通知到android平台接口
func PushNoticeToAndroid(appKey int, target, targetValue, title, body string) *PushNoticeToAndroidRequest {
	if target == "" || targetValue == "" {
		return nil
	}
	req := newRequset()
	req.Put("Version", "2016-08-01")
	req.Put("Action", "PushMessageToIos")
	req.Put("AppKey", strconv.Itoa(appKey))
	req.Put("Target", target)
	req.Put("TargetValue", targetValue)
	req.Put("Title", title)
	req.Put("Body", body)

	r := &PushNoticeToAndroidRequest{Request: req}
	return r
}

// SetPushExtParameters 设置推送额外的参数，可选
func (p *PushNoticeToAndroidRequest) SetPushExtParameters(extParameters map[string]interface{}) *PushNoticeToAndroidRequest {
	if p == nil || p.Request == nil {
		return nil
	}
	body, _ := json.Marshal(extParameters)
	p.Request.Put("ExtParameters", string(body))
	return p
}
