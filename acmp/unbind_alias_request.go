package acmp

import (
	"encoding/json"
	"errors"
	"strconv"
)

// UnbindAliasResponse 返回响应体
type UnbindAliasResponse struct {
	ErrorMessage
}

// String 序列化响应
func (l *UnbindAliasResponse) String() string {
	body, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(body)
}

// UnbindAliasRequest 解绑别名请求
type UnbindAliasRequest struct {
	Request *Request
}

// SetUnbindAll 设置是否解绑所有别名，可选
func (u *UnbindAliasRequest) SetUnbindAll(unbindAll bool) *UnbindAliasRequest {
	if u == nil || u.Request == nil {
		return nil
	}
	if unbindAll {
		u.Request.Put("UnbindAll", "true")
	} else {
		u.Request.Put("UnbindAll", "false")
	}
	return u
}

// SetAliasName 设置解绑别名，可选
func (u *UnbindAliasRequest) SetAliasName(aliasName string) *UnbindAliasRequest {
	if u == nil || u.Request == nil {
		return nil
	}
	u.Request.Put("AliasName", aliasName)
	return u
}

// DoActionWithException 发起http请求
func (u *UnbindAliasRequest) DoActionWithException() (resp *UnbindAliasResponse, err error) {
	if u != nil && u.Request != nil {
		resp := &UnbindAliasResponse{}
		body, httpCode, err := u.Request.Do("UnbindAlias")
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

// UnbindAlias 解除绑定别名请求接口
func UnbindAlias(appKey int, deviceID string) *UnbindAliasRequest {
	req := newRequset()
	req.Put("Version", "2016-08-01")
	req.Put("Action", "UnbindTag")
	req.Put("AppKey", strconv.Itoa(appKey))
	req.Put("DeviceId", deviceID)

	r := &UnbindAliasRequest{Request: req}
	return r
}
