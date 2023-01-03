package main

import (
	"flag"
	"log"
	"net/http"

	model "binginx.com/webhook/model"
	"binginx.com/webhook/notifier"
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
	router.POST("/webhook/dingding", func(c *gin.Context) {
		var notification model.Notification
		err := c.BindJSON(&notification)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		log.Println("notification:", notification)
		err = notifier.Send(notification, defaultRobot)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"message": "send to dingtalk successful!"})

	})
	router.POST("/webhook/sms", func(c *gin.Context) {
		var notification model.Notification
		err := c.BindJSON(&notification)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		log.Println("notification:", notification)
		err = notifier.SendSms(notification, defaultMobiles)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"message": "send to sms successful!"})
	})
	router.Run(":8090")
}
