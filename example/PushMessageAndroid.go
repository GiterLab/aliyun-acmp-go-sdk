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
	pushMsgRespon, err := acmp.PushMessage2Android(23267207, "target", "targetValue", "title", "body").
		DoActionWithException()
	if err != nil {
		fmt.Println("PushMessage2Android failed", err, pushMsgRespon.Error())
		os.Exit(0)
	}
	fmt.Println("PushMessage2Android successed", pushMsgRespon.String())
}
