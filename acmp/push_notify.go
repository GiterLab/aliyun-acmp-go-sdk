package acmp

import (
	"encoding/json"
	"errors"
)

type PushNotify2AndroidRequest struct {
	Request *Request
}

func (p *PushNotify2AndroidRequest) DoActionWithException()(resp *PushNotifyResponse,err error){
	if p!=nil&&p.Request!=nil {
		resp:=&PushNotifyResponse{}
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
	return nil,errors.New("SendRequest is nil")
}

func PushNotify2Android(target,targetValue,title,body string) *PushMessge2IosRequest {
	if target==""||targetValue=="" {
		return nil
	}
	req:=newRequset()
	req.Put("Version", "2016-08-01")
	req.Put("Action", "PushMessageToIos")
	req.Put("Target",target)
	req.Put("TargetValue",targetValue)
	req.Put("Title",title)
	req.Put("Body",body)

	r:=&PushMessge2IosRequest{Request:req}
	return r
}

func (p *PushNotify2AndroidRequest)SetPushExtParameters(extParameters map[string]interface{}) *PushNotify2AndroidRequest {
	body,_:=json.Marshal(extParameters)
	p.Request.Put("ExtParameters",string(body))
	return p
}