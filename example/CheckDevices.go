package main

import (
	"fmt"
	"github.com/GiterLab/aliyun-acmp-go-sdk/acmp"
	"os"
)

func main() {
	acmp.HTTPDebugEnable = true
	acmp.SetACLClient(ACCESSID, ACCESSKEY)

	// 检查设备
	checkDevicesRespones, err := acmp.CheckDevices(24640440, "deviceIds").DoActionWithException()
	if err != nil {
		fmt.Println("CheckDevices failed", err, checkDevicesRespones.Error())
		os.Exit(0)
	}
	fmt.Println("CheckDevices successed", checkDevicesRespones.String())
}
