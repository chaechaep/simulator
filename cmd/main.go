package main

import (
	"flag"
	"fmt"
	"github.com/chaechaep/simulator/config"
	"github.com/chaechaep/simulator/event"
	"github.com/chaechaep/simulator/log"
	"github.com/chaechaep/simulator/object"
	"math/rand"
	"os"
	"time"
)

func errorLog(user object.User, err error) {
	log.Log.Errorf("userId(%s) : %s", user.UserId, err)
}

func Start(userId, roomName string) {
	var roomId string
	var err error
	user := object.User{
		UserId:      userId,
		AccessToken: "",
		Password:    config.Cfg.Simulator.DefaultPassword,
		DeviceId:    "",
	}

	if err = user.Login(); err != nil {
		errorLog(user, err)
	} else {
		log.Log.Info(userId, " : login success")
		if roomId, err = event.GetRoomId(user.UserId, "#"+roomName+":plea.im"); err != nil {
			errorLog(user, err)
		} else {
			user.RoomId = roomId
		}
		if user.RoomId != "" {
			go func() {
				for {
					if err = user.GetSync(); err != nil {
						errorLog(user, err)
					}
				}
			}()
			go func() {
				for {
					r := rand.Intn(200)
					time.Sleep(time.Duration(1/config.Cfg.Simulator.TPS) * time.Second)
					if err = user.SendMessage("m.text", "Msg test"+string(r)); err != nil {
						errorLog(user, err)
					}
				}
			}()
			/*go func() {
				for {
					r := rand.Intn(100)
					time.Sleep(time.Duration(r) * time.Second)
					if err = user.ReadMarker(); err != nil {
						errorLog(user, err)
					}
				}
			}()*/
		}
	}
}

func createRoom(userId, roomName string) {
	var err error
	user := object.User{
		UserId:   userId,
		Password: config.Cfg.Simulator.DefaultPassword,
	}
	if err = user.Login(); err != nil {
		errorLog(user, err)
	} else {
		if err = user.CreateRoom(roomName); err != nil {
			errorLog(user, err)
		} else {
			fmt.Println("create room success : ", userId, " / ", roomName)
		}
	}
}

func joinRoom(userId, roomName string) {
	var err error
	var roomId string
	user := object.User{
		UserId:   userId,
		Password: config.Cfg.Simulator.DefaultPassword,
	}
	if err = user.Login(); err != nil {
		errorLog(user, err)
	} else {
		if roomId, err = event.GetRoomId(user.UserId, "#"+roomName+":plea.im"); err != nil {
			errorLog(user, err)
		} else {
			if err = user.JoinRoom(roomId); err != nil {
				errorLog(user, err)
			} else {
				fmt.Println("join room success : ", userId, " / ", roomName)
			}
		}
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

	for i := 0; i < config.Cfg.Simulator.SessionCount/2; i++ {
		time.Sleep(time.Duration(config.Cfg.Simulator.LoginDuration) * time.Second)
		go Start(config.Cfg.Simulator.CreateUserName+fmt.Sprintf("%04d", i), config.Cfg.Simulator.RoomNamePrefix+fmt.Sprintf("%04d", i))
	}

	for i := 0; i < config.Cfg.Simulator.SessionCount/2; i++ {
		time.Sleep(time.Duration(config.Cfg.Simulator.LoginDuration) * time.Second)
		go Start(config.Cfg.Simulator.JoinUserName+fmt.Sprintf("%04d", i), config.Cfg.Simulator.RoomNamePrefix+fmt.Sprintf("%04d", i))
	}

	fmt.Scanln()
}
