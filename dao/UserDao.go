package dao

import (
	"fmt"
	"starfall-go/entity"
	"starfall-go/util"
	"strconv"
)

var dbUser = util.DB.Table("user")

type UserDao struct {
}

func (UserDao) FindUserWithUserOrEmail(account string) (user entity.UserOut) {
	dbUser.Where("user = ? or email = ?", account, account).Find(&user)
	return
}

func (UserDao) FindUserWithUser(user string) (userObj entity.UserOut) {
	dbUser.Where("user = ?", user).Find(&userObj)
	return userObj
}

func (UserDao) FindUserWithEmail(email string) (user entity.UserOut) {
	dbUser.Where("email = ?", email).Find(&user)
	return
}

func (UserDao) Save(user entity.User) bool {
	err := dbUser.Save(&user).Error
	if err != nil {
		fmt.Println(err)
	}
	return err == nil
}

func (UserDao) UpdateExp(user string, exp, level int64) bool {
	var userObj entity.User
	row := dbUser.Where("user = ?", user).First(&userObj).Updates(entity.User{Exp: int64(exp), Level: int64(level)}).RowsAffected
	return util.Int64ToBool(row)
}

func (UserDao) UpdateAvatar(user, avatar string) bool {
	var userObj entity.User
	row := dbUser.Where("user = ?", user).First(&userObj).Update("avatar", avatar).RowsAffected
	return util.Int64ToBool(row)
}

func (UserDao) UpdateInfo(user, name, gender, birthday string) bool {
	var userObj entity.User
	genderInt, _ := strconv.Atoi(gender)
	row := dbUser.Where("user = ?", user).First(&userObj).Updates(entity.User{Name: name, Gender: int64(genderInt), Birthday: birthday}).RowsAffected
	return util.Int64ToBool(row)
}

func (UserDao) UpdatePassword(user, password string) bool {
	var userObj entity.User
	row := dbUser.Where("user = ?", user).First(&userObj).Update("password", password).RowsAffected
	return util.Int64ToBool(row)
}

func (UserDao) UpdateEmail(user, email string) bool {
	var userObj entity.User
	row := dbUser.Where("user = ?", user).First(&userObj).Update("email", email).RowsAffected
	return util.Int64ToBool(row)
}
