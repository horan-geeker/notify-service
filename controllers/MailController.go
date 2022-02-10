package controllers

import (
    "net/http"
    "notify-service/services"
    "io"
    "notify-service/library/response"
)

func SendMail(w http.ResponseWriter, r *http.Request) {
    params := r.URL.Query()
    emails, ok := params["email"]
    if !ok || len(emails[0]) <= 0{
        io.WriteString(w, response.JsonParamError("Url Param 'email' is missing"))
        return
    }
    email := emails[0]
    froms, ok := params["from"]
    if !ok || len(froms[0]) <= 0{
        io.WriteString(w, response.JsonParamError("Url Param 'from' is missing"))
        return
    }
    from := froms[0]
    titles, ok := params["title"]
    if !ok || len(titles[0]) <= 0{
        io.WriteString(w, response.JsonParamError("Url Param 'title' is missing"))
        return
    }
    title := titles[0]
    contents, ok := params["content"]
    if !ok || len(contents[0]) <= 0{
        io.WriteString(w, response.JsonParamError("Url Param 'content' is missing"))
        return
    }
    content := contents[0]
    go services.SendMailByQQ(email, from, title, content)
    io.WriteString(w, response.JsonDone())
}