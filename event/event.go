package event

import (
	"encoding/json"
	"fmt"
	"github.com/chaechaep/simulator/config"
	"github.com/chaechaep/simulator/types"
	"math/rand"
	"strings"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var roomEvent = []string{
	"m.room.aliases",
	"m.room.canonical_alias",
	"m.room.create",
	"m.room.join_rules",
	"m.room.member",
	"m.room.power_levels",
	"m.room.redaction",
}

var instantMessagingEvent = []string{
	"m.room.message",
	"m.room.message.feedback",
	"m.room.name",
	"m.room.topic",
	"m.room.avatar",
	"m.room.pinned_events",
}

func createTxnId() string {
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, 12)
	for i := range b {
		r := rand.Intn(len(charset))
		b[i] = charset[r]
	}
	return string(b)
}

func Contains(slice []string, eventType string) bool {
	for _, val := range slice {
		if eventType == val {
			return true
		}
	}
	return false
}
func SendEvent(accessToken string, roomId string, eventType string, reqValue interface{}, userId string) (ret types.SendEventResp, err error) {
	auth := accessToken
	stateKey := ""
	url := ""

	if eventType == "" {
		return ret, fmt.Errorf("eventType is not set")
	}
	if Contains(instantMessagingEvent, eventType) {
		url = config.Cfg.Simulator.BaseUrl + "/rooms/" + roomId + "/send/" + eventType + "/" + createTxnId()
	} else if Contains(roomEvent, eventType) {
		switch eventType {
		case "m.room.aliases":
			stateKey = strings.Split(userId, ":")[1]
		case "m.room.member":
			stateKey = userId
		}
		//state_key -> room event 참고
		url = config.Cfg.Simulator.BaseUrl + "/rooms/" + roomId + "/state/" + eventType + "/" + stateKey
	} else {
		return ret, fmt.Errorf("invalid eventType")
	}
	jsonStr, _ := json.Marshal(reqValue)

	err = Process("PUT", url, jsonStr, &ret, auth, userId)
	if err != nil {
		return ret, fmt.Errorf("send event failed : %s", err)
	}
	return ret, nil
}
