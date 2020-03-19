package event

import (
	"encoding/json"
	"fmt"
	"github.com/CHAEUNPARK/simulator/config"
	"github.com/CHAEUNPARK/simulator/types"
	"net/http"
)

func Login(userId string, password string) (ret types.LoginResp, err error) {
	auth := ""
	values := types.LoginReq{
		Type: "m.login.password",
		Identifier: types.UserIdentifier{
			Type: "m.id.user",
			User: userId,
		},
		Password:                 password,
		DeviceId:                 "testtest",
		InitialDeviceDisplayName: "",
	}
	jsonStr, _ := json.Marshal(values)
	/*
		req, err := http.NewRequest("POST", config.Cfg.BaseUrl+"/login", bytes.NewBuffer(jsonStr))
		if err != nil {
			return ret, err
		}
		req.Header.Add("Content-Type", "application/json")
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
				return ret, fmt.Errorf("login failed ( status code : %d, error code : %s, error msg : %s )", resp.StatusCode, errorRet.Errcode, errorRet.Error)
			}
		}
	*/

	result, err := Process("POST", config.Cfg.BaseUrl+"/login", jsonStr, types.LoginResp{}, auth)
	if err != nil {
		return ret, err
	}

	ret.UserId = result["user_id"].(string)
	ret.AccessToken = result["access_token"].(string)
	ret.DeviceId = result["device_id"].(string)

	return ret, nil
}

func Logout(accessToken string) error {
	req, err := http.NewRequest("POST", config.Cfg.BaseUrl+"/logout", nil)
	if err != nil {
		return fmt.Errorf("request create error : %s", err.Error())
	}

	req.Header.Add("Authorization", config.Cfg.AccessTokenPrefix+accessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("send request error : %s", err.Error())
	}

	if resp.StatusCode == http.StatusOK {
		return nil
	} else {
		errorRet := types.Error{}
		if json.NewDecoder(resp.Body).Decode(&errorRet); err != nil {
			return fmt.Errorf("response decode error : %s", err.Error())
		} else {
			return fmt.Errorf("logout falied ( status code : %d, error code : %s, error msg : %s )", resp.StatusCode, errorRet.Errcode, errorRet.Error)
		}
	}
	return nil
}
