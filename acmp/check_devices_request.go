package acmp

import (
	"encoding/json"
	"errors"
	"strconv"
)

// DeviceCheckInfo 检查设备信息
type DeviceCheckInfo struct {
	DeviceID  string `json:"DeviceId"`
	Available string `json:"Available"`
}

// DeviceCheckInfos 检查设备信息
type DeviceCheckInfos struct {
	DeviceCheckInfos []DeviceCheckInfo `json:"DeviceCheckInfos"`
}

// CheckDevicesResponse 检查设备响应
type CheckDevicesResponse struct {
	ErrorMessage
	DeviceCheckInfos *DeviceCheckInfos `json:"DeviceCheckInfos,omitempty"`
}

// GetDeviceInfos 获取响应里面的设备信息
func (c *CheckDevicesResponse) GetDeviceInfos() *DeviceCheckInfos {
	if c != nil && c.DeviceCheckInfos != nil {
		return c.DeviceCheckInfos
	}
	return nil
}

// String 序列化响应信息
func (c *CheckDevicesResponse) String() string {
	body, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(body)
}

// CheckDevicesRequest http请求结构体
type CheckDevicesRequest struct {
	Request *Request
}

// DoActionWithException 发起http请求
func (c *CheckDevicesRequest) DoActionWithException() (*CheckDevicesResponse, error) {
	if c != nil && c.Request != nil {
		resp := &CheckDevicesResponse{}
		body, httpCode, err := c.Request.Do("CheckDevices")
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

// CheckDevices 设备检查接口
//
func CheckDevices(appKey int, deviceIDs string) *CheckDevicesRequest {
	req := newRequset()
	req.Put("Version", "2016-08-01")
	req.Put("Action", "QueryDeviceInfo")
	req.Put("AppKey", strconv.Itoa(appKey))
	req.Put("DeviceIds", deviceIDs)

	r := &CheckDevicesRequest{Request: req}
	return r
}
