package event

import (
	"encoding/json"
	"fmt"
	"github.com/CHAEUNPARK/simulator/config"
	"github.com/CHAEUNPARK/simulator/types"
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
	resp := types.LoginResp{}
	err = Process("POST", config.Cfg.BaseUrl+"/login", jsonStr, &resp, auth)
	if err != nil {
		return ret, fmt.Errorf("login failed : %s", err)
	}

	ret = resp

	return ret, nil
}

func Logout(accessToken string) error {
	resp := types.JSONEmpty{}
	err := Process("POST", config.Cfg.BaseUrl+"/logout", nil, &resp, accessToken)
	if err != nil {
		return fmt.Errorf("logout failed : %s", err)
	}
	fmt.Println(resp)
	return nil
}
