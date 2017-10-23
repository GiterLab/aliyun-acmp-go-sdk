package push

import (
	"aliyun-acmp-go-sdk/acmp/bean"
	"aliyun-acmp-go-sdk/acmp/signature"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func PushNotice(rootUrl, accessSecret string, publicParam *bean.PublicParam, noticeParam *bean.NoticeParam) (responeString string, err error) {
	if publicParam == nil || noticeParam == nil {
		return "", errors.New("PushNotice param pointer shouldn't be nil")
	}
	publicParamStr, _ := publicParam.ToStringWithoutSignature()
	noticeParamStr, _ := noticeParam.ToString()
	urlstr := rootUrl + "/?" + publicParamStr + noticeParamStr
	signstr, err := signature.SignatureString(urlstr, http.MethodGet)
	if err != nil {
		return "", errors.New("PushNotice signature.SignatureString err")
	}
	signaturestr, err := signature.GetSignature(signstr, accessSecret)
	if err != nil {
		return "", errors.New("PushNotice signature.GetSignature err")
	}
	finalUrlStr := urlstr + "&Signature=" + signaturestr
	fmt.Println(finalUrlStr)
	resp, err := http.Get(finalUrlStr)
	if err != nil {
		return "", errors.New("PushNotice http.Get err")
	}
	resultByte, _ := ioutil.ReadAll(resp.Body)
	return string(resultByte), nil
}
