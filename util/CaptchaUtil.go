package util

import (
	"encoding/json"
	"fmt"
	"github.com/mojocn/base64Captcha"
	"time"
)

type captcha struct {
	id      int
	base64s string
	num     string
}

var redisUtil = RedisUtil{}
var store = base64Captcha.DefaultMemStore

func CreateAndSaveCaptcha() (id, b64s, code string, err error) {
	driver := base64Captcha.NewDriverDigit(40, 80, 4, 0.6, 50)
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, code, err = captcha.Generate()
	captchaJSON, err := json.Marshal(captcha)
	if err != nil {
		return "", "", "", err
	}
	fmt.Println("code_id:", id, "code:", code)
	redisUtil.SetObjWithExpireTime("captcha:"+id, captchaJSON, 1*time.Minute)
	return
}

func VerifyCaptchaCode(id, key string) bool {
	flag := store.Verify(id, key, true)
	if flag {
		redisUtil.Del("captcha:" + id)
	}
	return flag
}
