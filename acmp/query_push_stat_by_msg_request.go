package acmp

import (
	"encoding/json"
	"errors"
	"strconv"
)

// MsgPushStat 任务维度推送统计结构体
type MsgPushStat struct {
	MessageID     string `json:"MessageId"`
	AcceptCount   int    `json:"AcceptCount"`
	SentCount     int    `json:"SentCount"`
	ReceivedCount int    `json:"ReceivedCount"`
	OpenedCount   int    `json:"OpenedCount"`
	DeletedCount  int    `json:"DeletedCount"`
}

// MsgPushStats 任务维度推送统计结构体
type MsgPushStats struct {
	MsgPushStats []MsgPushStat `json:"MsgPushStats"`
}

// QueryPushStatByMsgResponse 响应结构体
type QueryPushStatByMsgResponse struct {
	ErrorMessage
	MsgPushStats *MsgPushStats `json:"MsgPushStats,omitempty"`
}

// GetMsgPushStats 获取响应结构体的任务维度推送统计结构体
func (q *QueryPushStatByMsgResponse) GetMsgPushStats() *MsgPushStats {
	if q != nil && q.MsgPushStats != nil {
		return q.MsgPushStats
	}
	return nil
}

// String 序列化
func (q *QueryPushStatByMsgResponse) String() string {
	body, err := json.Marshal(q)
	if err != nil {
		return ""
	}
	return string(body)
}

// QueryPushStatByMsgRequest http请求结构体
type QueryPushStatByMsgRequest struct {
	Request *Request
}

// DoActionWithException http请求
func (q *QueryPushStatByMsgRequest) DoActionWithException() (resp *QueryPushStatByMsgResponse, err error) {
	if q != nil && q.Request != nil {
		resp := &QueryPushStatByMsgResponse{}
		body, httpCode, err := q.Request.Do("QueryPushStatByMsg")
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

// QueryPushStatByMsg http请求接口
func QueryPushStatByMsg(appKey int, messageID string) *QueryPushStatByMsgRequest {
	req := newRequset()
	req.Put("Version", "2016-08-01")
	req.Put("Action", "QueryPushStatByMsg")
	req.Put("AppKey", strconv.Itoa(appKey))
	req.Put("MessageId", messageID)

	r := &QueryPushStatByMsgRequest{Request: req}
	return r
}
