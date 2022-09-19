package internal_package_demo

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	var f float64 = 1.22
	s := strconv.FormatFloat(f, 'f', 10, 64)
	agreedCompensation := strings.TrimRight(s, "0")
	if strings.HasSuffix(agreedCompensation, ".") {
		agreedCompensation = strings.TrimRight(agreedCompensation, ".")
	}
	fmt.Println(agreedCompensation)
}
