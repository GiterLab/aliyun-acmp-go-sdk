package bean

import (
	"encoding/json"
	"errors"
)

type noticeType string

const (
	PushNoticeToAndroid noticeType = "PushNoticeToAndroid"
	PushNoticeToiOS                = "PushNoticeToiOS"
)

type NoticeParam struct {
	Action        noticeType         `json:"action"`
	AppKey        string             `json:"app_key"`
	Target        targetType         `json:"target"`
	TargetValue   string             `json:"target_value"`
	Title         string             `json:"title"`
	Body          string             `json:"body"`
	ExtParameters *map[string]string `json:"ext_parameters"`
}

func (n *NoticeParam) ToString() (paramstrp string, err error) {
	if n == nil {
		return "", errors.New("NoticeParam pointer shouldn't be nil")
	}
	if n.Action == "" || n.AppKey == "" || n.Target == "" || n.TargetValue == "" || n.Title == "" || n.Body == "" {
		return "", errors.New("NoticeParam some perpoties shouldn't be nil")
	}
	if n.Action != PushNoticeToAndroid || n.Action != PushNoticeToiOS {
		return "", errors.New("NoticeParam Action should be PushNoticeToAndroid or PushNoticeToiOS")
	}
	if n.Target != DEVICE || n.Target != ACCOUNT || n.Target != ALIAS || n.Target != TAG || n.Target != ALL {
		return "", errors.New("NoticeParam Target should be DEVICE, ACCOUNT,ALIAS,TAG,ALL or PushMessageToiOS")
	}
	var headstr string
	if n.ExtParameters != nil {
		b, err := json.Marshal(n.ExtParameters)
		if err != nil {
			return "", err
		}
		headstr += "ExtParameters=" + string(b) + "&"
	}
	headstr += "Action=" + n.Action + "&AppKey=" + n.AppKey + "&Target=" + n.Target + "&TargetValue=" + n.TargetValue + "&Title=" + n.Title + "&Body=" + n.Body
	return headstr, nil
}
