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
	pushMsgRespon, err := acmp.PushMessageToiOSRequestos(23267207, "target", "targetValue", "body").
		SetPushTitle("title").SetPushExtParameters(extParameters).DoActionWithException()
	if err != nil {
		fmt.Println("PushMessageToiOSRequestos failed", err, pushMsgRespon.Error())
		os.Exit(0)
	}
	fmt.Println("PushMessageToiOSRequestos successed", pushMsgRespon.String())
}
