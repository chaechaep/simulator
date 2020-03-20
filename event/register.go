package event

import (
	"encoding/json"
	"fmt"
	"github.com/CHAEUNPARK/simulator/config"
	"github.com/CHAEUNPARK/simulator/types"
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
	err = Process("POST", config.Cfg.BaseUrl+"/register", jsonStr, &ret, "")
	if err != nil {
		return ret, fmt.Errorf("regist failed : %s", err)
	}
	return ret, nil
}
