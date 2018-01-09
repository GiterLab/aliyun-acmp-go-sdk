// Package acmp Copyright 2016 The GiterLab Authors. All rights reserved.
package acmp

import (
	"encoding/json"
	"errors"
)

// SummaryAppInfo APP概览列表信息
type SummaryAppInfo struct {
	AppKey   int    `json:"AppKey"`   // "AppKey": "129376929"
	AppName  string `json:"AppName"`  // "AppName": "移动推送APP1"
	Platform string `json:"Platform"` // "Platform": "Android"
}

// SummaryAppInfos  APP概览列表
type SummaryAppInfos struct {
	SummaryAppInfos []SummaryAppInfo `json:"SummaryAppInfo"`
}

// ListSummaryAppsResponse APP概览列表响应参数
type ListSummaryAppsResponse struct {
	ErrorMessage
	SummaryAppInfos *SummaryAppInfos `json:"SummaryAppInfos,omitempty"`
}

// GetSummaryAppInfos 获取APP概览列表
func (l *ListSummaryAppsResponse) GetSummaryAppInfos() *SummaryAppInfos {
	if l != nil && l.SummaryAppInfos != nil {
		return l.SummaryAppInfos
	}
	return nil
}

// String 序列化成JSON字符串
func (l ListSummaryAppsResponse) String() string {
	body, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(body)
}

// ListSummaryAppsRequest APP概览列表请求参数
type ListSummaryAppsRequest struct {
	Request *Request
}

// DoActionWithException 发起HTTP请求
func (l *ListSummaryAppsRequest) DoActionWithException() (resp *ListSummaryAppsResponse, err error) {
	if l != nil && l.Request != nil {
		resp := &ListSummaryAppsResponse{}
		body, httpCode, err := l.Request.Do("ListSummaryApps")
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

// ListSummaryApps APP概览列表接口
func ListSummaryApps() *ListSummaryAppsRequest {
	req := newRequset()
	req.Put("Version", "2016-08-01")
	req.Put("Action", "ListSummaryApps")

	r := &ListSummaryAppsRequest{Request: req}
	return r
}
