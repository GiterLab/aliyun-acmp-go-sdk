package main

import (
	"aliyun-acmp-go-sdk/acmp/signature"
	"fmt"
)

func main() {
	/*
		str:="abcde"
		var strp *string
		strp=&str
		keys:="cdefg"
		var keysp *string
		keysp=&keys
		strr,err:=hmacsha1.GetHmacStr(strp,keysp)
		if err==nil {
			fmt.Print(*strr)
		}*/

	httpurl := "http://cloudpush.aliyuncs.com/?Format=XML&AccessKeyId=testid&Action=GetDeviceInfos&SignatureMethod=HMAC-SHA1&RegionId=cn-hangzhou&Devices=e2ba19de97604f55b165576736477b74%2C92a1da34bdfd4c9692714917ce22d53d&SignatureNonce=c4f5f0de-b3ff-4528-8a89-fa478bda8d80&SignatureVersion=1.0&Version=2016-08-01&AppKey=23267207&Timestamp=2016-03-29T03%3A59%3A24Z"
	var httpurlp *string
	httpurlp = &httpurl
	httpmethod := "GET"
	var httpmethodp *string
	httpmethodp = &httpmethod
	signstr, err := signature.SignatureString(httpurlp, httpmethodp)
	if err == nil {
		fmt.Println(*signstr)
	}
	keys := "testsecret"
	var keysp *string
	keysp = &keys
	signstrr, err := signature.GetSignature(signstr, keysp)
	if err == nil {
		fmt.Println(*signstrr)
	}

}
