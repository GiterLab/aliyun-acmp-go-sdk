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
	pushMsgRespon, err := acmp.PushMessage2Ios("target", "targetValue", "body").
		SetPushTitle("title").SetPushExtParameters(extParameters).DoActionWithException()
	if err != nil {
		fmt.Println("PushMessage2Ios failed", err, pushMsgRespon.Error())
		os.Exit(0)
	}
	fmt.Println("PushMessage2Ios successed", pushMsgRespon.String())
}
