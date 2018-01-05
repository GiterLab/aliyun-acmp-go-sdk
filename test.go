package main

import (
	"aliyun-acmp-go-sdk/acmp/acmp"
	"fmt"
)

func main() {

	//acmp notify to terminal
	//publicParam's all property should be set
	publicParam := &acmp.PublicParam{}

	//notifyParam's most property should be set
	notifyParam := &acmp.NoticeParam{}
	notify := &acmp.Notify{}
	notify.SetRootUrl("rootUrl")
	notify.SetAccessSecret("accesssecret")
	notify.SetPublicParam(publicParam)
	notify.SetNoticeParam(notifyParam)
	returnNotify, err := notify.DoACMP()
	if err == nil {
		fmt.Println(returnNotify)
	}

	//acmp message to terminal
	//messageParam's most property should be set
	messageParam := &acmp.MessageParam{}
	message := &acmp.Message{}
	message.SetRootUrl("rootUrl")
	message.SetAccessSecret("accesssecret")
	message.SetPublicParam(publicParam)
	message.SetMessageParam(messageParam)
	returnMessage, err := message.DoACMP()
	if err == nil {
		fmt.Println(returnMessage)
	}

}
