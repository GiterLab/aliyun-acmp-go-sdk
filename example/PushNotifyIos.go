package main

import (
	"fmt"
	"github.com/GiterLab/aliyun-acmp-go-sdk/acmp"
	"os"
)

func main() {
	acmp.HTTPDebugEnable = true
	acmp.SetACLClient(ACCESSID, ACCESSKEY)

	// 推送通知给iOS设备
	extParameters := make(map[string]interface{}, 0)
	pushNoticeRespon, err := acmp.PushNoticeToiOS(23267207, "target", "targetValue", "apnsEnv", "body").
		SetPushTitle("title").SetPushExtParameters(extParameters).DoActionWithException()
	if err != nil {
		fmt.Println("PushNoticeToiOS failed", err, pushNoticeRespon.Error())
		os.Exit(0)
	}
	fmt.Println("PushNoticeToiOS successed", pushNoticeRespon.String())
}
