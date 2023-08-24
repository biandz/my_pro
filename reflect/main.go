package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p Person) Run(name string) (string, int) {
	str := fmt.Sprintf("%s is run....", name)
	fmt.Println(str)
	return str, 111
}

func (p Person) Eat(name string) (string, int) {
	str := fmt.Sprintf("%s is eat....", name)
	fmt.Println(str)
	return str, 222
}

func main() {
	//获取结构体结构相关数据
	var p = Person{Name: "bdz", Age: 18}
	tf := reflect.TypeOf(p)
	for i := 0; i < tf.NumField(); i++ {
		fmt.Println(tf.Field(i).Tag.Get("json"))
	}

	//获取结构体下的方法
	for i := 0; i < tf.NumMethod(); i++ {
		fmt.Println("methodName:", tf.Method(i).Name)
		fmt.Println("methodType:", tf.Method(i).Type)
		fmt.Println("methodIndex:", tf.Method(i).Index)
	}

	//获取具体实例的相关数据
	vf := reflect.ValueOf(p)
	for i := 0; i < vf.NumField(); i++ {
		a := vf.Field(i).Interface()
		switch a.(type) {
		case int:
			fmt.Println("int:", a.(int))
		case string:
			fmt.Println("string:", a.(string))
		}
	}

	//调用具体实例下的方法
	for i := 0; i < tf.NumMethod(); i++ {
		name := vf.FieldByName("Name").String()
		call := vf.MethodByName(tf.Method(i).Name).Call([]reflect.Value{reflect.ValueOf(name)})
		for _, value := range call {
			switch value.Kind() {
			case reflect.String:
				fmt.Println("打印返回值：", value.String())
			case reflect.Int:
				fmt.Println("打印返回值：", value.Int())
			}
		}
	}
}
