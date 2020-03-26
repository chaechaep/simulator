package event

import (
	"fmt"
	"github.com/chaechaep/simulator/config"
	"github.com/chaechaep/simulator/types"
)

func JoinRoom(accessToken string, roomId string) (ret types.JoinRoomResp, err error) {
	if roomId == "" {
		roomId = config.Cfg.DefaultRoomId
	}
	err = Process("POST", config.Cfg.BaseUrl+"/rooms/"+roomId+"/join", nil, &ret, accessToken)
	if err != nil {
		return ret, fmt.Errorf("join room failed : %s", err)
	}
	return ret, nil
}

func GetJoinedRooms(accessToken string) (ret types.JoinedRoomResp, err error) {
	err = Process("GET", config.Cfg.BaseUrl+"/joined_rooms", nil, &ret, accessToken)
	if err != nil {
		return ret, fmt.Errorf("get joined rooms failed : %s", err)
	}
	return ret, nil
}
