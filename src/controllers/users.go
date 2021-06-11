package controllers

import "net/http"

func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("oi"))
}
