package event

import (
	"encoding/json"
	"fmt"
	"github.com/chaechaep/simulator/config"
	"github.com/chaechaep/simulator/types"
)

func Login(userId string, password string, deviceId string) (ret types.LoginResp, err error) {
	auth := ""
	values := types.LoginReq{
		Type: "m.login.password",
		Identifier: types.UserIdentifier{
			Type: "m.id.user",
			User: userId,
		},
		Password:                 password,
		DeviceId:                 deviceId,
		InitialDeviceDisplayName: "",
	}
	jsonStr, _ := json.Marshal(values)
	err = Process("POST", config.Cfg.Simulator.BaseUrl+"/login", jsonStr, &ret, auth)
	if err != nil {
		return ret, fmt.Errorf("login failed : %s", err)
	}

	return ret, nil
}

func Logout(accessToken string) error {
	resp := types.JSONEmpty{}
	err := Process("POST", config.Cfg.Simulator.BaseUrl+"/logout", nil, &resp, accessToken)
	if err != nil {
		return fmt.Errorf("logout failed : %s", err)
	}
	return nil
}
