package event

import (
	"encoding/json"
	"fmt"
	"github.com/CHAEUNPARK/simulator/config"
	"github.com/CHAEUNPARK/simulator/types"
)

func Login(userId string, password string) (ret map[string]interface{}, err error) {
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

	result, err := Process("POST", config.Cfg.BaseUrl+"/login", jsonStr, types.LoginResp{}, auth)
	if err != nil {
		return ret, fmt.Errorf("login failed : %s", err)
	}

	ret = result

	return ret, nil
}

func Logout(accessToken string) error {
	result, err := Process("POST", config.Cfg.BaseUrl+"/logout", nil, types.JSONEmpty{}, accessToken)
	if err != nil {
		return fmt.Errorf("logout failed : %s", err)
	}
	fmt.Println(result)
	return nil
}
