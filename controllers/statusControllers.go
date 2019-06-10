package controllers

import (
	u "github.com/sotomskir/go-todos/utils"
	"net/http"
)

var Status = func(w http.ResponseWriter, r *http.Request) {
	resp := u.Message(true, "success")
	u.Respond(w, resp)
}
