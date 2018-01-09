package main

import (
	"fmt"
	"github.com/GiterLab/aliyun-acmp-go-sdk/acmp"
	"os"
)

func main() {
	acmp.HTTPDebugEnable = true
	acmp.SetACLClient(ACCESSID, ACCESSKEY)

	respListSummaryPush, err := acmp.ListSummaryPush("appKey", "startTime", "endTime", "pushType").
		SetPage(2).SetPageSize(20).DoActionWithException()
	if err != nil {
		fmt.Println("ListSummaryPush failed", err, respListSummaryPush.Error())
		os.Exit(0)
	}
	fmt.Println("ListSummaryPush successed", respListSummaryPush.String())
}
