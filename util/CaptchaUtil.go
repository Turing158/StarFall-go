package util

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
	"strings"
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
	//captchaJSON, err := json.Marshal(captcha)
	if err != nil {
		return "", "", "", err
	}
	fmt.Println("code_id:", id, "code:", code)
	//因为有store.Verify的原因，可以不使用redis缓存验证码
	//redisUtil.SetObjWithExpireTime("captcha:"+id, captchaJSON, 1*time.Minute)
	return
}

func VerifyCaptchaCode(id, key string) bool {
	flag := store.Verify(id, key, false)
	if flag {
		store.Verify(id, key, true)
	}
	//因为有store.Verify的原因，可以不使用redis缓存验证码
	//if flag {
	//	redisUtil.Del("captcha:" + id)
	//}
	return flag
}

func GetCodeAndIdByCode(codeStr string) (id, code string) {
	codeStrSplit := strings.Split(codeStr, ":")
	id = codeStrSplit[0]
	code = codeStrSplit[1]
	return
}
