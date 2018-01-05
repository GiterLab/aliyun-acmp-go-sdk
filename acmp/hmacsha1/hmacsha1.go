package hmacsha1

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
)

func GetHmacStr(paramstr string, keystr string) (hmacstr string, err error) {
	if paramstr == "" || keystr == "" {
		return "", errors.New("GetHmacStr parameter shouldn't be nil")
	}
	hmacsha1 := hmac.New(sha1.New, []byte(keystr+"&"))
	hmacsha1.Write([]byte(paramstr))
	fmt.Println(hmacsha1.Sum(nil))
	return base64.StdEncoding.EncodeToString(hmacsha1.Sum(nil)), nil
}
