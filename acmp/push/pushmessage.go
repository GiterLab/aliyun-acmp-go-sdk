package push

import (
	"aliyun-acmp-go-sdk/acmp/signature"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Message struct {
	RootUrl      string
	AccessSecret string
	PublicParam  *PublicParam
	MessageParam *MessageParam
}

func (m *Message) SetRootUrl(rootUrl string) {
	m.RootUrl = rootUrl
}

func (m *Message) SetAccessSecret(accessSecret string) {
	m.AccessSecret = accessSecret
}

func (m *Message) SetPublicParam(publicParam *PublicParam) {
	m.PublicParam = publicParam
}

func (m *Message) SetMessageParam(messageParam *MessageParam) {
	m.MessageParam = messageParam
}

func (m *Message) DoACMP() (responeString string, err error) {
	if m.RootUrl == "" || m.AccessSecret == "" || m.PublicParam == nil || m.MessageParam == nil {
		return "", errors.New("PushMessage param shouldn't be nil or null")
	}
	PublicParamStr, _ := m.PublicParam.ToStringWithoutSignature()
	noticeParamStr, _ := m.MessageParam.ToString()
	urlstr := m.RootUrl + "/?" + PublicParamStr + noticeParamStr
	signstr, err := signature.SignatureString(urlstr, http.MethodGet)
	if err != nil {
		return "", errors.New("PushMessage signature.SignatureString err")
	}
	signaturestr, err := signature.GetSignature(signstr, m.AccessSecret)
	if err != nil {
		return "", errors.New("PushMessage signature.GetSignature err")
	}
	finalUrlStr := urlstr + "&Signature=" + signaturestr
	fmt.Println(finalUrlStr)
	resp, err := http.Get(finalUrlStr)
	if err != nil {
		return "", errors.New("PushMessage http.Get err")
	}
	resultByte, _ := ioutil.ReadAll(resp.Body)
	return string(resultByte), nil
}
