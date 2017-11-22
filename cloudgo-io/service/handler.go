package service

import (
	"net/http"
	"github.com/render"
)


func jsHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        formatter.JSON(w, http.StatusOK, struct {
            Title      string `json:"title"`
        }{Title : "Wechat"})
    }
}

func htmlHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, r *http.Request) {
        r.ParseForm()
        formatter.HTML(w, http.StatusOK, "table", struct {
            Name     string
            Email    string
            Phone    string
        }{Name: r.Form["username"][0], Email: r.Form["email"][0], Phone: r.Form["phone"][0]})
    }
}

func unknowHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
            formatter.JSON(w, http.StatusNotImplemented, "501 Not Implemented")}
}