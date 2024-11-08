package util

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"strings"
)

const emailSender = "icethestral@vip.qq.com"
const emailUrl = "smtp.qq.com"
const emailPort = "25"
const emailKey = "abqxaufiwxomifac"

var emailObj = email.NewEmail()

func init() {
	emailObj.From = "" + emailSender
}

func SendEmail(toEmail, title, content string) {
	emailObj.To = []string{toEmail}
	emailObj.Subject = title
	emailObj.HTML = []byte(content)
	err := emailObj.Send(emailUrl+":"+emailPort, smtp.PlainAuth("", emailSender, emailKey, emailUrl))
	if err != nil {
		fmt.Println(err)
	}
}

func SendRegEmailCode(toEmail, code string) {
	title := "StarFall注册验证码"
	content := `您好！<br>	感谢注册此网站，也感谢你的大力支持<br>注意：如您并未正在操作，请勿轻信任何索要验证码的坏人!<center>↓↓↓↓您的注册验证码↓↓↓↓</center><center style='font-size:40px;border:1px solid black;border-radius: 2px'>` + strings.ToUpper(code) + `</center><center>↑↑↑↑您的注册验证码↑↑↑↑</center>`
	SendEmail(toEmail, title, content)
}

func SendCustomEmailCode(toEmail, code, role string) {
	title := "StarFall" + role + "验证码"
	content := `您好！<br>	感谢注册此网站，也感谢你的大力支持<br>注意：如您并未正在操作，请勿轻信任何索要验证码的坏人!<center>↓↓↓↓您的注册验证码↓↓↓↓</center><center style='font-size:40px;border:1px solid black;border-radius: 2px'>` + strings.ToUpper(code) + `</center><center>↑↑↑↑您的` + role + `验证码↑↑↑↑</center>`
	SendEmail(toEmail, title, content)
}
