package internal_package_demo

import (
	"fmt"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

type Person struct {
	Name   string `json:"name"`
	Age    *int   `json:"age,omitempty"`
	Number int    `json:"number"`
}

func TestJsonMarshal(t *testing.T) {
	p := Person{
		Name: "zhangsan",
		//Age:  1,
	}
	var l []Person
	l = append(l, p)
	b, _ := jsoniter.Marshal(l)
	fmt.Println(string(b)) //[{"name":"zhangsan","age":1}]
}

func TestJsonUnmarshal(t *testing.T) {
	//s1 := "{\"name\":\"zhangsan\"}"
	s1 := "{\"name\":\"zhangsan\",\"age\":1, \"other\":true}"
	var p Person
	err := jsoniter.Unmarshal([]byte(s1), &p)
	if err != nil {
		fmt.Println("err====>", err)
	}
	fmt.Printf("%+v\n", p)

}
