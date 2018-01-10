package main

import (
	"fmt"
	"github.com/GiterLab/aliyun-acmp-go-sdk/acmp"
	"os"
)

func main() {
	// pass
	acmp.HTTPDebugEnable = true
	acmp.SetACLClient(ACCESSID, ACCESSKEY)

	queryPushStatByMsgRespones, err := acmp.QueryPushStatByMsg(24629355, "510427").DoActionWithException()
	if err != nil {
		fmt.Println("QueryPushStatByApp failed", err, queryPushStatByMsgRespones.Error())
		os.Exit(0)
	}
	fmt.Println("QueryPushStatByApp successed", queryPushStatByMsgRespones.String())
}
