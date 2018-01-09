package main

import (
	"fmt"
	"github.com/GiterLab/aliyun-acmp-go-sdk/acmp"
	"os"
)

func main() {
	acmp.HTTPDebugEnable = true
	acmp.SetACLClient(ACCESSID, ACCESSKEY)

	respCancelPush, err := acmp.CancelPush("messageId").DoActionWithException()
	if err != nil {
		fmt.Println("CancelPush failed", err, respCancelPush.Error())
		os.Exit(0)
	}
	fmt.Println("CancelPush successed", respCancelPush.String())
}
