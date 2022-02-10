package config

import (
    "github.com/joho/godotenv"
    "github.com/jinzhu/configor"
)

var Mail = struct {
    TENCENT_MAIL_ADDRESS string
    TENCENT_MAIL_PASSWORD string
}{}

func init() {
    err := godotenv.Load()
    if err != nil {
        panic("Error loading .env file")
    }
    configor.Load(&Mail)
}