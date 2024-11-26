package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"starfall-go/dao"
	"starfall-go/entity"
	"starfall-go/util"
	"strconv"
	"strings"
	"time"
)

var result = entity.Result{}
var userDao = dao.UserDao{}

type UserService struct {
}

func (UserService) Login(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")
	codeAndId := c.PostForm("code")
	codeId, code := util.GetCodeAndIdByCode(codeAndId)

	if util.VerifyCaptchaCode(codeId, code) {
		user := userDao.FindUserWithUserOrEmail(account)
		if user.User != "" {
			passwordE, _ := util.AesDecrypt(user.Password)
			if passwordE == password {
				token := util.GenerateToken(util.UserClaim{User: user.User, Email: user.Email, Role: user.Role})
				c.JSON(200, result.OkWithObj(token))
				return
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("password is wrong"))
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("account is not exist"))
		return
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("code is wrong"))
}

func (UserService) GetUserInfo(c *gin.Context) {
	token := c.GetHeader("Authorization")
	_, claim, _ := util.ParseToken(token)
	user := claim.User
	if user == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("The User is null in the token"))
		return
	}
	userObj := userDao.FindUserWithUser(user)
	userObj.MaxExp = util.GetMaxExp(userObj.Level)
	c.JSON(200, result.OkWithObj(userObj))
}

func (UserService) Register(c *gin.Context) {
	user := c.PostForm("user")
	password := c.PostForm("password")
	email := c.PostForm("email")
	emailCode := c.PostForm("emailCode")
	codeAndId := c.PostForm("code")
	codeId, code := util.GetCodeAndIdByCode(codeAndId)

	if util.VerifyCaptchaCode(codeId, code) {
		if userDao.FindUserWithEmail(email).User == "" {
			if userDao.FindUserWithUser(user).User == "" {
				emailCodeInRedis := redisUtil.Get("emailCode:" + email)
				if emailCodeInRedis == emailCode {
					passwordCode, _ := util.AesEncrypt(password)
					userObj := entity.User{
						User:     user,
						Password: passwordCode,
						Email:    email,
						Birthday: time.Now().Format("2006-01-02"),
						Gender:   0,
						Level:    1,
						Exp:      0,
						Name:     "新用户" + strconv.Itoa(time.Now().Nanosecond()),
						Role:     "user",
						Avatar:   "default.png",
					}
					userDao.Save(userObj)
					redisUtil.Del("emailCode:" + email)
					c.JSON(200, result.Ok())
					return
				}
				c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("The Email code is wrong"))
				return
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("The user already exists"))
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("The Email already exists"))
		return
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("The code is wrong"))
}

func getEmailCode(c *gin.Context, role string) {
	email := c.PostForm("email")
	code := strings.ToUpper(util.RandomStr(6))
	redisUtil.SetWithExpireTime("emailCode:"+email, code, 10*time.Minute)
	if role == "注册" {
		util.SendRegEmailCode(email, code)
	} else {
		util.SendCustomEmailCode(email, code, role)
	}
	c.JSON(200, result.Ok())
}

func existEmail(email string) bool {
	if userDao.FindUserWithEmail(email).User != "" {
		return true
	}
	return false
}

func (UserService) GetRegEmailCode(c *gin.Context) {
	email := c.PostForm("email")
	if email == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("The email cannot be empty"))
	}
	if existEmail(email) {
		c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("The Email is already exists"))
		return
	}
	getEmailCode(c, "注册")
}

func (UserService) GetForgerPasswordEmailCode(c *gin.Context) {
	email := c.PostForm("email")
	if email == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("The email cannot be empty"))
	}
	if !existEmail(email) {
		c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("The Email is not exists"))
		return
	}
	getEmailCode(c, "忘记密码")
}

func (UserService) GetNewEmailCode(c *gin.Context) {
	getEmailCode(c, "新邮箱")
}

func (UserService) GetOldEmailCode(c *gin.Context) {
	getEmailCode(c, "旧邮箱")
}

func (UserService) CheckForgetPassword(c *gin.Context) {
	email := c.PostForm("email")
	emailCode := c.PostForm("emailCode")
	codeAndId := c.PostForm("code")
	id, code := util.GetCodeAndIdByCode(codeAndId)

	if util.VerifyCaptchaCode(id, code) {
		user := userDao.FindUserWithEmail(email)
		if !redisUtil.Has("emailCode:" + email) {
			c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("Email verification code has expired or not been sent"))
			return
		}
		redisEmailCode := redisUtil.Get("emailCode:" + email).(string)
		if strings.ToUpper(redisEmailCode) == strings.ToUpper(emailCode) {
			if user.User != "" {
				token := util.GenerateTokenWithExpire(util.UserClaim{
					User:  "",
					Email: email,
					Role:  "@ForgetPassword",
				}, time.Minute*5)
				redisUtil.Del("email:" + email)
				c.JSON(200, result.OkWithObj(token))
				return
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("The Email is not exists"))
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("The Email code is wrong"))
		return
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("The code is wrong"))
}

func (UserService) ForgerPassword(c *gin.Context) {
	token := c.GetHeader("Authorization")
	password := c.PostForm("password")
	_, claim, _ := util.ParseToken(token)
	if claim.Role != "@ForgetPassword" {
		c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("Not allow"))
		return
	}
	user := userDao.FindUserWithUserOrEmail(claim.Email)
	if user.User == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("No user exists"))
		return
	}
	passwordAes, _ := util.AesEncrypt(password)
	userDao.UpdatePassword(user.User, passwordAes)
	c.JSON(200, result.Ok())
}

func (UserService) SettingInfo(c *gin.Context) {
	b, _ := c.GetRawData()
	var data map[string]interface{}
	json.Unmarshal(b, &data)
	codeStr := data["code"].(string)
	id, code := util.GetCodeAndIdByCode(codeStr)
	if util.VerifyCaptchaCode(id, code) {
		token := c.GetHeader("Authorization")
		_, claim, _ := util.ParseToken(token)
		//更新用户操作
		user := userDao.FindUserWithUser(claim.User)
		user.MaxExp = util.GetMaxExp(user.Level)
		c.JSON(200, result.OkWithObj(user))
		return
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("The code is wrong"))
}

func (UserService) SettingPassword(c *gin.Context) {
	token := c.GetHeader("Authorization")
	oldPassword := c.PostForm("oldPassword")
	newPassword := c.PostForm("newPassword")
	codeStr := c.PostForm("code")
	id, code := util.GetCodeAndIdByCode(codeStr)
	_, claim, _ := util.ParseToken(token)
	if util.VerifyCaptchaCode(id, code) {
		user := userDao.FindUserWithUser(claim.User)
		if user.User == "" {
			password, _ := util.AesDecrypt(user.Password)
			if password == oldPassword {
				newPasswordAes, _ := util.AesEncrypt(newPassword)
				userDao.UpdatePassword(user.User, newPasswordAes)
				c.JSON(200, result.Ok())
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("The OldPassword is wrong"))
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("The token is not exist user"))
		return
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("The code is wrong"))
}

func (UserService) SettingAvatar(c *gin.Context) {

}

func (UserService) SettingEmail(c *gin.Context) {
	token := c.GetHeader("Authorization")
	_, claim, _ := util.ParseToken(token)
	newEmail := c.PostForm("newEmail")
	oldEmailCode := strings.ToUpper(c.PostForm("oldEmailCode"))
	newEmailCode := strings.ToUpper(c.PostForm("newEmailCode"))
	redisOldEmailCode := redisUtil.Get("email:" + claim.Email).(string)
	redisNewEmailCode := redisUtil.Get("email:" + newEmail).(string)
	if oldEmailCode == redisOldEmailCode {
		if newEmailCode == redisNewEmailCode {
			tmpUser := userDao.FindUserWithEmail(newEmail)
			if tmpUser.Email == "" {
				userDao.UpdateEmail(claim.User, newEmail)
				newClaim := util.UserClaim{
					User:  tmpUser.User,
					Email: tmpUser.Email,
					Role:  tmpUser.Role,
				}
				newToken := util.GenerateToken(newClaim)
				c.JSON(200, result.OkWithObj(newToken))
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("The Old Email code is wrong"))
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("The New Email code is wrong"))
		return
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("The Old Email code is wrong"))
}

func (UserService) FindUserByUser(c *gin.Context) {
	user := c.PostForm("user")
	userObj := userDao.FindUserWithUser(user)
	if userObj.User != "" {
		userObj.MaxExp = util.GetMaxExp(userObj.Level)
		c.JSON(200, result.OkWithObj(userObj))
		return
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("The User is not exist"))
}

var signInDao = dao.SignInDao{}

func (UserService) FindAlreadySignIn(c *gin.Context) {
	token := c.GetHeader("Authorization")
	_, claim, _ := util.ParseToken(token)
	pageStr := c.PostForm("page")
	page, error := strconv.Atoi(pageStr)
	if error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("The page is not a number"))
		return
	}
	signIns := signInDao.FindAllSignInByUserAndOffset(claim.User, (page-1)*6)
	var err bool
	var count int
	if len(signIns) == 1 {
		count = 1
	} else if len(signIns) == 0 {
		count = 0
	} else {
		count = 1
		timeStr := "2006-01-02"
		for i := 0; i < len(signIns)-1; i++ {
			newDate, err1 := time.Parse(timeStr, signIns[i].Date)
			oldDate, err2 := time.Parse(timeStr, signIns[i+1].Date)
			if err1 != nil || err2 != nil {
				err = true
				continue
			}
			if util.IsContinualDate(oldDate, newDate) {
				count++
			} else {
				break
			}
		}
	}
	var r entity.Result
	if err {
		r = result.OkWithObj(gin.H{
			"signIns":        signIns,
			"count":          signInDao.CountSignInByUser(claim.User),
			"continualCount": count,
			"error":          "The date data format is incorrect and will only affect the display of consecutive days",
		})
	} else {
		r = result.OkWithObj(gin.H{
			"signIns":        signIns,
			"count":          signInDao.CountSignInByUser(claim.User),
			"continualCount": count,
		})
	}
	c.JSON(200, r)
}

func (UserService) SignIn(c *gin.Context) {
	msg := c.PostForm("msg")
	emotion := c.PostForm("emotion")
	token := c.GetHeader("Authorization")
	_, claim, _ := util.ParseToken(token)
	date := time.Now().Format("2006-01-02")
	signIn := signInDao.FindSignInByUserAndDate(claim.User, date)
	if signIn.User == "" {
		addExp := int64(rand.Intn(50) + 20)
		user := userDao.FindUserWithUser(claim.User)
		exp := user.Exp + addExp
		level := user.Level
		diffExp := util.CheckAndLevelUp(exp, level)
		if diffExp >= 0 {
			exp = diffExp
			level++
		}
		msg = "[获得" + strconv.Itoa(int(addExp)) + "点经验] " + msg
		signInDao.InsertSignIn(entity.SignIn{
			User:    user.User,
			Date:    date,
			Message: msg,
			Emotion: emotion,
		})
		userDao.UpdateExp(user.User, exp, level)
		c.JSON(200, result.OkWithObj(addExp))
		return
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, result.ErrorWithMsg("You have already signed it today"))
}
