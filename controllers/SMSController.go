package controllers

import (
    "net/http"
    "io"
    "notify-service/library/response"
    "log"
    "notify-service/services"
)

func SendSmsCode(w http.ResponseWriter, r *http.Request) {
    params := r.URL.Query()
    phones, ok := params["phone"]
    if !ok || len(phones[0]) <= 0{
        log.Println("Url Param 'phone' is missing")
        io.WriteString(w, response.JsonParamError("Url Param 'phone' is missing"))
        return
    }
    phone := phones[0]
    codes, ok := params["code"]
    if !ok || len(codes[0]) <= 0{
        log.Println("Url Param 'code' is missing")
        io.WriteString(w, response.JsonParamError("Url Param 'code' is missing"))
        return
    }
    code := codes[0]
    go services.SendSMSCodeByAliyun(phone, code)
    io.WriteString(w, response.JsonDone())
}
