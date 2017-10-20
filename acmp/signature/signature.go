package signature

import (
	"errors"
	"net/url"
	"fmt"
	"sort"
	"aliyun-acmp-push-go-sdk/models/hmacsha1"
	"strings"
)

func SignatureString(httprequrl *string,httpmethord *string) (signstr *string,err error) {
	if httprequrl==nil {
		return nil,errors.New("SignatureString httprequrl pointer shouldn't be nil")
	}
	u,err:=url.Parse(*httprequrl)
	if err!=nil {
		return nil,errors.New(fmt.Sprint("SignatureString httprequrl parse error %s",err))
	}
	uParam,err:=url.ParseQuery(u.RawQuery)
	if err!=nil {
		return nil,errors.New(fmt.Sprint("SignatureString httprequrl ParseQuery error %s",err))
	}
	i:=0
	strslice := make([]string, len(uParam))
	for k,v:=range uParam{
		data := url.Values{}
		data.Add(k, v[0])
		strslice[i] = data.Encode()
		strslice[i] = aliyunEncodeOver(strslice[i])
		fmt.Println(k+":---->"+v[0])
		i++
	}
	sort.Strings(strslice)
	temp:=*httpmethord+"&" + percentEncode("/") + "&" + percentEncode(strings.Join(strslice, "&"))
	return &temp,nil
}

func GetSignature(urlencodestr *string,accesssecret *string) (signstr *string,err error) {
	signstr,err= hmacsha1.GetHmacStr(urlencodestr,accesssecret)
	if err!=nil {
		return nil,errors.New(fmt.Sprint("SignatureString GetHmacStr error %s",err))
	}
	return signstr,nil
}

// 一般支持 URL 编码的库（比如 Java 中的 java.net.URLEncoder）都是按照“application/x-www-form-urlencoded”的MIME类型的规则进行编码的。
// 实现时可以直接使用这类方式进行编码，
// 把编码后的字符串中加号（+）替换成%20、星号（*）替换成%2A、%7E 替换回波浪号（~）, 即可得到所需要的编码字符串
func percentEncode(s string) string {
	s = url.QueryEscape(s)
	s = strings.Replace(s, "+", "%20", -1)
	s = strings.Replace(s, "*", "%2A", -1)
	s = strings.Replace(s, "%7E", "~", -1)
	s = strings.Replace(s, "%26", "&", -1)
	return s
}

// 把编码后的字符串中加号（+）替换成%20、星号（*）替换成%2A、%7E 替换回波浪号（~）, 即可得到所需要的编码字符串
func aliyunEncodeOver(s string) string {
	s = strings.Replace(s, "+", "%20", -1)
	s = strings.Replace(s, "*", "%2A", -1)
	s = strings.Replace(s, "%7E", "~", -1)
	return s
}
