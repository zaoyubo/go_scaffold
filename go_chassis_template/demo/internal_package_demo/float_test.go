package internal_package_demo

import (
	"fmt"
	"testing"
)

func TestFloat(t *testing.T) {
	var a float64 = 1.5
	var b float64 = 1.3
	var result float64 = a - b

	if result == 0.2 {
		fmt.Println("相等")
	} else {
		fmt.Println("不相等")
	}
	fmt.Printf("%.20f\n ", result)

	res := IsEqual(result, 0.2)
	fmt.Println(res)
}

// MIN 为用户自定义的比较精度
const MIN = 0.000001
const MIN_NEG = -0.000001

func IsEqual(f1, f2 float64) bool {
	sub := f1 - f2
	if sub < 0 {
		return sub > MIN_NEG
	} else {
		return sub < MIN

	}
}
