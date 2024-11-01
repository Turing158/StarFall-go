package dao

import "starfall-go/util"

var db = util.DB.Table("user")

type UserDao struct {
}
