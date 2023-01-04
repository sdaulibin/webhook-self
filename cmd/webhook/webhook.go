package main

import (
	"flag"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	h            bool
	DefaultRobot string
	//go build -defaultMobiles="13705328368,15063049112,15621757755,13210002670,15553208391"
	DefaultMobiles string
)

func init() {
	flag.BoolVar(&h, "h", false, "help")
	flag.StringVar(&DefaultRobot, "defaultRobot", "", "global dingtalk robot webhook, you can overwrite by alert rule with annotations dingtalkRobot")
	flag.StringVar(&DefaultMobiles, "defaultMobiles", "15553208888", "短信发送名单")
	log.SetPrefix("【告警系统】")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
}

func main() {
	flag.Parse()
	if h {
		flag.Usage()
		return
	}
	router := gin.Default()
	router.POST("/webhook/dingtalk", Dingtalk)
	router.POST("/webhook/sms", Sms)
	router.Run(":8090")
}
