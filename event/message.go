package event

import (
	"fmt"
	"github.com/CHAEUNPARK/simulator/types"
)

func SendMessage(accessToken string, roomId string, msgType string, messages ...interface{}) error {
	//Todo:msgType == m.image는 일단 나중에 생각
	var values interface{}
	fmt.Println(msgType)
	values = types.MsgText{
		MsgType: "m.text",
		Body:    messages[0].(string),
	}

	result, err := SendEvent(accessToken, roomId, "m.room.message", values)
	if err != nil {
		return fmt.Errorf("send event failed : %s", err)
	}
	fmt.Println(result)

	return nil
}
