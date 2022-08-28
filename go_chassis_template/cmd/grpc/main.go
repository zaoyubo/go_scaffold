package main

import (
	"log"

	chassis "github.com/go-chassis/go-chassis/v2"
	"template/apps/grpc"
	_ "template/initialize"
)

//if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/{path}/{to}/rest/server/
func main() {
	api := grpc.Init()
	api.Register()
	if err := chassis.Init(); err != nil {
		log.Fatal("Init failed." + err.Error())
		return
	}
	chassis.Run()
}
