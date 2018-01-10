package main

import (
	"fmt"
	"github.com/GiterLab/aliyun-acmp-go-sdk/acmp"
	"os"
)

func main() {
	acmp.HTTPDebugEnable = true
	acmp.SetACLClient(ACCESSID, ACCESSKEY)

	// 查询去重设备信息
	queryUniqueDeviceStatRespones, err := acmp.QueryUniqueDeviceStat(24640440, 1515499851, 1515599851).DoActionWithException()
	if err != nil {
		fmt.Println("QueryUniqueDeviceStat failed", err, queryUniqueDeviceStatRespones.Error())
		os.Exit(0)
	}
	fmt.Println("QueryUniqueDeviceStat successed", queryUniqueDeviceStatRespones.String())
}
