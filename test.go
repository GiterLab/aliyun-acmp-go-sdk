package main

import (
	"aliyun-acmp-go-sdk/acmp/bean"
	"aliyun-acmp-go-sdk/acmp/push"
	"fmt"
)

func main() {

	//push notify to terminal
	//publicParam's all property should be set
	publicParam := &bean.PublicParam{}

	//notifyParam's most property should be set
	notifyParam := &bean.NoticeParam{}
	notify := &push.Notify{}
	notify.SetRootUrl("rootUrl")
	notify.SetAccessSecret("accesssecret")
	notify.SetPublicParam(publicParam)
	notify.SetNoticeParam(notifyParam)
	returnNotify, err := notify.DoPush(notify)
	if err == nil {
		fmt.Println(returnNotify)
	}

	//push message to terminal
	//messageParam's most property should be set
	messageParam := &bean.MessageParam{}
	message := &push.Message{}
	message.SetRootUrl("rootUrl")
	message.SetAccessSecret("accesssecret")
	message.SetPublicParam(publicParam)
	message.SetMessageParam(messageParam)
	returnMessage, err := message.DoPush(message)
	if err == nil {
		fmt.Println(returnMessage)
	}

}
