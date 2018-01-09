package acmp

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

// MessageInfo 消息结构体
type MessageInfo struct {
	MessageId  string
	Type       int
	Status     int
	Title      string
	Body       string
	Summary    string
	AppName    string
	AppKey     int
	DeviceType int
	PushTime   int64
}

// SummaryMessageInfos 消息结构体
type SummaryMessageInfos struct {
	Total        int
	Page         int
	PageSize     int
	MessageInfos []*MessageInfo
}

// ListSummaryPushResponse 请求响应
type ListSummaryPushResponse struct {
	ErrorMessage
	SummaryMessageInfos *SummaryMessageInfos
}

// GetSummaryMessageInfos 获取请求结构体的get方法
func (l *ListSummaryPushResponse) GetSummaryMessageInfos() *SummaryMessageInfos {
	if l != nil && l.SummaryMessageInfos != nil {
		return l.SummaryMessageInfos
	}
	return nil
}

// String 序列化响应
func (l *ListSummaryPushResponse) String() string {
	body, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(body)
}

// ListSummaryPushMessageInfoRequest http请求结构体
type ListSummaryPushMessageInfoRequest struct {
	Request *Request
}

// DoActionWithException 发起http请求
func (l *ListSummaryPushMessageInfoRequest) DoActionWithException() (resp *ListSummaryPushResponse, err error) {
	if l != nil && l.Request != nil {
		resp := &ListSummaryPushResponse{}
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

// ListSummaryPush 查询推送列表接口
func ListSummaryPush(appKey int, pushType string, startTime, endTime int64) *ListSummaryPushMessageInfoRequest {
	req := newRequset()
	req.Put("Version", "2016-08-01")
	req.Put("Action", "ListPushRecords")
	req.Put("AppKey", strconv.Itoa(appKey))
	req.Put("StartTime", time.Unix(startTime, 0).Format("YYYY-MM-DDThh:mm:ssZ"))
	req.Put("EndTime", time.Unix(endTime, 0).Format("YYYY-MM-DDThh:mm:ssZ"))
	req.Put("PushType", pushType)

	r := &ListSummaryPushMessageInfoRequest{Request: req}
	return r
}

func (l *ListSummaryPushMessageInfoRequest) SetPage(page int) *ListSummaryPushMessageInfoRequest {
	if l == nil || l.Request == nil {
		return nil
	}
	l.Request.Put("Page", strconv.Itoa(page))

	return l
}

func (l *ListSummaryPushMessageInfoRequest) SetPageSize(pageSize int) *ListSummaryPushMessageInfoRequest {
	if l == nil || l.Request == nil {
		return nil
	}
	l.Request.Put("PageSize", strconv.Itoa(pageSize))

	return l
}
