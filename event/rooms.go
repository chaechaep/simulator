package event

import (
	"encoding/json"
	"fmt"
	"github.com/CHAEUNPARK/simulator/config"
	"github.com/CHAEUNPARK/simulator/types"
	"net/http"
)

func JoinRoom() (ret types.JoinRoomResp, err error) {
	req, err := http.NewRequest("POST", config.Cfg.BaseUrl+"/rooms/"+config.Cfg.DefaultRoomId+"/join", nil)
	if err != nil {
		return ret, fmt.Errorf("request create error : %s", err.Error())
	}
	req.Header.Add("Authorizaion", config.Cfg.AccessTokenPrefix)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ret, fmt.Errorf("send request error : %s", err.Error())
	}
	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
			return ret, fmt.Errorf("response decode error : %s", err.Error())
		}
	} else {
		errorRet := types.Error{}
		if err := json.NewDecoder(resp.Body).Decode(&errorRet); err != nil {
			return ret, fmt.Errorf("response decode error : %s", err.Error())
		} else {
			return ret, fmt.Errorf("join room failed ( status code : %d, error code : %s, error msg : %s )", resp.StatusCode, errorRet.Errcode, errorRet.Error)
		}
	}
	return ret, nil
}

func GetJoinedRooms() (ret types.JoinedRoomResp, err error) {
	req, err := http.NewRequest("GET", config.Cfg.BaseUrl+"/joined_rooms", nil)
	if err != nil {
		return ret, fmt.Errorf("request create error : %s", err.Error())
	}
	//Todo : User Type 생성 후 AccessToken 및 UserId 가지고 있기
	req.Header.Add("Authorization", config.Cfg.AccessTokenPrefix)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ret, fmt.Errorf("send request error : %s", err.Error())
	}
	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
			return ret, fmt.Errorf("response decode error : %s", err.Error())
		}
	} else {
		errorRet := types.Error{}
		if err := json.NewDecoder(resp.Body).Decode(&errorRet); err != nil {
			return ret, fmt.Errorf("response decode error : %s", err.Error())
		} else {
			return ret, fmt.Errorf("get joined room falied ( status code : %d, error code : %s, error msg : %s )", resp.StatusCode, errorRet.Errcode, errorRet.Error)
		}
	}
	return ret, nil
}
