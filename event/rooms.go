package event

import (
	"encoding/json"
	"fmt"
	"github.com/CHAEUNPARK/simulator/config"
	"github.com/CHAEUNPARK/simulator/types"
	"net/http"
)

func JoinRoom(accessToken string) (ret types.JoinRoomResp, err error) {
	req, err := http.NewRequest("POST", config.Cfg.BaseUrl+"/rooms/"+config.Cfg.DefaultRoomId+"/join", nil)
	if err != nil {
		return ret, fmt.Errorf("request create error : %s", err.Error())
	}
	req.Header.Add("Authorizaion", config.Cfg.AccessTokenPrefix+accessToken)
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

func GetJoinedRooms(accessToken string) (ret map[string]interface{}, err error) {
	result, err := Process("GET", config.Cfg.BaseUrl+"/joined_rooms", nil, types.JoinedRoomResp{}, accessToken)
	if err != nil {
		return ret, fmt.Errorf("get joined rooms failed : %s", err)
	}
	ret = result
	/*
		req, err := http.NewRequest("GET", config.Cfg.BaseUrl+"/joined_rooms", nil)
		if err != nil {
			return ret, fmt.Errorf("request create error : %s", err.Error())
		}
		req.Header.Add("Authorization", config.Cfg.AccessTokenPrefix+accessToken)
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

	*/
	return ret, nil
}
