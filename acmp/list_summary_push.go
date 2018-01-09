package acmp

import (
	"encoding/json"
	"errors"
	"strconv"
)

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

type SummaryMessageInfos struct {
	Total        int
	Page         int
	PageSize     int
	MessageInfos []*MessageInfo
}

type ListSummaryPushResponse struct {
	ErrorMessage
	SummaryMessageInfos *SummaryMessageInfos
}

func (l *ListSummaryPushResponse) GetSummaryMessageInfos() *SummaryMessageInfos {
	if l != nil && l.SummaryMessageInfos != nil {
		return l.SummaryMessageInfos
	}
	return nil
}

func (l *ListSummaryPushResponse) String() string {
	body, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(body)
}

type ListSummaryPushMessageInfoRequest struct {
	Request *Request
}

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

func ListSummaryPush(startTime, endTime, pushType string) *ListSummaryPushMessageInfoRequest {
	req := newRequset()
	req.Put("Version", "2016-08-01")
	req.Put("Action", "ListPushRecords")
	req.Put("StartTime", startTime)
	req.Put("EndTime", endTime)
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
