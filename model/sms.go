package model

type SmsMessage struct {
	EcName    string `json:"ecName"`
	ApId      string `json:"apId"`
	Mobiles   string `json:"mobiles"`
	Content   string `json:"content"`
	Sign      string `json:"sign"`
	AddSerial string `json:"addSerial"`
	Mac       string `json:"mac"`
}
