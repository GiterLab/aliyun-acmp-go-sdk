package push

import (
	"aliyun-acmp-go-sdk/acmp/bean"
	"aliyun-acmp-go-sdk/acmp/signature"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func PushMessage(rootUrl, accessSecret string, publicParam *bean.PublicParam, noticeParam *bean.MessageParam) (responeString string, err error) {
	if publicParam == nil || noticeParam == nil {
		return "", errors.New("PushMessage param pointer shouldn't be nil")
	}
	publicParamStr, _ := publicParam.ToStringWithoutSignature()
	noticeParamStr, _ := noticeParam.ToString()
	urlstr := rootUrl + "/?" + publicParamStr + noticeParamStr
	signstr, err := signature.SignatureString(urlstr, http.MethodGet)
	if err != nil {
		return "", errors.New("PushMessage signature.SignatureString err")
	}
	signaturestr, err := signature.GetSignature(signstr, accessSecret)
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
