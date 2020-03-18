package event

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/CHAEUNPARK/simulator/config"
	"github.com/CHAEUNPARK/simulator/types"
	"net/http"
)

func Register() (ret types.RegisterResp, err error) {
	values := types.RegisterReq{
		Auth: types.AuthenticationData{
			Type:    "m.login.dummy",
			Session: "",
		},
		UserName:                 "testtest1",
		Password:                 "testtest",
		DeviceId:                 "testtest",
		InitialDeviceDisplayName: "testdevice1",
		InhibitLogin:             false,
	}
	jsonStr, _ := json.Marshal(values)
	req, err := http.NewRequest("POST", config.Cfg.BaseUrl+"/register", bytes.NewBuffer(jsonStr))
	if err != nil {
		return ret, err
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ret, err
	}
	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
			return ret, fmt.Errorf("decode failed : %s", err.Error())
		}
	} else {
		errorRet := types.Error{}
		if err := json.NewDecoder(resp.Body).Decode(&errorRet); err != nil {
			return ret, fmt.Errorf("decode failed : %s", err.Error())
		} else {
			return ret, fmt.Errorf("register error( status code : %d, error code : %s, error msg : %s )", resp.StatusCode, errorRet.Errcode, errorRet.Error)
		}
	}
	return ret, nil
}
