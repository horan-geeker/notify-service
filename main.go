package main

import (
    "net/http"
    "log"
    "notify-service/controllers"
)

func main() {
    http.HandleFunc("/send/sms", controllers.SendSmsCode)
    http.HandleFunc("/send/mail", controllers.SendMail)
    log.Println("server start")
    log.Fatal(http.ListenAndServe(":80", nil))
}
