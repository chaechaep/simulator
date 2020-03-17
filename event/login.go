package event

import (
	"bytes"
	"encoding/json"
	"fmt"
	config2 "github.com/CHAEUNPARK/simulator/config"
	"github.com/CHAEUNPARK/simulator/types"
	"net/http"
)

func Login() (ret types.LoginResp, err error) {
	values := types.LoginReq{
		Type: "m.login.password",
		Identifier: types.UserIdentifier{
			Type: "m.id.user",
			User: "testtest1",
		},
		Password:                 "testtest",
		DeviceId:                 "testtest",
		InitialDeviceDisplayName: "",
	}
	jsonStr, _ := json.Marshal(values)
	req, err := http.NewRequest("POST", config2.Cfg.BaseUrl+"/login", bytes.NewBuffer(jsonStr))
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
			return ret, err
		}
	} else {
		return ret, fmt.Errorf("BadRequest")
	}
	return ret, nil
}

func Logout(accessToken string) error {

	return nil
}
