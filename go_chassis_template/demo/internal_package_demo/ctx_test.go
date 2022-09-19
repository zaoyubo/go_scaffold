package internal_package_demo

import (
	"context"
	"fmt"
	"testing"
)

func TestCtx(t *testing.T) {
	ctx := context.Background()
	//ctx = context.WithValue(ctx, "country", "BR")
	v, ok := ctx.Value("country").(string)
	if !ok {
		fmt.Println("country convert non ok")
		fmt.Println(v)
	} else {
		fmt.Println(v)
	}

}
