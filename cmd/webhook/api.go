package main

import (
	"log"
	"net/http"

	"binginx.com/webhook/model"
	"binginx.com/webhook/notifier"
	"github.com/gin-gonic/gin"
)

func Dingtalk(c *gin.Context) {
	var notification model.Notification
	err := c.BindJSON(&notification)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("notification:", notification)
	err = notifier.Send(notification, DefaultRobot)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "send to dingtalk successful!"})

}

func Sms(c *gin.Context) {
	var notification model.Notification
	err := c.BindJSON(&notification)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("notification:", notification)
	err = notifier.SendSms(notification, DefaultMobiles)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "send to sms successful!"})

}
