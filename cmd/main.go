package main

import (
	"flag"
	"fmt"
	"github.com/chaechaep/simulator/config"
	"github.com/chaechaep/simulator/event"
	"github.com/chaechaep/simulator/log"
	"github.com/chaechaep/simulator/object"
	"github.com/chaechaep/simulator/types"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func Start(userId string) {
	var err error
	user := object.User{
		UserId:      userId,
		AccessToken: "",
		Password:    config.Cfg.Simulator.DefaultPassword,
		DeviceId:    "",
	}

	if err = user.Login(); err != nil {
		log.Log.Errorf("userId(%s) : %s", user.UserId, err)
	} else {
		log.Log.Info(userId, "login success")
	}

	go func() {
		for {
			if err = user.GetSync(); err != nil {
				log.Log.Errorf("userId(%s) : %s", user.UserId, err)
			}
		}
	}()
	go func() {
		for {
			r := rand.Intn(200)
			time.Sleep(time.Duration(config.Cfg.Simulator.SendMessageDuration) * time.Second)
			if err = user.SendMessage("m.text", "Msg test"+string(r)); err != nil {
				log.Log.Errorf("userId(%s) : %s", user.UserId, err)
			}
		}
	}()
	go func() {
		for {
			r := rand.Intn(100)
			time.Sleep(time.Duration(r) * time.Second)
			if err = user.ReadMarker(); err != nil {
				log.Log.Errorf("userId(%s) : %s", user.UserId, err)
			}
		}
	}()
}

func checkMemberCount(user object.User, roomId string) int {
	cnt, err := user.GetJoinedMembers(roomId)
	if err != nil {
		log.Log.Errorf("userId(%s) : %s", user.UserId, err)
	}
	return cnt
}
func AdminStart() {
	var err error
	roomList := config.Cfg.Simulator.RoomList
	user := object.User{
		UserId:      "chaeuntest",
		AccessToken: "",
		Password:    "ehlswkd123!",
		DeviceId:    "",
		RoomId:      "",
		Sync:        types.SyncResp{},
	}
	if err = user.Login(); err != nil {
		log.Log.Errorf("userId(%s) : %s", user.UserId, err)
	}
	userCnt := 0

	go func() {
		if err = user.GetSync(); err != nil {
			log.Log.Errorf("userId(%s) : %s", user.UserId, err)
		}

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
				time.Sleep(time.Second)
				cnt := checkMemberCount(user, roomId)
				if cnt > config.Cfg.Simulator.RoomMemberCount {
					userCnt += cnt - 1
					err := user.ChangeJoinRule(roomId, "invite")
					if err != nil {
						log.Log.Errorf("userId(%s) : %s", user.UserId, err)
					}
					break
				}
				if userCnt >= config.Cfg.Simulator.CreateUserCount {
					break
				}
			}
		}()
	}
}

var (
	configFile = flag.String("c", "config.json", "The path to the config file.")
)

func main() {
	flag.Parse()

	fmt.Println("program: ", os.Args[0])
	fmt.Println("config: ", *configFile)
	err := config.Init(*configFile)
	if err != nil {
		fmt.Println("config load failed : ", err)
		return
	}

	ProcessName := os.Args[0]
	if config.Cfg.Log.LogFile != "" {
		ProcessName = config.Cfg.Log.ProcessName
	}

	log.Init(config.Cfg.Log.LogFile, config.Cfg.Log.LogLevel, ProcessName)

	go AdminStart()
	for i := 0; i < config.Cfg.Simulator.CreateUserCount; i++ {
		time.Sleep(time.Duration(config.Cfg.Simulator.LoginDuration) * time.Second)
		go Start("testUser" + strconv.Itoa(i))
	}
	fmt.Scanln()
}
