package main

import (
	"fmt"
	"github.com/GiterLab/aliyun-acmp-go-sdk/acmp"
	"os"
)

func main() {
	acmp.HTTPDebugEnable = true
	acmp.SetACLClient(ACCESSID, ACCESSKEY)

	// 推送通知给android设备
	extParameters := make(map[string]interface{}, 0)
	pushNoticeRespones, err := acmp.PushNoticeToAndroid(23267207, "target", "targetValue", "title", "body").
		SetPushExtParameters(extParameters).DoActionWithException()
	if err != nil {
		fmt.Println("PushNoticeToAndroid failed", err, pushNoticeRespones.Error())
		os.Exit(0)
	}
	fmt.Println("PushNoticeToAndroid successed", pushNoticeRespones.String())
}
