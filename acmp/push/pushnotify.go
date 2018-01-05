package push

import (
	"aliyun-acmp-go-sdk/acmp/signature"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Notify struct {
	RootUrl      string
	AccessSecret string
	PublicParam  *PublicParam
	NoticeParam  *NoticeParam
}

func (m *Notify) SetRootUrl(rootUrl string) {
	m.RootUrl = rootUrl
}

func (m *Notify) SetAccessSecret(accessSecret string) {
	m.AccessSecret = accessSecret
}

func (m *Notify) SetPublicParam(publicParam *PublicParam) {
	m.PublicParam = publicParam
}

func (m *Notify) SetNoticeParam(noticeParam *NoticeParam) {
	m.NoticeParam = noticeParam
}

func (n *Notify) DoACMP() (responeString string, err error) {
	if n.RootUrl == "" || n.AccessSecret == "" || n.PublicParam == nil || n.NoticeParam == nil {
		return "", errors.New("PushNotice param pointer shouldn't be nil")
	}
	publicParamStr, _ := n.PublicParam.ToStringWithoutSignature()
	noticeParamStr, _ := n.NoticeParam.ToString()
	urlstr := n.RootUrl + "/?" + publicParamStr + noticeParamStr
	signstr, err := signature.SignatureString(urlstr, http.MethodGet)
	if err != nil {
		return "", errors.New("PushNotice signature.SignatureString err")
	}
	signaturestr, err := signature.GetSignature(signstr, n.AccessSecret)
	if err != nil {
		return "", errors.New("PushNotice signature.GetSignature err")
	}
	finalUrlStr := urlstr + "&Signature=" + signaturestr
	fmt.Println("finalUrlStr---->:", finalUrlStr)
	resp, err := http.Get(finalUrlStr)
	if err != nil {
		return "", errors.New("PushNotice http.Get err")
	}
	resultByte, _ := ioutil.ReadAll(resp.Body)
	return string(resultByte), nil
}
