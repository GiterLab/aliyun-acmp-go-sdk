package hmacsha1

import (
	"crypto/hmac"
	"crypto/sha1"
	"errors"
	"encoding/base64"
	"fmt"
)

func GetHmacStr(paramstr *string, keystr *string) (hmacstr *string, err error) {
	if paramstr == nil || keystr == nil {
		return nil, errors.New("GetHmacStr parameter pointer shouldn't be nil")
	}
	bytestr:=[]byte(*keystr+"&")
	fmt.Println(bytestr)
	hmacsha1 := hmac.New(sha1.New, []byte(*keystr+"&"))
	hmacsha1.Write([]byte(*paramstr))
	fmt.Println(hmacsha1.Sum(nil))
	temp := base64.StdEncoding.EncodeToString(hmacsha1.Sum(nil))
	return &temp, nil
}
