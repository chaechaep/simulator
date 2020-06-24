package event

import (
	"encoding/json"
	"fmt"
	"github.com/chaechaep/simulator/config"
	"github.com/chaechaep/simulator/types"
	"math/rand"
	url2 "net/url"
	"time"
)

func JoinRoom(accessToken string, roomId string) (ret types.JoinRoomResp, err error) {
	if roomId == "" {
		roomId = config.Cfg.Simulator.DefaultRoomId
	}
	err = Process("POST", config.Cfg.Simulator.BaseUrl+"/join/"+roomId, nil, &ret, accessToken)
	if err != nil {
		return ret, fmt.Errorf("join room failed : %s", err)
	}
	return ret, nil
}

func GetJoinedRooms(accessToken string) (ret types.JoinedRoomResp, err error) {
	err = Process("GET", config.Cfg.Simulator.BaseUrl+"/joined_rooms", nil, &ret, accessToken)
	if err != nil {
		return ret, fmt.Errorf("get joined rooms failed : %s", err)
	}
	return ret, nil
}

func ReadMarker(accessToken string, eventId string, roomId string) error {
	ret := types.JSONEmpty{}
	if eventId == "" {
		return fmt.Errorf("eventId not set")
	}
	values := types.ReadMarkerReq{
		MFullyRead: eventId,
		MRead:      "",
	}
	jsonStr, _ := json.Marshal(values)
	url := config.Cfg.Simulator.BaseUrl + "/rooms/" + roomId + "/read_markers"
	err := Process("POST", url, jsonStr, &ret, accessToken)
	if err != nil {
		return fmt.Errorf("read marker failed : %s", err)
	}
	return nil
}

func ChangeJoinRule(accessToken string, roomId string, userId string, joinRule string) error {
	values := types.JoinRule{JoinRule: joinRule}

	_, err := SendEvent(accessToken, roomId, "m.room.join_rules", values, userId)
	if err != nil {
		fmt.Errorf("set join rule failed : %s", err)
	}

	return nil
}

func GetRoomId(roomAlias string) (ret string, err error) {
	if roomAlias == "" {
		return ret, fmt.Errorf("room alias not set")
	}
	respValue := types.GetRoomIdResp{}
	url := config.Cfg.Simulator.BaseUrl + "/directory/room/" + url2.QueryEscape(roomAlias)
	err = Process("GET", url, nil, &respValue, "")
	if err != nil {
		return ret, fmt.Errorf("get room id failed : %s", err)
	}
	ret = respValue.RoomId
	return ret, nil
}

func GetJoinedMembers(accessToken string, roomId string) (ret int, err error) {
	resp := types.GetJoinedMembersResp{}
	url := config.Cfg.Simulator.BaseUrl + "/rooms/" + roomId + "/joined_members"
	err = Process("GET", url, nil, &resp, accessToken)
	if err != nil {
		return 0, fmt.Errorf("get joined members failed : %s", err)
	}
	ret = len(resp.Joined)
	return ret, nil
}

func createRandomString() string {
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, 12)
	for i := range b {
		r := rand.Intn(len(charset))
		b[i] = charset[r]
	}
	return string(b)
}

func CreateRoom(accessToken string) (ret types.CreateRoomResp, err error) {
	url := config.Cfg.Simulator.BaseUrl + "/createRoom"
	randomStr := createRandomString()
	values := types.CreateRoomReq{
		Visibility:    "public",
		RoomAliasName: config.Cfg.Simulator.RoomNamePrefix + "-" + randomStr,
		Name:          config.Cfg.Simulator.RoomNamePrefix + "-" + randomStr,
		Topic:         "Matrix Test",
		Preset:        "public_chat",
	}
	jsonStr, _ := json.Marshal(values)
	if err := Process("POST", url, jsonStr, &ret, accessToken); err != nil {
		return ret, fmt.Errorf("create room failed : %s", err)
	}

	return ret, nil
}

func GetPublicRooms(accessToken, nextBatch string) (ret types.GetPublicRoomsResp, err error) {
	url := config.Cfg.Simulator.BaseUrl + "/publicRooms?limit=100"
	if nextBatch != "" {
		url += "&since=" + nextBatch
	}
	if err := Process("GET", url, nil, &ret, accessToken); err != nil {
		return ret, fmt.Errorf("get public rooms failed : %s", err)
	}
	return ret, nil
}
