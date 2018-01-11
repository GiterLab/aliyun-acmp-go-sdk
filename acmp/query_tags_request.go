package acmp

import (
	"encoding/json"
	"errors"
	"strconv"
)

// TagInfo 标签信息结构体
type TagInfo struct {
	TagName string `json:"TagName"`
}

// TagInfos 标签信息
type TagInfos struct {
	TagInfos []TagInfo `json:"TagInfos"`
}

// QueryTagsRespones 查询标签信息响应体
type QueryTagsRespones struct {
	ErrorMessage
	TagInfos *TagInfos `json:"TagInfos,omitempty"`
}

// GetTagInfos 获取响应体内标签信息
func (q *QueryTagsRespones) GetTagInfos() *TagInfos {
	if q != nil && q.TagInfos != nil {
		return q.TagInfos
	}
	return nil
}

// String 序列化响应
func (q *QueryTagsRespones) String() string {
	body, err := json.Marshal(q)
	if err != nil {
		return ""
	}
	return string(body)
}

// QueryTagsRequest http请求结构体
type QueryTagsRequest struct {
	Request *Request
}

// DoActionWithException 发起http请求
func (q *QueryTagsRequest) DoActionWithException() (resp *QueryTagsRespones, err error) {
	if q != nil && q.Request != nil {
		resp := &QueryTagsRespones{}
		body, httpCode, err := q.Request.Do("QueryTags")
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

// QueryTags 查询绑定标签请求接口
func QueryTags(appKey int, clientKey, keyType string) *QueryTagsRequest {
	req := newRequset()
	req.Put("Version", "2016-08-01")
	req.Put("Action", "QueryTags")
	req.Put("AppKey", strconv.Itoa(appKey))
	req.Put("ClientKey", clientKey)
	req.Put("KeyType", keyType)

	r := &QueryTagsRequest{Request: req}
	return r
}
