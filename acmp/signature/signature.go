package signature

import (
	"aliyun-acmp-go-sdk/acmp/hmacsha1"
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

func SignatureString(httprequrl string, httpmethord string) (signstr string, err error) {
	if httprequrl == "" {
		return "", errors.New("SignatureString httprequrl pointer shouldn't be \"\"")
	}
	u, err := url.Parse(httprequrl)
	if err != nil {
		return "", errors.New(fmt.Sprint("SignatureString httprequrl parse error %s", err))
	}
	uParam, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return "", errors.New(fmt.Sprint("SignatureString httprequrl ParseQuery error %s", err))
	}
	i := 0
	strslice := make([]string, len(uParam))
	for k, v := range uParam {
		data := url.Values{}
		data.Add(k, v[0])
		strslice[i] = data.Encode()
		strslice[i] = aliyunEncodeOver(strslice[i])
		fmt.Println(k + ":---->" + v[0])
		i++
	}
	sort.Strings(strslice)
	return httpmethord + "&" + percentEncode("/") + "&" + percentEncode(strings.Join(strslice, "&")), nil
}

func GetSignature(urlencodestr string, accesssecret string) (signstr string, err error) {
	signstr, err = hmacsha1.GetHmacStr(urlencodestr, accesssecret)
	if err != nil {
		return "", errors.New(fmt.Sprint("SignatureString GetHmacStr error %s", err))
	}
	return signstr, nil
}

func percentEncode(s string) string {
	s = url.QueryEscape(s)
	s = strings.Replace(s, "+", "%20", -1)
	s = strings.Replace(s, "*", "%2A", -1)
	s = strings.Replace(s, "%7E", "~", -1)
	return s
}

func aliyunEncodeOver(s string) string {
	s = strings.Replace(s, "+", "%20", -1)
	s = strings.Replace(s, "*", "%2A", -1)
	s = strings.Replace(s, "%7E", "~", -1)
	return s
}
