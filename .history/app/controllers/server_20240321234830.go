package controllers

import (
	"net/http"
	"todo_go/config"
)

func StartMainServer() error{
	return http.ListenAndServe(":" + config.Config.Port)
}
