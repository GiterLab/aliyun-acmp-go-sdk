package acmp

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

// AppPushStat APP维度推送统计结构体
type AppPushStat struct {
	Time          string `json:"Time"`
	AcceptCount   int    `json:"AcceptCount"`
	SentCount     int    `json:"SentCount"`
	ReceivedCount int    `json:"ReceivedCount"`
	OpenedCount   int    `json:"OpenedCount"`
	DeletedCount  int    `json:"DeletedCount"`
}

// AppPushStats APP维度推送统计结构体
type AppPushStats struct {
	AppPushStats []AppPushStat `json:"AppPushStats"`
}

// QueryPushStatByAppResponse 响应结构体
type QueryPushStatByAppResponse struct {
	ErrorMessage
	AppPushStats *AppPushStats `json:"AppPushStats,omitempty"`
}

// GetAppPushStats 获取响应体的APP维度推送统计结构体
func (q *QueryPushStatByAppResponse) GetAppPushStats() *AppPushStats {
	if q != nil && q.AppPushStats != nil {
		return q.AppPushStats
	}
	return nil
}

// String 序列化
func (q *QueryPushStatByAppResponse) String() string {
	body, err := json.Marshal(q)
	if err != nil {
		return ""
	}
	return string(body)
}

// QueryPushStatByAppRequest http请求结构体
type QueryPushStatByAppRequest struct {
	Request *Request
}

// DoActionWithException 发起http请求
func (q *QueryPushStatByAppRequest) DoActionWithException() (resp *QueryPushStatByAppResponse, err error) {
	if q != nil && q.Request != nil {
		resp := &QueryPushStatByAppResponse{}
		body, httpCode, err := q.Request.Do("QueryPushStatByApp")
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

// QueryPushStatByApp http请求接口
func QueryPushStatByApp(appKey int, startTime, endTime int64, granularity string) *QueryPushStatByAppRequest {
	req := newRequset()
	req.Put("Version", "2016-08-01")
	req.Put("Action", "QueryPushStatByApp")
	req.Put("AppKey", strconv.Itoa(appKey))
	req.Put("StartTime", time.Unix(startTime, 0).Format("2006-01-02T15:04:05Z"))
	req.Put("EndTime", time.Unix(endTime, 0).Format("2006-01-02T15:04:05Z"))
	req.Put("Granularity", granularity)

	r := &QueryPushStatByAppRequest{Request: req}
	return r
}
