package push

import (
	"aliyun-acmp-go-sdk/acmp/bean"
	"aliyun-acmp-go-sdk/acmp/signature"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Notify struct {
	RootUrl string
	AccessSecret string
	PublicParam *bean.PublicParam
	NoticeParam *bean.NoticeParam
}

func PushNotice(notify *Notify) (responeString string, err error) {
	if notify.PublicParam == nil || notify.NoticeParam == nil {
		return "", errors.New("PushNotice param pointer shouldn't be nil")
	}
	publicParamStr, _ := notify.PublicParam.ToStringWithoutSignature()
	noticeParamStr, _ := notify.NoticeParam.ToString()
	urlstr := notify.RootUrl + "/?" + publicParamStr + noticeParamStr
	signstr, err := signature.SignatureString(urlstr, http.MethodGet)
	if err != nil {
		return "", errors.New("PushNotice signature.SignatureString err")
	}
	signaturestr, err := signature.GetSignature(signstr, notify.AccessSecret)
	if err != nil {
		return "", errors.New("PushNotice signature.GetSignature err")
	}
	finalUrlStr := urlstr + "&Signature=" + signaturestr
	fmt.Println("finalUrlStr---->:",finalUrlStr)
	resp, err := http.Get(finalUrlStr)
	if err != nil {
		return "", errors.New("PushNotice http.Get err")
	}
	resultByte, _ := ioutil.ReadAll(resp.Body)
	return string(resultByte), nil
}
