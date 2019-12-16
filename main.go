package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func main() {
	// fmt.Println("Hello word")
	var jsonString = `{"name":"john wick","age":"27"}`
	var jsonData = []byte(jsonString)

	var data User
	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("user	:" + data.Name)
	fmt.Println("age		:" + data.Age)
}
