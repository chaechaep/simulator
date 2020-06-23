package event

import (
	"encoding/json"
	"fmt"
	"github.com/chaechaep/simulator/config"
	"github.com/chaechaep/simulator/types"
)

func SendMessage(accessToken string, roomId string, userId string, msgType string, messages ...interface{}) error {
	//Todo:msgType == m.image는 일단 나중에 생각
	var values interface{}
	fmt.Println(msgType)
	values = types.MsgText{
		MsgType: "m.text",
		Body:    messages[0].(string),
	}

	result, err := SendEvent(accessToken, roomId, "m.room.message", values, userId)
	if err != nil {
		return fmt.Errorf("send event failed : %s", err)
	}
	fmt.Println(result)

	return nil
}

func Typing(accessToken string, roomId string, userId string) error {
	values := types.TypingReq{
		Typing:  true,
		Timeout: 3000,
	}
	jsonStr, _ := json.Marshal(values)
	url := config.Cfg.Simulator.BaseUrl + "/rooms/" + roomId + "/typing/" + userId
	err := Process("PUT", url, jsonStr, &types.JSONEmpty{}, accessToken)
	if err != nil {
		return fmt.Errorf("typing failed : %s", err)
	}
	return nil
}
