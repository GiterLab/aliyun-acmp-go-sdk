package main

import (
	"fmt"
	"github.com/GiterLab/aliyun-acmp-go-sdk/acmp"
	"os"
)

func main() {
	acmp.HTTPDebugEnable = true
	acmp.SetACLClient(ACCESSID, ACCESSKEY)

	extParameters := make(map[string]interface{}, 0)
	pushNotifyRespon, err := acmp.PushNotify2Ios("target", "targetValue", "apnsEnv", "body").
		SetPushTitle("title").SetPushExtParameters(extParameters).DoActionWithException()
	if err != nil {
		fmt.Println("PushNotify2Ios failed", err, pushNotifyRespon.Error())
		os.Exit(0)
	}
	fmt.Println("PushNotify2Ios successed", pushNotifyRespon.String())
}
