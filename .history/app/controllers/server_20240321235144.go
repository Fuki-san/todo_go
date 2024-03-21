package controllers

import (
	"net/http"
	"todo_go/config"
)

func StartMainServer() error{
	http.HandleFunc("/", func())
	return http.ListenAndServe(":" + config.Config.Port, nil)
}
