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
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func GetAll() ([]Todo, error) {
	db, err := NewSqlHandler()
	defer db.Close()
	if err != nil {
		fmt.Println("error", err)
		return []Todo{}, err
	}

	var todos []Todo
	db.Table("todo").Select("title, detail").Find(&todos)
	return todos, nil
}
