package acmp

import (
	"encoding/json"
	"errors"
	"strconv"
)

type MsgPushStat struct {
	MessageId     string
	AcceptCount   int
	SentCount     int
	ReceivedCount int
	OpenedCount   int
	DeletedCount  int
}

type MsgPushStats struct {
	MsgPushStats []*MsgPushStat
}

type QueryPushStatByMsgResponse struct {
	ErrorMessage
	MsgPushStats *MsgPushStats
}

func (q *QueryPushStatByMsgResponse) GetMsgPushStats() *MsgPushStats {
	if q != nil && q.MsgPushStats != nil {
		return q.MsgPushStats
	}
	return nil
}

func (q *QueryPushStatByMsgResponse) String() string {
	body, err := json.Marshal(q)
	if err != nil {
		return ""
	}
	return string(body)
}

type QueryPushStatByMsgRequest struct {
	Request *Request
}

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

func QueryPushStatByMsg(appKey int, messageId string) *QueryPushStatByMsgRequest {
	req := newRequset()
	req.Put("Version", "2016-08-01")
	req.Put("Action", "QueryPushStatByMsg")
	req.Put("AppKey", strconv.Itoa(appKey))
	req.Put("MessageId", messageId)

	r := &QueryPushStatByMsgRequest{Request: req}
	return r
}
