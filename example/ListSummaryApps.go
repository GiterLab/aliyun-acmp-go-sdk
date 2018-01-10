package main

import (
	"fmt"
	"os"

	"github.com/GiterLab/aliyun-acmp-go-sdk/acmp"
)

// modify it to yours
const (
	ACCESSID  = "your_accessid"  // "your_accessid"
	ACCESSKEY = "your_accesskey" // "your_accesskey"
)

func main() {
	// pass
	acmp.HTTPDebugEnable = true
	acmp.SetACLClient(ACCESSID, ACCESSKEY)

	// 获取APP概览列表
	respListSummaryApps, err := acmp.ListSummaryApps().DoActionWithException()
	if err != nil {
		fmt.Println("ListSummaryApps failed", err, respListSummaryApps.Error())
		os.Exit(0)
	}
	fmt.Println("ListSummaryApps succeed", respListSummaryApps.String())
}
