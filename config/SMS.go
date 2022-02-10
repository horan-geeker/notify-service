package config

import (
    "github.com/joho/godotenv"
    "github.com/jinzhu/configor"
)

var SMS = struct {
    ALIYUN_ACCESS_KEY_ID string
    ALIYUN_ACCESS_KEY_SECRET string
    ALIYUN_SMS_SIGN_NAME string
    ALIYUN_SMS_TEMPLATE_CODE string
}{}

func init() {
    err := godotenv.Load()
    if err != nil {
        panic("Error loading .env file")
    }
    configor.Load(&SMS)
}