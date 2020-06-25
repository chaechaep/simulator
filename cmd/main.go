package main

import (
	"flag"
	"fmt"
	"github.com/chaechaep/simulator/config"
	"github.com/chaechaep/simulator/log"
	"github.com/chaechaep/simulator/object"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func errorLog(user object.User, err error) {
	log.Log.Errorf("userId(%s) : %s", user.UserId, err)
}

func Start(userId string) {
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
		joinedRoomList, err := user.GetJoinedRooms()
		if err != nil {
			errorLog(user, err)
		}
		if len(joinedRoomList) == 0 {
			nextBatch := ""
			for user.RoomId == "" {
				if rooms, err := user.GetPublicRooms(nextBatch); err != nil {
					errorLog(user, err)
				} else {
					for _, room := range rooms.Chunk {
						if room.NumJoinedMembers < 2 {
							if err := user.JoinRoom(room.RoomId); err != nil {
								errorLog(user, err)
							}
							break
						}
					}
					if user.RoomId == "" {
						if rooms.NextBatch == "" {
							if rooms.TotalRoomCountEstimate < config.Cfg.Simulator.CreateUserCount/2 {
								for {
									time.Sleep(5 * time.Second)
									if err := user.CreateRoom(); err != nil {
										errorLog(user, err)
									} else {
										break
									}
								}
							} else {
								break
							}
						} else {
							nextBatch = rooms.NextBatch
						}
					} else {
						break
					}
				}
			}
		} else {
			user.RoomId = joinedRoomList[0]
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
					time.Sleep(time.Duration(config.Cfg.Simulator.SendMessageDuration) * time.Second)
					if err = user.SendMessage("m.text", "Msg test"+string(r)); err != nil {
						errorLog(user, err)
					}
				}
			}()
			go func() {
				for {
					r := rand.Intn(100)
					time.Sleep(time.Duration(r) * time.Second)
					if err = user.ReadMarker(); err != nil {
						errorLog(user, err)
					}
				}
			}()
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

	//go AdminStart()
	for i := 0; i < config.Cfg.Simulator.CreateUserCount; i++ {
		time.Sleep(time.Duration(config.Cfg.Simulator.LoginDuration) * time.Second)
		go Start(config.Cfg.Simulator.UserNamePrefix + strconv.Itoa(i))
	}
	fmt.Scanln()
}
