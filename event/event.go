package event

import (
	"encoding/json"
	"fmt"
	"github.com/chaechaep/simulator/config"
	"github.com/chaechaep/simulator/types"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func createTxnId() string {
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, 12)
	for i := range b {
		r := rand.Intn(len(charset))
		b[i] = charset[r]
	}
	return string(b)
}

func SendEvent(accessToken string, roomId string, eventType string, reqValue interface{}) (ret types.SendEventResp, err error) {
	auth := accessToken
	stateKey := ""
	url := ""
	if eventType == "" {
		return ret, fmt.Errorf("eventType is not set")
	}
	if eventType == "m.room.message" {

		url = config.Cfg.BaseUrl + "/rooms/" + roomId + "/send/" + eventType + "/" + createTxnId()
	} else {
		//state_key -> room event 참고
		url = config.Cfg.BaseUrl + "/rooms/" + roomId + "/state/" + eventType + stateKey
	}
	jsonStr, _ := json.Marshal(reqValue)

	err = Process("PUT", url, jsonStr, &ret, auth)
	if err != nil {
		return ret, fmt.Errorf("send event failed : %s", err)
	}
	return ret, nil
}
