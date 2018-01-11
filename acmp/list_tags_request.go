package acmp

import (
	"encoding/json"
	"errors"
	"strconv"
)

// ListTagsRespones 查询标签列表信息响应体
type ListTagsRespones struct {
	ErrorMessage
	TagInfos *TagInfos `json:"TagInfos,omitempty"`
}

// ListTagsRequest http请求结构体
type ListTagsRequest struct {
	Request *Request
}

// GetTagInfos 获取响应体里面标签列表信息
func (l *ListTagsRespones) GetTagInfos() *TagInfos {
	if l != nil && l.TagInfos != nil {
		return l.TagInfos
	}
	return nil
}

// String 序列化响应
func (l *ListTagsRespones) String() string {
	body, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(body)
}

// DoActionWithException 发起http请求
func (q *ListTagsRequest) DoActionWithException() (resp *ListTagsRespones, err error) {
	if q != nil && q.Request != nil {
		resp := &ListTagsRespones{}
		body, httpCode, err := q.Request.Do("BindTag")
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

// ListTags 查看标签列表请求接口
func ListTags(appKey int) *ListTagsRequest {
	req := newRequset()
	req.Put("Version", "2016-08-01")
	req.Put("Action", "BindTag")
	req.Put("AppKey", strconv.Itoa(appKey))

	r := &ListTagsRequest{Request: req}
	return r
}
