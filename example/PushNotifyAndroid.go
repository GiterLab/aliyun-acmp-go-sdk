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
	pushNotifyRespon, err := acmp.PushNotify2Android("target", "targetValue", "title", "body").
		SetPushExtParameters(extParameters).DoActionWithException()
	if err != nil {
		fmt.Println("PushNotify2Android failed", err, pushNotifyRespon.Error())
		os.Exit(0)
	}
	fmt.Println("PushNotify2Android successed", pushNotifyRespon.String())
}
