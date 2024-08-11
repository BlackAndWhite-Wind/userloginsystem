package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB
var JWTSecretKey = []byte("your_secret_key")
var SMTPHost = "smtp.example.com"
var SMTPPort = 587
var SMTPUser = "874321762@qq.com"
var SMTPPassword = "qwer697410"
var SmsApiKey = "your-sms-api-key"                 // 短信服务 API 密钥
var SmsApiUrl = "https://sms-api.example.com/send" // 短信服务 API URL

func InitDB() {
	dsn := "root:qwer697410@tcp(127.0.0.1:3306)/userlogin?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}
}
