#!/bin/sh

SCRIPT_DIR=$(dirname "$0")
cd "$SCRIPT_DIR" || exit

go get -u github.com/gin-gonic/gin
go get -u github.com/gin-contrib/cors
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get -u github.com/dgrijalva/jwt-go
go get -u github.com/go-redis/redis
go get -u github.com/mojocn/base64Captcha
go get -u github.com/jordan-wright/email
go get -u github.com/gorilla/websocket