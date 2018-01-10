package main

import (
	"fmt"
	"github.com/GiterLab/aliyun-acmp-go-sdk/acmp"
	"os"
)

func main() {
	acmp.HTTPDebugEnable = true
	acmp.SetACLClient(ACCESSID, ACCESSKEY)

	// 推送消息给android设备
	pushMsgRespones, err := acmp.PushMessageToAndroid(23267207, "target", "targetValue", "title", "body").
		DoActionWithException()
	if err != nil {
		fmt.Println("PushMessageToAndroid failed", err, pushMsgRespones.Error())
		os.Exit(0)
	}
	fmt.Println("PushMessageToAndroid successed", pushMsgRespones.String())
}
