package controllers

import (
	"MyToDo_backend/config"
	"net/http"
)

func StartMainServer() error {
	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
