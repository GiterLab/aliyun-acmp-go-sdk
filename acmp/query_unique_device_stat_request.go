package acmp

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

// QueryUniqueDeviceStatResponse 查询去重设备返回响应
type QueryUniqueDeviceStatResponse struct {
	ErrorMessage
	AppDeviceStats *AppDeviceStats `json:"AppDeviceStats,omitempty"`
}

// GetAppDeviceStats 获取响应体里面的设备状态
func (q *QueryUniqueDeviceStatResponse) GetAppDeviceStats() *AppDeviceStats {
	if q != nil && q.AppDeviceStats != nil {
		return q.AppDeviceStats
	}
	return nil
}

// String 序列化响应
func (q *QueryUniqueDeviceStatResponse) String() string {
	body, err := json.Marshal(q)
	if err != nil {
		return ""
	}
	return string(body)
}

// QueryUniqueDeviceStatRequest http请求结构体
type QueryUniqueDeviceStatRequest struct {
	Request *Request
}

// DoActionWithException 发起http请求
func (q *QueryUniqueDeviceStatRequest) DoActionWithException() (*QueryUniqueDeviceStatResponse, error) {
	if q != nil && q.Request != nil {
		resp := &QueryUniqueDeviceStatResponse{}
		body, httpCode, err := q.Request.Do("QueryUniqueDeviceStat")
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

// QueryUniqueDeviceStat 获取去重设备接口
func QueryUniqueDeviceStat(appKey int, startTime, endTime int64) *QueryUniqueDeviceStatRequest {
	req := newRequset()
	req.Put("Version", "2016-08-01")
	req.Put("Action", "QueryDeviceStat")
	req.Put("AppKey", strconv.Itoa(appKey))
	req.Put("StartTime", time.Unix(startTime, 0).Format("2006-01-02T15:04:05Z"))
	req.Put("EndTime", time.Unix(endTime, 0).Format("2006-01-02T15:04:05Z"))

	r := &QueryUniqueDeviceStatRequest{Request: req}
	return r
}
