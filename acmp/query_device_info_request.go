package acmp

import (
	"encoding/json"
	"errors"
	"strconv"
)

// DeviceInfo 设备信息结构体
type DeviceInfo struct {
	BoundAccount   string `json:"BoundAccount"`
	BoundTag       string `json:"BoundTag"`
	BoundAlias     string `json:"BoundAlias"`
	DeviceID       string `json:"DeviceId"`
	DeviceToken    string `json:"DeviceToken"`
	DeviceType     int    `json:"DeviceType"`
	IsOnline       bool   `json:"IsOnline"`
	LastOnlineTime string `json:"LastOnlineTime"`
}

// QueryDeviceInfoResponse 查询设备信息响应
type QueryDeviceInfoResponse struct {
	ErrorMessage
	DeviceInfo *DeviceInfo `json:"DeviceInfo,omitempty"`
}

// GetDeviceInfo 获取响应里面的设备信息
func (q *QueryDeviceInfoResponse) GetDeviceInfo() *DeviceInfo {
	if q != nil && q.DeviceInfo != nil {
		return q.DeviceInfo
	}
	return nil
}

// String 序列化响应
func (q *QueryDeviceInfoResponse) String() string {
	body, err := json.Marshal(q)
	if err != nil {
		return ""
	}
	return string(body)
}

// QueryDeviceInfoRequest http请求结构体
type QueryDeviceInfoRequest struct {
	Request *Request
}

// DoActionWithException 发起http请求
func (q *QueryDeviceInfoRequest) DoActionWithException() (*QueryDeviceInfoResponse, error) {
	if q != nil && q.Request != nil {
		resp := &QueryDeviceInfoResponse{}
		body, httpCode, err := q.Request.Do("QueryDeviceInfo")
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

// QueryDeviceInfo 查询设备信息接口
func QueryDeviceInfo(appKey int, deviceID string) *QueryDeviceInfoRequest {
	req := newRequset()
	req.Put("Version", "2016-08-01")
	req.Put("Action", "QueryDeviceInfo")
	req.Put("AppKey", strconv.Itoa(appKey))
	req.Put("DeviceId", deviceID)

	r := &QueryDeviceInfoRequest{Request: req}
	return r
}
