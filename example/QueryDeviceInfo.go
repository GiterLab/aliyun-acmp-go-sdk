package main

import (
	"fmt"
	"github.com/GiterLab/aliyun-acmp-go-sdk/acmp"
	"os"
)

func main() {
	acmp.HTTPDebugEnable = true
	acmp.SetACLClient(ACCESSID, ACCESSKEY)

	// 查询设备信息
	queryDeviceInfoRespones, err := acmp.QueryDeviceInfo(24640440, "deviceId").DoActionWithException()
	if err != nil {
		fmt.Println("QueryDeviceInfo failed", err, queryDeviceInfoRespones.Error())
		os.Exit(0)
	}
	fmt.Println("QueryDeviceInfo successed", queryDeviceInfoRespones.String())
}
