package push

import (
	"aliyun-acmp-go-sdk/acmp/bean"
	"aliyun-acmp-go-sdk/acmp/signature"
	"errors"
	"fmt"
	"net/http"
	"io/ioutil"
)

func PushMessage(rootUrl, accessSecret *string, publicParam *bean.PublicParam, noticeParam *bean.MessageParam) (responeString *string, err error) {
	if publicParam == nil || noticeParam == nil {
		return nil, errors.New("PushMessage param pointer shouldn't be nil")
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
		return nil, errors.New("PushMessage signature.SignatureString err")
	}
	signaturestr, err := signature.GetSignature(signstr, accessSecret)
	if err != nil {
		return nil, errors.New("PushMessage signature.GetSignature err")
	}
	finalUrlStr := urlstr + "&Signature=" + *signaturestr
	fmt.Println(finalUrlStr)
	resp, err := http.Get(finalUrlStr)
	if err != nil {
		return nil, errors.New("PushMessage http.Get err")
	}
	resultByte,_:=ioutil.ReadAll(resp.Body)
	temp:=string(resultByte)
	return &temp, nil
}
