package bean

import "errors"

//type messageActionType int
//
//var pushMessageToAndroid="PushMessageToAndroid"
//var pushMessageToiOS="PushMessageToiOS"
//
//var MessgeActionType  = [...]*string{
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

type messageType string

const (
	PushMessageToAndroid messageType = "PushMessageToAndroid"
	PushMessageToiOS                 = "PushMessageToiOS"
)

type targetType string

const (
	DEVICE  targetType = "DEVICE"
	ACCOUNT            = "ACCOUNT"
	ALIAS              = "ALIAS"
	TAG                = "TAG"
	ALL                = "ALL"
)

type MessageParam struct {
	Action      *messageType `json:"action"`
	AppKey      *string      `json:"app_key"`
	Target      *targetType  `json:"target"`
	TargetValue *string      `json:"target_value"`
	Title       *string      `json:"title"`
	Body        *string      `json:"body"`
}

func (this *MessageParam) ToString() (paramstrp *string, err error) {
	if this == nil {
		return nil, errors.New("MessageParam pointer shouldn't be nil")
	}
	if this.Action == nil || this.AppKey == nil || this.Target == nil || this.TargetValue == nil || this.Title == nil || this.Body == nil {
		return nil, errors.New("MessageParam some perpoties shouldn't be nil")
	}
	if *this.Action != PushMessageToAndroid || *this.Action != PushMessageToiOS {
		return nil, errors.New("MessageParam Action should be PushMessageToAndroid or PushMessageToiOS")
	}
	if *this.Target != DEVICE || *this.Target != ACCOUNT || *this.Target != ALIAS || *this.Target != TAG || *this.Target != ALL {
		return nil, errors.New("MessageParam Target should be DEVICE, ACCOUNT,ALIAS,TAG,ALL or PushMessageToiOS")
	}
	var headstr string
	headstr += "Action=" + *this.Action + "&AppKey=" + *this.AppKey + "&Target=" + *this.Target + "&TargetValue=" + *this.TargetValue + "&Title=" + *this.Title + "&Body=" + *this.Body
	return &headstr, nil
}
