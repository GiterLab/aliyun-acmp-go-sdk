package acmp

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

// MessageInfo 消息结构体
type MessageInfo struct {
	MessageID  string `json:"MessageId"`
	Type       int    `json:"Type"`
	Status     int    `json:"Status"`
	Title      string `json:"Title"`
	Body       string `json:"Body"`
	Summary    string `json:"Summary"`
	AppName    string `json:"AppName"`
	AppKey     int    `json:"AppKey"`
	DeviceType int    `json:"DeviceType"`
	PushTime   int64  `json:"PushTime"`
}

// SummaryMessageInfos 消息结构体
type SummaryMessageInfos struct {
	Total        int            `json:"Total"`
	Page         int            `json:"Page"`
	PageSize     int            `json:"PageSize"`
	MessageInfos []*MessageInfo `json:"MessageInfos"`
}

// ListPushRecordsResponse 请求响应
type ListPushRecordsResponse struct {
	ErrorMessage
	SummaryMessageInfos *SummaryMessageInfos `json:"SummaryMessageInfos,omitempty"`
}

// GetSummaryMessageInfos 获取请求结构体的get方法
func (l *ListPushRecordsResponse) GetSummaryMessageInfos() *SummaryMessageInfos {
	if l != nil && l.SummaryMessageInfos != nil {
		return l.SummaryMessageInfos
	}
	return nil
}

// String 序列化响应
func (l *ListPushRecordsResponse) String() string {
	body, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(body)
}

// ListPushRecordsRequest http请求结构体
type ListPushRecordsRequest struct {
	Request *Request
}

// DoActionWithException 发起http请求
func (l *ListPushRecordsRequest) DoActionWithException() (resp *ListPushRecordsResponse, err error) {
	if l != nil && l.Request != nil {
		resp := &ListPushRecordsResponse{}
		body, httpCode, err := l.Request.Do("ListPushRecords")
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

// ListPushRecords 查询推送列表接口
func ListPushRecords(appKey int, pushType string, startTime, endTime int64) *ListPushRecordsRequest {
	req := newRequset()
	req.Put("Version", "2016-08-01")
	req.Put("Action", "ListPushRecords")
	req.Put("AppKey", strconv.Itoa(appKey))
	req.Put("StartTime", time.Unix(startTime, 0).Format("2006-01-02T15:04:05Z"))
	req.Put("EndTime", time.Unix(endTime, 0).Format("2006-01-02T15:04:05Z"))
	req.Put("PushType", pushType)

	r := &ListPushRecordsRequest{Request: req}
	return r
}

// SetPage 设置请求起始页码
func (l *ListPushRecordsRequest) SetPage(page int) *ListPushRecordsRequest {
	if l == nil || l.Request == nil {
		return nil
	}
	l.Request.Put("Page", strconv.Itoa(page))

	return l
}

// SetPageSize 设置每页大小
func (l *ListPushRecordsRequest) SetPageSize(pageSize int) *ListPushRecordsRequest {
	if l == nil || l.Request == nil {
		return nil
	}
	l.Request.Put("PageSize", strconv.Itoa(pageSize))

	return l
}
