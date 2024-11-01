package entity

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Result struct {
	Msg    string `json:"msg,omitempty"`
	Object any    `json:"object,omitempty"`
}

func (Result) Ok() Result {
	return Result{
		Msg:    "success",
		Object: nil,
	}
}

func (Result) OkWithObj(object any) Result {
	return Result{
		Msg:    "success",
		Object: object,
	}
}

func (Result) OkWithMsgAndObj(msg string, object any) Result {
	return Result{
		Msg:    msg,
		Object: object,
	}
}

func (Result) Error() gin.H {
	return gin.H{
		"ERROR": "NULL",
		"time":  time.Now(),
	}
}

func (Result) ErrorWithMsg(msg string) gin.H {
	return gin.H{
		"ERROR": msg,
		"time":  time.Now(),
	}
}
