package utils

import (
	"git.homebank.kz/homebank-halykid/api-halykid/internal/model"
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

func Map(user *model.UserFullInfo) model.UserToFront {

	return model.UserToFront{
		FirstName:  user.Name,
		LastName:   user.Surname,
		Phone:      user.TrustedPhone,
		BirthDay:   user.BirthDate,
		Gender:     user.Gender,
		OpenWayId:  user.OpenWayId,
		Middlename: user.Middlename,
		UserId:     user.UserId,
	}
}
