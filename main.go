package main

import (
	"errors"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func main() {
	// fmt.Println("Hello word")
	// var jsonString = `{"name":"john wicks","age":"27"}`
	// var jsonData = []byte(jsonString)

	// var data User
	// var err = json.Unmarshal(jsonData, &data)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Println("user	:" + data.Name)
	// fmt.Println("age		:" + data.Age)

	// var slice = []string{"a", "b", "c", "d"}

	// var slice2 = slice[1:2]

	// fmt.Println(slice)
	// fmt.Println(slice2)

	// slice = []string{"1", "2", "3"}

	// fmt.Println(slice)
	// fmt.Println(slice2)

	// number := 10.0

	// result, err := calculate(number)

	// fmt.Println(result)
	// fmt.Println(err)
	// pointer()
	x := 5
	y := 6
	add(x, y)
	fmt.Println(min(x, y))
	fmt.Println(multiple(x, y))
	fmt.Println(division(x, y))
}

func calculate(d float64) (float64, error) {
	if d == 8 {
		return 0, errors.New("tidak sama dengan 8")
	}
	cal := d + 10
	return cal, nil
}

func pointer() {
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value) :", numberA)
	fmt.Println("numberA (address) :", &numberA)
	fmt.Println("numberB (value) :", *numberB)
	fmt.Println("numberB (address) :", numberB)

	fmt.Println("")

	numberA = 5

	fmt.Println("numberA (value) :", numberA)
	fmt.Println("numberA (address) :", &numberA)
	fmt.Println("numberB (value) :", *numberB)
	fmt.Println("numberB (address) :", numberB)

}

func add(x int, y int) {
	fmt.Println(x + y)
}

func min(x int, y int) int {
	return x - y
}

func multiple(x int, y int) int {
	return x * y
}

func division(x int, y int) int {
	return x / y
}
