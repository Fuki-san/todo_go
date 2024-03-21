package controllers

import (
	"html/template"
	"net/http"

	"github.com/labstack/gommon/log"
)

func top(w http.ResponseWriter, r *http.Request){
	t, err:= template.ParseFiles("app/views/templates/top.html")
	if err != nil{
		log.Fatal(err)
	}
	t.Execute(w, {{ .Title }})
} 
