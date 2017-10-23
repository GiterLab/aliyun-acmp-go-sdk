package push

import (
	"aliyun-acmp-go-sdk/acmp/bean"
	"aliyun-acmp-go-sdk/acmp/signature"
	"errors"
	"fmt"
	"net/http"
	"io/ioutil"
)

func PushNotice(rootUrl, accessSecret *string, publicParam *bean.PublicParam, noticeParam *bean.NoticeParam) (responeString *string, err error) {
	if publicParam == nil || noticeParam == nil {
		return nil, errors.New("PushNotice param pointer shouldn't be nil")
	}
	publicParamStr, _ := publicParam.ToStringWithoutSignature()
	noticeParamStr, _ := noticeParam.ToString()
	urlstr := *rootUrl + "/?" + *publicParamStr + *noticeParamStr
	var httpurlp *string
	httpurlp = &urlstr
	method := http.MethodGet
	var httpMethodP *string
	httpMethodP = &method
	signstr, err := signature.SignatureString(httpurlp, httpMethodP)
	if err != nil {
		return nil, errors.New("PushNotice signature.SignatureString err")
	}
	signaturestr, err := signature.GetSignature(signstr, accessSecret)
	if err != nil {
		return nil, errors.New("PushNotice signature.GetSignature err")
	}
	finalUrlStr := urlstr + "&Signature=" + *signaturestr
	fmt.Println(finalUrlStr)
	resp, err := http.Get(finalUrlStr)
	if err != nil {
		return nil, errors.New("PushNotice http.Get err")
	}
	resultByte,_:=ioutil.ReadAll(resp.Body)
	temp:=string(resultByte)
	return &temp, nil
}
