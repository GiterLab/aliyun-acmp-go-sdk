package main

import (
	"fmt"
	"github.com/GiterLab/aliyun-acmp-go-sdk/acmp"
	"os"
)

func main() {
	acmp.HTTPDebugEnable = true
	acmp.SetACLClient(ACCESSID, ACCESSKEY)

	// 查询推送列表
	respListSummaryPush, err := acmp.ListSummaryPush(23267207, "startTime", 1515499851, 1515599851).
		SetPage(2).SetPageSize(20).DoActionWithException()
	if err != nil {
		fmt.Println("ListSummaryPush failed", err, respListSummaryPush.Error())
		os.Exit(0)
	}
	fmt.Println("ListSummaryPush successed", respListSummaryPush.String())
}
