package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Active bool   `json:"active"`
}

func MarshalStructToJSON(s interface{}) ([]byte, error) {
	v := reflect.ValueOf(s)

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("Эрор: ", v.Kind())
	}

	return json.Marshal(s)
}

func main() {
	user := User{Name: "Виктор", Age: 36, Active: true}

	jsonData, err := MarshalStructToJSON(user)
	if err != nil {
		fmt.Println("Ошибка сериализации: ", err)
		return
	}

	fmt.Print("JSON: ", string(jsonData))

}
