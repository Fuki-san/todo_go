package controllers

import (
	"html/template"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request){
	t, err:= template.ParseFiles("app/views/templates/top.html")
	if err != nil{
		
	}
	t.Execute(w,"Hello World!")
} 
