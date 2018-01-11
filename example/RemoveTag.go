package main

import (
	"fmt"
	"github.com/GiterLab/aliyun-acmp-go-sdk/acmp"
	"os"
)

func main() {
	acmp.HTTPDebugEnable = true
	acmp.SetACLClient(ACCESSID, ACCESSKEY)

	respRemoveTag, err := acmp.RemoveTag(123213, "tagName").DoActionWithException()
	if err != nil {
		fmt.Println("RemoveTag failed", err, respRemoveTag.Error())
		os.Exit(0)
	}
	fmt.Println("RemoveTag successed", respRemoveTag.String())
}
