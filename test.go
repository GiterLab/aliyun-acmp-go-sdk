package main

import (
	"aliyun-acmp-go-sdk/acmp/push"
	"fmt"
)

func main() {

	//push notify to terminal
	//publicParam's all property should be set
	publicParam := &push.PublicParam{}

	//notifyParam's most property should be set
	notifyParam := &push.NoticeParam{}
	notify := &push.Notify{}
	notify.SetRootUrl("rootUrl")
	notify.SetAccessSecret("accesssecret")
	notify.SetPublicParam(publicParam)
	notify.SetNoticeParam(notifyParam)
	returnNotify, err := notify.DoACMP()
	if err == nil {
		fmt.Println(returnNotify)
	}

	//push message to terminal
	//messageParam's most property should be set
	messageParam := &push.MessageParam{}
	message := &push.Message{}
	message.SetRootUrl("rootUrl")
	message.SetAccessSecret("accesssecret")
	message.SetPublicParam(publicParam)
	message.SetMessageParam(messageParam)
	returnMessage, err := message.DoACMP()
	if err == nil {
		fmt.Println(returnMessage)
	}

}
