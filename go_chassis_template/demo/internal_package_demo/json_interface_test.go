package internal_package_demo

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestJson(t *testing.T) {
	type PP struct {
		Context context.Context `json:"context"`
	}
	ctx := context.WithValue(context.Background(), "k", "v")
	p1 := PP{
		Context: ctx,
	}
	result, err := json.MarshalIndent(p1, "", "")
	if err != nil {
		fmt.Println(err)
		panic("序列化失败")
	} else {
		fmt.Println("序列化结果::")
		fmt.Println(string(result))
	}
	//反序列化
	p2 := PP{}

	err2 := json.Unmarshal(result, &p2)
	if err2 != nil {
		fmt.Println(err2)
		panic("反序列化失败")
	} else {
		fmt.Println("反序列话结果:::")
		fmt.Println(p2)
	}
}
