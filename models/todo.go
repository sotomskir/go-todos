package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	u "github.com/sotomskir/go-todos/utils"
)

type Todo struct {
	gorm.Model
	Name string `json:"name"`
	UserId uint `json:"user_id"`
}

/*
 This struct function validate the required parameters sent through the http request body

returns message and true if the requirement is met
*/
func (todo *Todo) Validate() (map[string] interface{}, bool) {

	if todo.Name == "" {
		return u.Message(false, "Contact name should be on the payload"), false
	}

	if todo.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (todo *Todo) Create() (map[string] interface{}) {

	if resp, ok := todo.Validate(); !ok {
		return resp
	}

	GetDB().Create(todo)

	resp := u.Message(true, "success")
	resp["data"] = todo
	return resp
}

func GetTodo(id uint) (*Todo) {

	contact := &Todo{}
	err := GetDB().Table("todos").Where("id = ?", id).First(contact).Error
	if err != nil {
		return nil
	}
	return contact
}

func GetTodoForUser(id int64, userId uint) (*Todo) {

	contact := &Todo{}
	err := GetDB().Table("todos").Where("id = ? and user_id = ?", id, userId).First(contact).Error
	if err != nil {
		return nil
	}
	return contact
}

func GetTodos(user uint) ([]*Todo) {

	todos := make([]*Todo, 0)
	err := GetDB().Table("todos").Where("user_id = ?", user).Find(&todos).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return todos
}

func DeleteTodo(user uint, id int64) (error) {
	todo := GetTodoForUser(id, user)
	err := GetDB().Table("todos").Delete(todo).Error
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

