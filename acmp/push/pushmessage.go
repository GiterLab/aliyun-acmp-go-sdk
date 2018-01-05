package push

import (
	"aliyun-acmp-go-sdk/acmp/bean"
	"aliyun-acmp-go-sdk/acmp/signature"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Message struct {
	RootUrl      string
	AccessSecret string
	PublicParam  *bean.PublicParam
	MessageParam *bean.MessageParam
}

func (m *Message) SetRootUrl(rootUrl string) {
	m.RootUrl = rootUrl
}

func (m *Message) SetAccessSecret(accessSecret string) {
	m.AccessSecret = accessSecret
}

func (m *Message) SetPublicParam(publicParam *bean.PublicParam) {
	m.PublicParam = publicParam
}

func (m *Message) SetMessageParam(messageParam *bean.MessageParam) {
	m.MessageParam = messageParam
}

func (m *Message) DoPush(message *Message) (responeString string, err error) {
	if message.RootUrl == "" || message.AccessSecret == "" || message.PublicParam == nil || message.MessageParam == nil {
		return "", errors.New("PushMessage param shouldn't be nil or null")
	}
	PublicParamStr, _ := message.PublicParam.ToStringWithoutSignature()
	noticeParamStr, _ := message.MessageParam.ToString()
	urlstr := message.RootUrl + "/?" + PublicParamStr + noticeParamStr
	signstr, err := signature.SignatureString(urlstr, http.MethodGet)
	if err != nil {
		return "", errors.New("PushMessage signature.SignatureString err")
	}
	signaturestr, err := signature.GetSignature(signstr, message.AccessSecret)
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
