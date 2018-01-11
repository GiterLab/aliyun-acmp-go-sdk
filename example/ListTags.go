package main

import (
	"fmt"
	"github.com/GiterLab/aliyun-acmp-go-sdk/acmp"
	"os"
)

func main() {
	acmp.HTTPDebugEnable = true
	acmp.SetACLClient(ACCESSID, ACCESSKEY)

	respListTags, err := acmp.ListTags(123213).DoActionWithException()
	if err != nil {
		fmt.Println("ListTags failed", err, respListTags.Error())
		os.Exit(0)
	}
	fmt.Println("ListTags successed", respListTags.String())
}
