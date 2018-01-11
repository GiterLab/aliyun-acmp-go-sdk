package main

import (
	"fmt"
	"github.com/GiterLab/aliyun-acmp-go-sdk/acmp"
	"os"
)

func main() {
	acmp.HTTPDebugEnable = true
	acmp.SetACLClient(ACCESSID, ACCESSKEY)

	respUnbindAlias, err := acmp.UnbindAlias(123213, "deviceID").
		SetUnbindAll(false).SetAliasName("aliasName").DoActionWithException()
	if err != nil {
		fmt.Println("UnbindAlias failed", err, respUnbindAlias.Error())
		os.Exit(0)
	}
	fmt.Println("UnbindAlias successed", respUnbindAlias.String())
}
