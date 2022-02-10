package services

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
    "notify-service/config"
    "encoding/json"
    "log"
    "time"
)


type templateJson struct {
    Code string `json:"code"`
}

func SendSMSCodeByAliyun(phone string, code string) {
    client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", config.SMS.ALIYUN_ACCESS_KEY_ID, config.SMS.ALIYUN_ACCESS_KEY_SECRET)

    template := templateJson{code}
    templateStr, err := json.Marshal(template)

    request := dysmsapi.CreateSendSmsRequest()
    request.Scheme = "http"
    request.PhoneNumbers = phone
    request.SignName = config.SMS.ALIYUN_SMS_SIGN_NAME
    request.TemplateCode = config.SMS.ALIYUN_SMS_TEMPLATE_CODE
    request.TemplateParam = string(templateStr)

    begin := time.Now()
    response, err := client.SendSms(request)
    if err != nil {
        log.Println("error:", err)
    }
    log.Println("aliyun sms duration: ", time.Since(begin), "response:", response)
}
