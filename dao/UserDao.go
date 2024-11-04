package dao

import (
	"starfall-go/entity"
	"starfall-go/util"
)

var dbUser = util.DB.Table("user")

type User entity.User
type UserDao struct {
}

func (UserDao) FindUserWithUserOrEmail(account string) User {
	user := User{}
	dbUser.Where("user = ? or email = ?", account, account).Find(&user)
	return user
}

func (UserDao) FindUserWithUser(user string) User {
	userObj := User{}
	dbUser.Where("user = ?", user).Find(&userObj)
	return userObj
}
