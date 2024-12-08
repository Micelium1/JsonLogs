package service

import (
	"net/http"
	"text/template"
)

func FrontInit(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/front.html")
	if err != nil {
		http.Error(w, "Ошибка загрузки фронта", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
