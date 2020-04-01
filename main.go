package main

import (
	"fmt"
	"github.com/chaechaep/simulator/config"
	"github.com/chaechaep/simulator/event"
	"github.com/chaechaep/simulator/object"
	"github.com/chaechaep/simulator/types"
	"math/rand"
	"strings"
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
	go func() {
		for {
			user.GetSync()
		}
	}()
	go func() {
		for {
			rand.Seed(time.Now().UnixNano())
			r := rand.Intn(100)
			time.Sleep(time.Duration(r) * time.Second)
			user.SendMessage("m.text", "Msg test"+string(r))
		}
	}()
	go func() {
		for {
			r := rand.Intn(200)
			time.Sleep(time.Duration(r) * time.Second)
			user.ReadMarker()
		}
	}()
}

func checkMemeberCount(user object.User, roomId string) int {
	cnt, _ := user.GetJoinedMembers(roomId)
	return cnt
}
func AdminStart() {
	roomList := config.Cfg.RoomList
	user := object.User{
		UserId:      "chaeuntest",
		AccessToken: "",
		Password:    "ehlswkd123!",
		DeviceId:    "",
		RoomId:      "",
		Sync:        types.SyncResp{},
	}
	user.Login()
	go func() {
		user.GetSync()
	}()
	for _, v := range roomList {
		roomId := ""
		if strings.HasPrefix(v, "#") {
			roomId, _ = event.GetRoomId(v)
		} else {
			roomId = v
		}
		go func() {
			for {
				cnt := checkMemeberCount(user, roomId)
				if cnt > 6 {
					err := user.ChangeJoinRule(roomId, "invite")
					if err != nil {
						fmt.Println(err)
					}
					break
				}
			}
		}()
	}
}
func main() {
	//userList := []string{
	//	"testtest3", "testtest4", "testtest5",
	//}

	confFile := "C://Users/user/GolandProjects/go/src/github.com/chaechaep/simulator/config.json"
	err := config.Init(confFile)
	if err != nil {
		fmt.Println("config load failed : ", err)
		return
	}
	//userList := config.Cfg.UserList
	//go AdminStart()
	//
	//for _, user := range userList {
	//	go Start(user)
	//}
	//
	//fmt.Scanln()

	user := object.User{
		UserId:      "chaeuntest",
		AccessToken: "",
		Password:    "ehlswkd123!",
		DeviceId:    "",
		RoomId:      "",
		Sync:        types.SyncResp{},
	}
	user.Login()
	//user.ChangeJoinRule(config.Cfg.DefaultRoomId, "invite")
	for _, v := range config.Cfg.RoomList {
		roomId := v
		if strings.HasPrefix(v, "#") {
			roomId, err = event.GetRoomId(v)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(roomId)
		}
		cnt, err := user.GetJoinedMembers(roomId)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(cnt)
	}

}
