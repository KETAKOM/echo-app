package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewSqlHandler() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:33000)/echo_app?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	return db, nil
}

type Todo struct {
	ID     uint32 `json:"id" form:"id"`
	Title  string `json:"title" form:"title"`
	Detail string `json:"detail" form:"detail"`
}

func GetAll() ([]Todo, error) {
	db, err := NewSqlHandler()
	defer db.Close()
	if err != nil {
		fmt.Println("error", err)
		return []Todo{}, err
	}

	var todos []Todo
	db.Table("todo").Select("id, title, detail").Find(&todos)
	return todos, nil
}

func CreateTodo(t *Todo) error {
	db, err := NewSqlHandler()
	defer db.Close()
	if err != nil {
		fmt.Println("error", err)
		return err
	}

	db.Table("todo").Create(t)
	return nil
}

func UpdateTodo(t *Todo) error {
	db, err := NewSqlHandler()
	defer db.Close()
	if err != nil {
		fmt.Println("error", err)
		return err
	}

	db.Table("todo").Update(t)
	return nil
}

func DeleteTodo(t *Todo) error {
	db, err := NewSqlHandler()
	defer db.Close()
	if err != nil {
		fmt.Println("error", err)
		return err
	}

	db.Table("todo").Delete(t)
	return nil
}
