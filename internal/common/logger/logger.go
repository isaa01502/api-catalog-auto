package logger

import (
	"api-catalog-auto/internal/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	"log"
	"os"
)

// PanicErrorCode общий код ошибки при паники
const (
	PanicErrorCode = -10201
)

var ErrorLogger *log.Logger
var WarningLogger *log.Logger
var InfoLogger *log.Logger

type AppLogger interface {
	HttpPanicHandler(ctx *gin.Context, recovered interface{})
	Error(action string, err error, message string)
	Warn(action, message string)
	//Debug(action, message, customMessage string, userId string, clientId string, data interface{}, additionalFields *map[string]interface{})
	Info(action, message string)
}

// Note заложена логика работы с логами
type app struct {
}

type messageStruct struct {
	Action  string
	Error   error
	Message string
}

type Options struct {
	CustomMessage    string
	Data             interface{}
	AdditionalFields *map[string]interface{}
}

func New() AppLogger {
	return &app{}
}

// HttpPanicHandler метод для обработки паники для GIN
func (l app) HttpPanicHandler(ctx *gin.Context, recovered interface{}) {
	code, mess := PanicErrorCode, fmt.Sprint(recovered)

	ctx.JSON(http.StatusUnprocessableEntity, errorToBase(code, mess, nil))
	return
}

func (l app) Error(action string, err error, message string) {
	m := messageStruct{
		Action:  action,
		Error:   err,
		Message: message,
	}
	errMes := fmt.Sprintf("%s: %s - %s", "ERROR:", m.Message, m.Error.Error())

	ErrorLogger = log.New(os.Stderr, errMes, log.Ldate|log.Ltime|log.Lshortfile)
}

func (l app) Warn(action, message string) {
	m := messageStruct{
		Action:  action,
		Message: message,
	}
	warnMes := fmt.Sprintf("%s: %s - %s", "WARNING:", m.Message)

	WarningLogger = log.New(os.Stdout, warnMes, log.Ldate|log.Ltime|log.Lshortfile)
}

func (l app) Info(action, message string) {
	m := messageStruct{
		Action:  action,
		Message: message,
	}
	infoMes := fmt.Sprintf("%s: %s - %s", "INFO:", m.Message)

	InfoLogger = log.New(os.Stdout, infoMes, log.Ldate|log.Ltime|log.Lshortfile)
}

//func (l app) Debug(action, message, customMessage string, userId string, clientId string, data interface{}, additionalFields *map[string]interface{}) {
//	l.log.Debug(loggergo.Message{
//		Action:           action,
//		Message:          &message,
//		CustomMessage:    &customMessage,
//		UserID:           &userId,
//		ClientID:         &clientId,
//		Data:             data,
//		AdditionalFields: additionalFields,
//	})
//}

// errorToBase - ответ об ошибке при обработке запроса
func errorToBase(code int, message string, data interface{}) common.BaseResponse {
	return common.BaseResponse{Code: code, Message: message, Data: data}
}
