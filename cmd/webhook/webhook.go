package main

import (
	"flag"
	"log"

	"binginx.com/webhook/api"
	"github.com/gin-gonic/gin"
)

var (
	h            bool
	defaultRobot string
	//go build -defaultMobiles="13705328368,15063049112,15621757755,13210002670,15553208391"
	defaultMobiles string
)

func init() {
	flag.BoolVar(&h, "h", false, "help")
	flag.StringVar(&defaultRobot, "defaultRobot", "", "global dingtalk robot webhook, you can overwrite by alert rule with annotations dingtalkRobot")
	flag.StringVar(&defaultMobiles, "defaultMobiles", "", "短信发送名单")
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

	router.POST("/webhook/dingding", api.Dingding)
	router.POST("/webhook/sms", api.Sms)
	router.Run(":8090")
}
