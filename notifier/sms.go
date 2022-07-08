package notifier

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"binginx.com/webhook/model"
	"binginx.com/webhook/transformer"
)

func SendSms(notification model.Notification, defaultMobiles string) (err error) {
	smsMessage, err := transformer.TransformToSms(notification, defaultMobiles)

	if err != nil {
		return
	}

	data, err := json.Marshal(smsMessage)
	if err != nil {
		return
	}
	log.Println("组装后的信息:", string(data))

	sEnc := base64.StdEncoding.EncodeToString([]byte(data))

	log.Println("组装后的信息-base64:", sEnc)

	req, err := http.NewRequest(
		"POST",
		"http://112.35.1.155:1992/sms/norsubmit",
		bytes.NewBufferString(sEnc))

	if err != nil {
		log.Println("sms robot url not found ignore:")
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return
	}

	defer resp.Body.Close()
	out, err := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(out))
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	return
}

// Send send markdown message to dingtalk
func Send(notification model.Notification, defaultRobot string) (err error) {

	markdown, robotURL, err := transformer.TransformToMarkdown(notification)

	if err != nil {
		return
	}

	data, err := json.Marshal(markdown)
	if err != nil {
		return
	}
	log.Println("组装后的信息:", string(data))
	var dingTalkRobotURL string

	if robotURL != "" {
		dingTalkRobotURL = robotURL
	} else {
		dingTalkRobotURL = defaultRobot
	}

	if len(dingTalkRobotURL) == 0 {
		return nil
	}
	req, err := http.NewRequest(
		"POST",
		dingTalkRobotURL,
		bytes.NewBuffer(data))

	log.Print("url:", dingTalkRobotURL)
	if err != nil {
		log.Println("dingtalk robot url not found ignore:")
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return
	}

	defer resp.Body.Close()
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)

	return
}
