package acmp

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

type AppDeviceStat struct {
	Time       string `json:"Time"`
	Count      int `json:"Count"`
	DeviceType string `json:"DeviceType"`
}

type AppDeviceStats struct {
	AppDeviceStats []*AppDeviceStat `json:"AppDeviceStats"`
}

type QueryDeviceStatResponse struct {
	ErrorMessage
	AppDeviceStats *AppDeviceStats `json:"AppDeviceStats,omitempty"`
}

func (q *QueryDeviceStatResponse) GetAppDeviceStats() *AppDeviceStats {
	if q != nil && q.AppDeviceStats != nil {
		return q.AppDeviceStats
	}
	return nil
}

func (q *QueryDeviceStatResponse) String() string {
	body, err := json.Marshal(q)
	if err != nil {
		return ""
	}
	return string(body)
}

type QueryDeviceStatRequest struct {
	Request *Request
}

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
