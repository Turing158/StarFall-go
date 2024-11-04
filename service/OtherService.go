package service

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"starfall-go/util"
	"strconv"
	"strings"
)

type OtherService struct {
}

var redisUtil = util.RedisUtil{}

func (OtherService) GetCodeImage(c *gin.Context) {
	oldId := c.Query("captcha_id")
	if oldId != "" {
		redisUtil.Del("captcha:" + oldId)
	}
	id, base64s, _, err := util.CreateAndSaveCaptcha()
	if err != nil {
		return
	}
	data, err := base64.StdEncoding.DecodeString(strings.Split(base64s, ",")[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Header("Content-length", strconv.Itoa(len(data)))
	c.Header("captcha_id", id)
	c.Data(200, "image/png", data)
}
