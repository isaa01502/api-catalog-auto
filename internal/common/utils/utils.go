package utils

import (
	"github.com/gin-gonic/gin"
)

func GetUserIP(c *gin.Context) string {
	userIP := c.GetHeader("x-user-ip")
	if len(userIP) > 0 {
		return userIP
	}

	userIP = c.GetHeader("ip")
	if len(userIP) > 0 {
		return userIP
	}

	userIP = c.GetHeader("x-real-ip")
	return userIP
}

// GetDeviceID Возвращает идентификатор девайса
func GetDeviceID(c *gin.Context) string {
	deviceId := c.GetHeader("deviceid")
	if len(deviceId) > 0 {
		return deviceId
	}

	deviceId = c.GetHeader("DeviceID")
	return deviceId
}
