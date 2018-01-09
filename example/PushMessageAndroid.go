package main

import (
	"fmt"
	"github.com/GiterLab/aliyun-acmp-go-sdk/acmp"
	"os"
)

func main() {
	acmp.HTTPDebugEnable = true
	acmp.SetACLClient(ACCESSID, ACCESSKEY)

	pushMsgRespon, err := acmp.PushMessage2Android("target", "targetValue", "title", "body").
		DoActionWithException()
	if err != nil {
		fmt.Println("PushMessage2Android failed", err, pushMsgRespon.Error())
		os.Exit(0)
	}
	fmt.Println("PushMessage2Android successed", pushMsgRespon.String())
}
