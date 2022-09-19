package external_package_demo

import (
	"fmt"
	"testing"
	"time"
)

const Secret = "abc"

func TestGenToken(t *testing.T) {
	type User struct {
		Username string `json:"username"`
	}
	now := time.Now().Unix()
	claims := jwt.MapClaims{
		"timestamp": now,
		"data":      User{"zhangsan"}, // 会以map格式序列化，而不是json
		//"data2":      data, 可任意添加，自定义数据
		//"exp": now, // 过期时间，解析时会校验。exp 为指定参数，不能改其他名字
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(Secret))
	if err != nil {
		return
	}
	fmt.Println(tokenString)
	/*
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7InVzZXJuYW1lIjoiemhhbmdzYW4ifSwidGltZXN0YW1wIjoxNjYxODI3NDEwfQ.JGnQiLGnsr_UTXSt42mYc6A_zHRi1TTApPoS9SMw85A"
	*/
}

func TestParseToken(t *testing.T) {
	jwtStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7InVzZXJuYW1lIjoiemhhbmdzYW4ifSwidGltZXN0YW1wIjoxNjYxODI3NDEwfQ.JGnQiLGnsr_UTXSt42mYc6A_zHRi1TTApPoS9SMw85A"
	//jwtStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7InVzZXJuYW1lIjoiemhhbmdzYW4ifSwiZXhwIjoxNjYxODMyNTYzLCJ0aW1lc3RhbXAiOjE2NjE4MzI1NjN9.ABok7vuTUvuxrnmPjFhSFyPvId30Ctoedz6Ts0bIL-A"// 会过期的 jwt
	jwtToken, err := jwt.ParseWithClaims(jwtStr, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		t.Error(err.Error())
	}
	if !jwtToken.Valid {
		return
	}
	mapClaims := jwtToken.Claims.(jwt.MapClaims)
	fmt.Printf("claims:%+v", mapClaims)
	data := mapClaims["data"]
	fmt.Println(data)
	dat_map := data.(map[string]interface{}) // 会以map格式序列化，而不是json
	fmt.Println(dat_map["username"])
	/*
		claims:map[data:map[username:zhangsan] timestamp:1.66182741e+09]map[username:zhangsan]
		zhangsan
	*/
}
