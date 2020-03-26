package main

import (
	"fmt"
	"github.com/chaechaep/simulator/config"
	"github.com/chaechaep/simulator/object"
)

func main() {

	confFile := "C://Users/user/GolandProjects/go/src/github.com/chaechaep/simulator/config.json"
	err := config.Init(confFile)
	if err != nil {
		fmt.Println("config load failed : ", err)
		return
	}

	user := &object.User{
		UserId: "testtesttest5",
		//UserId:		"testtest1",
		AccessToken: "",
		//Password:    "testtest",
		Password: config.Cfg.DefaultPassword,
	}
	//ret, err := event.GetRegAvailable("testtesttest")
	//fmt.Println(ret)
	user.Login()
	user.GetJoinedRooms()
	user.SendMessage("m.text", "dlrjqhsofrjek")
	user.Logout()

}
