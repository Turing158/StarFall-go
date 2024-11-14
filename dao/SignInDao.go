package dao

import (
	"starfall-go/entity"
	"starfall-go/util"
)

type SignInDao struct {
}

var dbSignIn = util.DB.Table("sign_in")

func (SignInDao) FindSignInByUserAndDate(user, date string) (signIn entity.SignIn) {
	dbSignIn.Where("user = ? and date = ?", user, date).Find(&signIn)
	return
}

func (SignInDao) FindAllSignInByUser(user string) (signIns []entity.SignIn) {
	dbSignIn.Where("user = ?", user).Order("date desc").Find(&signIns)
	return
}

func (SignInDao) FindAllSignInByUserAndOffset(user string, offset int) (signIns []entity.SignIn) {
	dbSignIn.Where("user = ?", user).Order("date desc").Limit(6).Offset(offset).Find(&signIns)
	return
}

func (SignInDao) CountSignInByUser(user string) (count int64) {
	dbSignIn.Where("user = ?", user).Find(&entity.SignIn{}).Count(&count)
	return
}

func (SignInDao) InsertSignIn(signIn entity.SignIn) bool {
	result := dbSignIn.Create(signIn).RowsAffected
	return util.Int64ToBool(result)
}
