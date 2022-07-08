package transformer

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"

	"binginx.com/webhook/model"
)

const (
	EcName    = "xxxxxx科技股份有限公司。"
	Sign      = "xxxx"
	ApId      = "xxxx"
	SecretKey = "xxxx"
)

func TransformToSms(notification model.Notification, defaultMobiles string) (smsMessage *model.SmsMessage, err error) {
	groupKey := notification.GroupKey
	status := notification.Status

	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("### 通知组%s(当前状态:%s) \n", groupKey, status))

	buffer.WriteString(fmt.Sprintf("#### 告警项:\n"))

	for _, alert := range notification.Alerts {
		annotations := alert.Annotations
		buffer.WriteString(fmt.Sprintf("##### %s\n > %s\n", annotations["summary"], annotations["description"]))
		buffer.WriteString(fmt.Sprintf("\n> 开始时间：%s\n", alert.StartsAt.Format("15:04:05")))
	}

	fmt.Println("组装前:", buffer.String())

	if defaultMobiles == "" {
		defaultMobiles = "x,x,x,x,x"
	}

	smsMessage = &model.SmsMessage{
		EcName:    EcName,
		ApId:      ApId,
		Mobiles:   defaultMobiles,
		Content:   buffer.String(),
		Sign:      Sign,
		AddSerial: "",
	}
	log.Println("defaultMobiles:", defaultMobiles)
	smsMessage.Mac = md5sms(smsMessage.EcName + smsMessage.ApId + SecretKey + smsMessage.Mobiles + smsMessage.Content + smsMessage.Sign + "")
	return
}

// TransformToMarkdown transform alertmanager notification to dingtalk markdow message
func TransformToMarkdown(notification model.Notification) (markdown *model.DingTalkMarkdown, robotURL string, err error) {

	groupKey := notification.GroupKey
	status := notification.Status

	// annotations := notification.CommonAnnotations
	// robotURL = annotations["dingtalkRobot"]
	// robotURL = "https://oapi.dingtalk.com/robot/send?access_token=208a84e6e1d3e854ac55f106362949ff7cd791a7c5cbc92a6f6fc81e915cb764"
	robotURL = "https://oapi.xxxx.com/robot/send?access_token=xxxx3a5a2905da220c21cb047b47445fd9a33a1cb0c1b49cde193588664"
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("###核酸检测 通知组%s(当前状态:%s) \n", groupKey, status))

	buffer.WriteString(fmt.Sprintf("#### 告警项:\n"))

	for _, alert := range notification.Alerts {
		annotations := alert.Annotations
		buffer.WriteString(fmt.Sprintf("##### %s\n > %s\n", annotations["summary"], annotations["description"]))
		buffer.WriteString(fmt.Sprintf("\n> 开始时间：%s\n", alert.StartsAt.Format("15:04:05")))
	}

	log.Println("组装前:", buffer.String())

	markdown = &model.DingTalkMarkdown{
		MsgType: "markdown",
		Markdown: &model.Markdown{
			Title: fmt.Sprintf("通知组：%s(当前状态:%s)", groupKey, status),
			Text:  buffer.String(),
		},
		At: &model.At{
			IsAtAll: false,
		},
	}
	log.Println("组装后-text:", markdown.Markdown.Text)
	return
}

func md5sms(s string) string {
	d := []byte(s)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}
