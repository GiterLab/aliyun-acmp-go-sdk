package main

import (
	"fmt"
	"github.com/GiterLab/aliyun-acmp-go-sdk/acmp"
	"os"
)

func main() {
	acmp.HTTPDebugEnable = true
	acmp.SetACLClient(ACCESSID, ACCESSKEY)

	respUnbindTag, err := acmp.UnbindTag(123213, "clientKey", "keyType", "tagName").DoActionWithException()
	if err != nil {
		fmt.Println("UnbindTag failed", err, respUnbindTag.Error())
		os.Exit(0)
	}
	fmt.Println("UnbindTag successed", respUnbindTag.String())
}
