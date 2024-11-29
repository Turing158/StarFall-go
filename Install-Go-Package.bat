@echo off
cd /d "%~dp0"
echo The current directory for the Go project is %~dp0

@echo on
echo Downloading and installing: github.com/gin-contrib/cors
go get -u github.com/gin-contrib/cors

echo Downloading and installing: gorm.io/gorm
go get -u gorm.io/gorm

echo Downloading and installing: gorm.io/driver/mysql
go get -u gorm.io/driver/mysql

echo Downloading and installing: github.com/dgrijalva/jwt-go
go get -u github.com/dgrijalva/jwt-go

echo Downloading and installing: github.com/go-redis/redis
go get -u github.com/go-redis/redis

echo Downloading and installing: github.com/mojocn/base64Captcha
go get -u github.com/mojocn/base64Captcha

echo Downloading and installing: github.com/jordan-wright/email
go get -u github.com/jordan-wright/email

echo Downloading and installing: github.com/gorilla/websocket
go get -u github.com/gorilla/websocket

pause
