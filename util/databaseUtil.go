package util

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dbName = "starfall"
const port = "3306"
const username = "root"
const password = "xwh2003"

const dns = username + ":" + password + "@tcp(localhost:" + port + ")/" + dbName

var DB *gorm.DB

func init() {
	db, err := gorm.Open(mysql.Open(dns))
	if err != nil {
		fmt.Println(err)
		return
	}
	DB = db
}
