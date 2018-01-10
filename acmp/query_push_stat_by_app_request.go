package acmp

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

type AppPushStat struct {
	Time string
	AcceptCount int
	SentCount int
	ReceivedCount int
	OpenedCount int
	DeletedCount int
}

type AppPushStats struct {
	AppPushStats []*AppPushStat
}

type QueryPushStatByAppResponse struct {
	ErrorMessage
	AppPushStats *AppPushStats
}

func (q *QueryPushStatByAppResponse)GetAppPushStats() *AppPushStats {
	if q!=nil&&q.AppPushStats!=nil {
		return q.AppPushStats
	}
	return nil
}

func (q *QueryPushStatByAppResponse)String() string {
	body,err:=json.Marshal(q)
	if err!=nil {
		return ""
	}
	return string(body)
}

type QueryPushStatByAppRequest struct {
	Request *Request
}

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

func QueryPushStatByApp(appKey int,startTime,endTime int64,granularity string) *QueryPushStatByAppRequest {
	req:=&Request{}
	req.Put("Version", "2016-08-01")
	req.Put("Action", "QueryPushStatByApp")
	req.Put("AppKey", strconv.Itoa(appKey))
	req.Put("StartTime", time.Unix(startTime, 0).Format("2006-01-02T15:04:05Z"))
	req.Put("EndTime", time.Unix(endTime, 0).Format("2006-01-02T15:04:05Z"))
	req.Put("Granularity",granularity)

	r:=&QueryPushStatByAppRequest{Request:req}
	return r
}


