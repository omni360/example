package main

import "reflect"
import "fmt"

type User struct {
	Id    int
	Name  string
	Age   int
	Title string
}

func (this User) Test() string {
	return this.Name
}

var intva int = 33
var strva string = "dsds"
var useva User

func main() {
	var val interface{} = &User{1, "Tom", 12, "nan"}
	v := reflect.ValueOf(val)
	fmt.Println(v)
	m := v.MethodByName("Test")
	rets := m.Call([]reflect.Value{})
	fmt.Println(rets)
	fmt.Println(reflect.ValueOf(intva))
	fmt.Println(reflect.ValueOf(strva))
	fmt.Println(reflect.ValueOf(useva))
}
