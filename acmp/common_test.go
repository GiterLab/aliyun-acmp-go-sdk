package acmp

import (
	"testing"
)

func Test_stringToSign(t *testing.T) {
	c := new(Client)
	c.EndPoint = "http://cloudpush.aliyuncs.com/"
	c.AccessID = "testid"
	c.AccessKey = "testsecret"

	req := newRequset()
	// 1. 系统参数
	req.Put("SignatureMethod", "HMAC-SHA1")
	req.Put("SignatureNonce", "c4f5f0de-b3ff-4528-8a89-fa478bda8d80")
	req.Put("AccessKeyId", c.AccessID)
	req.Put("SignatureVersion", "1.0")
	req.Put("Timestamp", "2016-03-29T03:59:24Z")
	req.Put("Format", "XML")
	// 2. 业务API参数
	req.Put("Action", "GetDeviceInfos")
	req.Put("Version", "2016-08-01")
	req.Put("RegionId", "cn-hangzhou")
	req.Put("AppKey", "23267207")
	req.Put("Devices", "e2ba19de97604f55b165576736477b74,92a1da34bdfd4c9692714917ce22d53d")
	stringToSign := req.CalcStringToSign("GET")
	stringToSignResult := `GET&%2F&AccessKeyId%3Dtestid%26Action%3DGetDeviceInfos%26AppKey%3D23267207%26Devices%3De2ba19de97604f55b165576736477b74%252C92a1da34bdfd4c9692714917ce22d53d%26Format%3DXML%26RegionId%3Dcn-hangzhou%26SignatureMethod%3DHMAC-SHA1%26SignatureNonce%3Dc4f5f0de-b3ff-4528-8a89-fa478bda8d80%26SignatureVersion%3D1.0%26Timestamp%3D2016-03-29T03%253A59%253A24Z%26Version%3D2016-08-01`
	if stringToSign != stringToSignResult {
		t.Error("calcStringToSign failed")
	}
}
