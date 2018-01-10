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

	// 推送通知给android设备
	extParameters := make(map[string]interface{}, 0)
	pushNoticeRespones, err := acmp.PushNoticeToAndroid(24639402, "ACCOUNT", "59f02c23-3d61-4ac7-a07e-a9bb2a7970bf", "this is title", "test").
		SetPushExtParameters(extParameters).DoActionWithException()
	if err != nil {
		fmt.Println("PushNoticeToAndroid failed", err, pushNoticeRespones.Error())
		os.Exit(0)
	}
	fmt.Println("PushNoticeToAndroid successed", pushNoticeRespones.String())
}
