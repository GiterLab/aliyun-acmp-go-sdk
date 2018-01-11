package main

import (
	"fmt"
	"github.com/GiterLab/aliyun-acmp-go-sdk/acmp"
	"os"
)

func main() {
	acmp.HTTPDebugEnable = true
	acmp.SetACLClient(ACCESSID, ACCESSKEY)

	respQueryAliases, err := acmp.QueryAliases(123213, "deviceID").DoActionWithException()
	if err != nil {
		fmt.Println("QueryAliases failed", err, respQueryAliases.Error())
		os.Exit(0)
	}
	fmt.Println("QueryAliases successed", respQueryAliases.String())
}
