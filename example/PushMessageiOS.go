package main

import (
	"fmt"
	"github.com/GiterLab/aliyun-acmp-go-sdk/acmp"
	"os"
)

func main() {
	acmp.HTTPDebugEnable = true
	acmp.SetACLClient(ACCESSID, ACCESSKEY)

	// 推送消息给ios设备
	extParameters := make(map[string]interface{}, 0)
	pushMsgRespones, err := acmp.PushMessageToiOS(23267207, "target", "targetValue", "body").
		SetPushTitle("title").SetPushExtParameters(extParameters).DoActionWithException()
	if err != nil {
		fmt.Println("PushMessageToiOSRequest failed", err, pushMsgRespones.Error())
		os.Exit(0)
	}
	fmt.Println("PushMessageToiOSRequest successed", pushMsgRespones.String())
}
