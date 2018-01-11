package main

import (
	"fmt"
	"github.com/GiterLab/aliyun-acmp-go-sdk/acmp"
	"os"
)

func main() {
	acmp.HTTPDebugEnable = true
	acmp.SetACLClient(ACCESSID, ACCESSKEY)

	respQueryTags, err := acmp.QueryTags(123213, "clientKey", "keyType").DoActionWithException()
	if err != nil {
		fmt.Println("QueryTags failed", err, respQueryTags.Error())
		os.Exit(0)
	}
	fmt.Println("QueryTags successed", respQueryTags.String())
}
