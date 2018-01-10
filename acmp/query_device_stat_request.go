package acmp

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

// AppDeviceStat 设备新增与留存
type AppDeviceStat struct {
	Time       string `json:"Time"`
	Count      int    `json:"Count"`
	DeviceType string `json:"DeviceType,omitempty"`
}

// AppDeviceStats 设备新增与留存结构体
type AppDeviceStats struct {
	AppDeviceStats []AppDeviceStat `json:"AppDeviceStats"`
}

// QueryDeviceStatResponse 响应结构体
type QueryDeviceStatResponse struct {
	ErrorMessage
	AppDeviceStats *AppDeviceStats `json:"AppDeviceStats,omitempty"`
}

// GetAppDeviceStats 获取响应的设备新增与留存信息
func (q *QueryDeviceStatResponse) GetAppDeviceStats() *AppDeviceStats {
	if q != nil && q.AppDeviceStats != nil {
		return q.AppDeviceStats
	}
	return nil
}

// String 序列化
func (q *QueryDeviceStatResponse) String() string {
	body, err := json.Marshal(q)
	if err != nil {
		return ""
	}
	return string(body)
}

// QueryDeviceStatRequest 请求结构体
type QueryDeviceStatRequest struct {
	Request *Request
}

// DoActionWithException 发起http请求
func (q *QueryDeviceStatRequest) DoActionWithException() (*QueryDeviceStatResponse, error) {
	if q != nil && q.Request != nil {
		resp := &QueryDeviceStatResponse{}
		body, httpCode, err := q.Request.Do("QueryDeviceStat")
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

// QueryDeviceStat 发起http请求接口
func QueryDeviceStat(appKey int, startTime, endTime int64, deviceType, queryType string) *QueryDeviceStatRequest {
	req := newRequset()
	req.Put("Version", "2016-08-01")
	req.Put("Action", "QueryDeviceStat")
	req.Put("AppKey", strconv.Itoa(appKey))
	req.Put("StartTime", time.Unix(startTime, 0).Format("2006-01-02T15:04:05Z"))
	req.Put("EndTime", time.Unix(endTime, 0).Format("2006-01-02T15:04:05Z"))
	req.Put("DeviceType", deviceType)
	req.Put("QueryType", queryType)

	r := &QueryDeviceStatRequest{Request: req}
	return r
}
