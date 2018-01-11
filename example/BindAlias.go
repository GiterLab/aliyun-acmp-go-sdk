package main

import (
	"fmt"
	"github.com/GiterLab/aliyun-acmp-go-sdk/acmp"
	"os"
)

func main() {
	acmp.HTTPDebugEnable = true
	acmp.SetACLClient(ACCESSID, ACCESSKEY)

	respBindAlias, err := acmp.BindAlias(123213, "DeviceID", "AliasName").DoActionWithException()
	if err != nil {
		fmt.Println("BindAlias failed", err, respBindAlias.Error())
		os.Exit(0)
	}
	fmt.Println("BindAlias successed", respBindAlias.String())
}
