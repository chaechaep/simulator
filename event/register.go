package event

import (
	"encoding/json"
	"fmt"
	"github.com/chaechaep/simulator/config"
	"github.com/chaechaep/simulator/types"
)

func Register(userName string, password string, deviceId string) (ret types.RegisterResp, err error) {
	values := types.RegisterReq{
		Auth: types.AuthenticationData{
			Type:    "m.login.dummy",
			Session: "",
		},
		UserName:                 userName,
		Password:                 password,
		DeviceId:                 deviceId,
		InitialDeviceDisplayName: "testdevice1",
		InhibitLogin:             false,
	}
	jsonStr, _ := json.Marshal(values)
	err = Process("POST", config.Cfg.Simulator.BaseUrl+"/register", jsonStr, &ret, "")
	if err != nil {
		return ret, fmt.Errorf("regist failed : %s", err)
	}
	return ret, nil
}

func GetRegAvailable(userName string) (bool, error) {
	respValue := types.RegAvailableResp{}
	url := config.Cfg.Simulator.BaseUrl + "/register/available?username=" + userName
	err := Process("GET", url, nil, &respValue, "")
	if err != nil {
		return false, fmt.Errorf("get reg available failed : %s", err)
	}
	return true, nil
}
