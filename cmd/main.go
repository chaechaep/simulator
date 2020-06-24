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
								if err := user.CreateRoom(); err != nil {
									time.Sleep(5 * time.Second)
									errorLog(user, err)
								} else {
									go func() {
										for {
											time.Sleep(time.Duration(config.Cfg.Simulator.LoginDuration) * time.Second)
											cnt, err := checkMemberCount(user)
											if err != nil {
												errorLog(user, err)
											}
											if cnt >= 2 {
												err := user.ChangeJoinRule("invite")
												if err != nil {
													errorLog(user, err)
												} else {
													break
												}
											}
										}
									}()
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

func checkMemberCount(user object.User) (int, error) {
	cnt, err := user.GetJoinedMembers(user.RoomId)
	if err != nil {
		return 0, err
	}
	return cnt, nil
}

/*
func AdminStart() {
	var err error
	roomList := config.Cfg.Simulator.RoomList
	user := object.User{
		UserId:      config.Cfg.Admin.AdminId,
		AccessToken: "",
		Password:    config.Cfg.Admin.AdminPassword,
		DeviceId:    "",
		RoomId:      "",
		Sync:        types.SyncResp{},
	}
	if err = user.Login(); err != nil {
		log.Log.Errorf("userId(%s) : %s", user.UserId, err)
	} else {
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
					cnt, err := checkMemberCount(user, roomId)
					if err != nil {
						log.Log.Errorf("userId(%s) : %s", user.UserId, err)
						break
					}
					if cnt > config.Cfg.Simulator.RoomMemberCount {
						userCnt += cnt - 1
						err := user.ChangeJoinRule(roomId, "invite")
						if err != nil {
							log.Log.Errorf("userId(%s) : %s", user.UserId, err)
							break
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
}

*/

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
