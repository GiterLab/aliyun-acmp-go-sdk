package acmp

import (
	"encoding/json"
	"errors"
)

type PushMessge2AndroidRequest struct {
	Request *Request
}

func (p *PushMessge2AndroidRequest) DoActionWithException()(resp *PushMessageResponse,err error){
	if p!=nil&&p.Request!=nil {
		resp:=&PushMessageResponse{}
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
	return nil,errors.New("SendRequest is nil")
}

func PushMessage2Android(target,targetValue,title,body string) *PushMessge2AndroidRequest {
	if target==""||targetValue=="" {
		return nil
	}
	req:=newRequset()
	req.Put("Version", "2016-08-01")
	req.Put("Action", "PushMessageToAndroid")
	req.Put("Target",target)
	req.Put("TargetValue",targetValue)
	req.Put("Title",title)
	req.Put("Body",body)

	r:=&PushMessge2AndroidRequest{Request:req}
	return r
}