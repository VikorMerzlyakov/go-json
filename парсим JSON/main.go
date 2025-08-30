package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsm map[string]interface{}
	jsonStr := `{
	"id": 1,
  "data": {
    "title": "Hello, Go!",
    "tags": ["json", "tutorial", "golang"],
    "views": 150
  },
  "active": true
}`
	err := json.Unmarshal([]byte(jsonStr), &jsm)
	if err != nil {
		panic(err)
	}

	switch v := jsm["id"].(type) {
	case float64:
		fmt.Println("ID: ", int(v))
	default:
		fmt.Println("Хуй а не ид")
	}

	switch v := jsm["active"].(type) {
	case bool:
		fmt.Println("Active: ", bool(v))
	default:
		fmt.Println("Хуй а не bool")
	}

	switch d := jsm["data"].(type) {
	case map[string]interface{}:
		switch t := d["title"].(type) {
		case string:
			fmt.Println("Title: ", string(t))
		default:
			fmt.Println("Хуй а не title")
		}
		switch tag := d["tags"].(type) {
		case []interface{}:
			fmt.Println("Tag: ")
			for _, tg := range tag {
				fmt.Print(tg, " ")
			}
		default:
			fmt.Println("Хуй а не tags")

		}

	default:
		fmt.Println("Хуй а не data")
	}

}
