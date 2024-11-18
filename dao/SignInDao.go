package dao

import (
	"starfall-go/entity"
	"starfall-go/util"
)

type SignInDao struct {
}

func (SignInDao) FindSignInByUserAndDate(user, date string) (signIn entity.SignIn) {
	util.DB.Table("sign_in").Where("user = ? and date = ?", user, date).Find(&signIn)
	return
}

func (SignInDao) FindAllSignInByUser(user string) (signIns []entity.SignIn) {
	util.DB.Table("sign_in").Where("user = ?", user).Order("date desc").Find(&signIns)
	return
}

func (SignInDao) FindAllSignInByUserAndOffset(user string, offset int) (signIns []entity.SignIn) {
	util.DB.Table("sign_in").Where("user = ?", user).Order("date desc").Limit(6).Offset(offset).Find(&signIns)
	return
}

func (SignInDao) CountSignInByUser(user string) (count int64) {
	util.DB.Table("sign_in").Where("user = ?", user).Find(&entity.SignIn{}).Count(&count)
	return
}

func (SignInDao) InsertSignIn(signIn entity.SignIn) bool {
	result := util.DB.Table("sign_in").Create(signIn).RowsAffected
	return util.Int64ToBool(result)
}
