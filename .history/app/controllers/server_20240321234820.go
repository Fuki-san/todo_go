package controllers

import "net/http"

func StartMainServer() error{
	return http.ListenAndServe(":")
}
