package acmp

import "errors"

//type messageActionType int
//
//var pushMessageToAndroid="PushMessageToAndroid"
//var pushMessageToiOS="PushMessageToiOS"
//
//var MessgeActionType  = [...]string{
//	&pushMessageToAndroid,
//	&pushMessageToiOS,
//}
//
////type MessgeActionType string
//
//const (
//	PushMessageToAndroid messageActionType=iota
//	PushMessageToiOS
//)

const (
	PushMessageToAndroid string = "PushMessageToAndroid"
	PushMessageToiOS            = "PushMessageToiOS"
)

const (
	DEVICE  string = "DEVICE"
	ACCOUNT        = "ACCOUNT"
	ALIAS          = "ALIAS"
	TAG            = "TAG"
	ALL            = "ALL"
)

type MessageParam struct {
	Action      string `json:"action"`
	AppKey      string `json:"app_key"`
	Target      string `json:"target"`
	TargetValue string `json:"target_value"`
	Title       string `json:"title"`
	Body        string `json:"body"`
}

func (m *MessageParam) ToString() (paramstrp string, err error) {
	if m == nil {
		return "", errors.New("MessageParam pointer shouldn't be nil")
	}
	if m.Action == "" || m.AppKey == "" || m.Target == "" || m.TargetValue == "" || m.Title == "" || m.Body == "" {
		return "", errors.New("MessageParam some perpoties shouldn't be nil")
	}
	if m.Action != PushMessageToAndroid || m.Action != PushMessageToiOS {
		return "", errors.New("MessageParam Action should be PushMessageToAndroid or PushMessageToiOS")
	}
	if m.Target != DEVICE || m.Target != ACCOUNT || m.Target != ALIAS || m.Target != TAG || m.Target != ALL {
		return "", errors.New("MessageParam Target should be DEVICE, ACCOUNT,ALIAS,TAG,ALL or PushMessageToiOS")
	}
	var headstr string
	headstr += "Action=" + m.Action + "&AppKey=" + m.AppKey + "&Target=" + m.Target + "&TargetValue=" + m.TargetValue + "&Title=" + m.Title + "&Body=" + m.Body
	return headstr, nil
}
