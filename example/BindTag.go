package main

import (
	"fmt"
	"github.com/GiterLab/aliyun-acmp-go-sdk/acmp"
	"os"
)

func main() {
	acmp.HTTPDebugEnable = true
	acmp.SetACLClient(ACCESSID, ACCESSKEY)

	respBindTag, err := acmp.BindTag(123213, "clientKey", "keyType", "tagName").DoActionWithException()
	if err != nil {
		fmt.Println("BindTag failed", err, respBindTag.Error())
		os.Exit(0)
	}
	fmt.Println("BindTag successed", respBindTag.String())
}
