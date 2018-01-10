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

	queryPushStatByAppRespones, err := acmp.QueryPushStatByApp(24629355, 1515499851, 1515599851, "DAY").DoActionWithException()
	if err != nil {
		fmt.Println("QueryPushStatByApp failed", err, queryPushStatByAppRespones.Error())
		os.Exit(0)
	}
	fmt.Println("QueryPushStatByApp successed", queryPushStatByAppRespones.String())
}
