package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sotomskir/go-todos/models"
	u "github.com/sotomskir/go-todos/utils"
	"net/http"
	"strconv"
)

var CreateTodo = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user") . (uint)
	todo := &models.Todo{}

	err := json.NewDecoder(r.Body).Decode(todo)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	todo.UserId = user
	resp := todo.Create()
	u.Respond(w, resp)
}

var GetTodosFor = func(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user") . (uint)
	data := models.GetTodos(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var DeleteTodo = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := r.Context().Value("user") . (uint)
	id, _ := strconv.ParseInt(vars["id"], 10, 8)
	data := models.DeleteTodo(userId, id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
