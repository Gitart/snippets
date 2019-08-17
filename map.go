package main

import (
	"fmt"
	"reflect"
)

func main() {
	myData := make(map[string]interface{})
	myData["Name"] = "Tony"
	myData["Age"] = 23

	result := &MyStruct{}
	FillStruct(myData, result)

	fmt.Println(result.Name)
}

func FillStruct(data map[string]interface{}, result interface{}) {
	t := reflect.ValueOf(result).Elem()
	for k, v := range data {
		val := t.FieldByName(k)
		val.Set(reflect.ValueOf(v))
	}
}

type MyStruct struct {
	Name string
	Age  int
}
