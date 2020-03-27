package main

import (
	"fmt"
	"github.com/chaechaep/simulator/config"
	"github.com/chaechaep/simulator/object"
	"math/rand"
	"time"
)

func Start(userId string) {
	user := object.User{
		UserId:      userId,
		AccessToken: "",
		Password:    config.Cfg.DefaultPassword,
		DeviceId:    "",
	}

	user.Login()
	for {
		r := rand.Intn(100)
		time.Sleep(time.Duration(r) * time.Millisecond)
		user.SendMessage("m.text", "Msg test")
	}

}
func main() {
	//userList := []string {
	//	"testtest3", "testtest4", "testtest5",
	//}

	confFile := "C://Users/user/GolandProjects/go/src/github.com/chaechaep/simulator/config.json"
	err := config.Init(confFile)
	if err != nil {
		fmt.Println("config load failed : ", err)
		return
	}
	//for _, user := range userList{
	//	go Start(user)
	//}
	//
	//fmt.Scanln()

	user := object.User{
		UserId:      "testtest1",
		AccessToken: "",
		Password:    "testtest",
		DeviceId:    "",
	}
	user.Login()
	user.GetJoinedRooms()
}
