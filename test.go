package main

import (
	"fmt"
	"aliyun-acmp-push-go-sdk/models/signature"
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

	httpurl:="http://cloudpush.aliyuncs.com/?Format=XML&AccessKeyId=testid&Action=GetDeviceInfos&SignatureMethod=HMAC-SHA1&RegionId=cn-hangzhou&Devices=e2ba19de97604f55b165576736477b74%2C92a1da34bdfd4c9692714917ce22d53d&SignatureNonce=c4f5f0de-b3ff-4528-8a89-fa478bda8d80&SignatureVersion=1.0&Version=2016-08-01&AppKey=23267207&Timestamp=2016-03-29T03%3A59%3A24Z"
	var httpurlp *string
	httpurlp=&httpurl
	httpmethod:="GET"
	var httpmethodp *string
	httpmethodp=&httpmethod
	signstr,err:=signature.SignatureString(httpurlp,httpmethodp)
	if err==nil {
		fmt.Println(*signstr)
	}

	str:="GET&%2F&AccessKeyId%3Dtestid&Action%3DGetDeviceInfos&AppKey%3D23267207&Devices%3De2ba19de97604f55b165576736477b74%252C92a1da34bdfd4c9692714917ce22d53d&Format%3DXML&RegionId%3Dcn-hangzhou&SignatureMethod%3DHMAC-SHA1&SignatureNonce%3Dc4f5f0de-b3ff-4528-8a89-fa478bda8d80&SignatureVersion%3D1.0&Timestamp%3D2016-03-29T03%253A59%253A24Z&Version%3D2016-08-01"
	fmt.Println(str)
	var strp *string
	strp=&str
	keys:="testsecret"
	var keysp *string
	keysp=&keys
	signstrr,err:=signature.GetSignature(signstr,keysp)
	strpr,err:=signature.GetSignature(strp,keysp)
	if err==nil {
		fmt.Println(*signstrr)
		fmt.Println(*strpr)
	}


}
