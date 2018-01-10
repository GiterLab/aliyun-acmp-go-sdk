package main

import (
	"fmt"
	"github.com/GiterLab/aliyun-acmp-go-sdk/acmp"
	"os"
)

func main() {
	// pass
	acmp.HTTPDebugEnable = true
	acmp.SetACLClient(ACCESSID, ACCESSKEY)

	queryDeviceStatRespones, err := acmp.QueryDeviceStat(24629355, 1515540755, 1515550755, "iOS", "TOTAL").DoActionWithException()
	if err != nil {
		fmt.Println("QueryDeviceStat failed", err, queryDeviceStatRespones.Error())
		os.Exit(0)
	}
	fmt.Println("QueryDeviceStat successed", queryDeviceStatRespones.String())
}
