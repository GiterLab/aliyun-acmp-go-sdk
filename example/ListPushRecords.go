package main

import (
	"fmt"
	"github.com/GiterLab/aliyun-acmp-go-sdk/acmp"
	"os"
)

func main() {
	//pass
	acmp.HTTPDebugEnable = true
	acmp.SetACLClient(ACCESSID, ACCESSKEY)

	// 查询推送列表
	respListSummaryPush, err := acmp.ListPushRecords(24629355, "NOTICE", 1515499851, 1515599851).
		SetPage(2).SetPageSize(20).DoActionWithException()
	if err != nil {
		fmt.Println("ListSummaryPush failed", err, respListSummaryPush.Error())
		os.Exit(0)
	}
	fmt.Println("ListSummaryPush successed", respListSummaryPush.String())
}
