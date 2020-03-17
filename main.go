package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	config2 "github.com/CHAEUNPARK/simulator/config"
	"github.com/CHAEUNPARK/simulator/types"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {

	confFile := "C://Users/user/GolandProjects/go/src/github.com/CHAEUNPARK/simulator/config.json"
	config, err := config2.Init(confFile)
	if err != nil {
		fmt.Println("config load failed : ", err)
		return
	}
	//login := Login{}
	//login.Type = "m.login.password"
	//user := User{}
	//user.Type = "m.id.user"
	//user.User = "chaeuntest"
	////login.User = "chaeuntest"
	//login.Identifier = user
	//login.Password = "ehlswkd123!"
	//
	//pbytes, _ := json.Marshal(login)
	//buff := bytes.NewBuffer(pbytes)
	//resp, err = http.Post("http://14.0.81.136:8008/_matrix/client/r0/login", "application/json", buff)
	//if err != nil{
	//	panic(err)
	//}
	//
	//data, err = ioutil.ReadAll(resp.Body)
	//if err != nil{
	//	panic(err)
	//}
	//fmt.Println(string(data))
	/*
			req, err := http.NewRequest("GET", "http://14.0.81.136:8008/_matrix/client/r0/account/whoami", nil)
			if err != nil {
				panic(err)
			}
			req.Header.Add("Authorization", "Bearer MDAxNWxvY2F0aW9uIHBsZWEuaW0KMDAxM2lkZW50aWZpZXIga2V5CjAwMTBjaWQgZ2VuID0gMQowMDI2Y2lkIHVzZXJfaWQgPSBAY2hhZXVudGVzdDpwbGVhLmltCjAwMTZjaWQgdHlwZSA9IGFjY2VzcwowMDIxY2lkIG5vbmNlID0gbmRrcWthY21Oa21yLjpGMQowMDJmc2lnbmF0dXJlIGOC3upFWiChXeMaFRkpOgJq1DEkU74KBRB36okpSgtNCg")
			client := &http.Client{}
			resp, err = client.Do(req)
			if err != nil {
				panic(err)
			}
			data, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(data))
			values := map[string]string{"msgtype": "m.text", "body": "testtest"}
			jsonStr, _ := json.Marshal(values)
			req, err = http.NewRequest("PUT", "http://14.0.81.136:8008/_matrix/client/r0/rooms/!WMxrVkaEfBZBGZOviw%3Aplea.im/send/m.room.message/36", bytes.NewBuffer(jsonStr))
			if err != nil {
				panic(err)
			}
			req.Header.Add("content-type", "application/json")
			req.Header.Add("Authorization", "Bearer MDAxNWxvY2F0aW9uIHBsZWEuaW0KMDAxM2lkZW50aWZpZXIga2V5CjAwMTBjaWQgZ2VuID0gMQowMDI2Y2lkIHVzZXJfaWQgPSBAY2hhZXVudGVzdDpwbGVhLmltCjAwMTZjaWQgdHlwZSA9IGFjY2VzcwowMDIxY2lkIG5vbmNlID0gbmRrcWthY21Oa21yLjpGMQowMDJmc2lnbmF0dXJlIGOC3upFWiChXeMaFRkpOgJq1DEkU74KBRB36okpSgtNCg")
			resp, err = client.Do(req)
			if err != nil {
				panic(err)
			}
			data, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(data))
		// register 부분
		values := types.RegisterReq{
			Auth: types.Auth{
				Type:    "m.login.dummy", // m.login.password 안먹힘
				Session: "",
			},
			UserName:                 "testtest1",
			Password:                 "testtest",
			DeviceId:                 "testtest",
			InitialDeviceDisplayName: "testdevice1",
			InhibitLogin:             false,
		}
		jsonStr, _ := json.Marshal(values)
		req, err := http.NewRequest("POST", "http://14.0.81.136:8008/_matrix/client/r0/register?kind=user", bytes.NewBuffer(jsonStr))
		if err != nil {
			panic(err)
		}
		req.Header.Add("Content-Type", "application/json")
		client := &http.Client{}
		resp, err = client.Do(req)
		if err != nil {
			panic(err)
		}
		if resp.StatusCode == http.StatusOK {
			result := types.RegisterResp{}
			if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
				panic(err)
			} else {
				fmt.Println("success")
				fmt.Println(result)
				fmt.Println("Bearer " + result.AccessToken)
			}
		} else {
			data, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s\n", string(data))
		}
	*/

	//	Login 시작 (Auth 필요없음)
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
	req, err := http.NewRequest("POST", config.BaseUrl+"/login", bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	result := types.LoginResp{}
	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			panic(err)
		} else {
			fmt.Println("success")
			fmt.Println("UserId : ", result.UserId)
			fmt.Println("Access Token : ", result.AccessToken)
			fmt.Println("DeviceId : ", result.DeviceId)
			fmt.Println("Well-known : ", result.WellKnown)
		}
	} else {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", string(data))
	}
	// join room
	// Default Room Id : !WMxrVkaEfBZBGZOviw%3Aplea.im
	/*
		req, err = http.NewRequest("POST", "http://14.0.81.136:8008/_matrix/client/r0/rooms/!WMxrVkaEfBZBGZOviw%3Aplea.im/join", nil)
		if err != nil {
			panic(err)
		}
		req.Header.Add("Authorization", "Bearer "+result.AccessToken)
		resp, err = client.Do(req)
		if err != nil {
			panic(err)
		}
		fmt.Println(resp.StatusCode)
		data, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", string(data))
	*/
	//	joined rooms
	req, err = http.NewRequest("GET", config.BaseUrl+"/joined_rooms", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", config.AccessTokenPrefix+result.AccessToken)
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(data))
	fmt.Println(url.QueryEscape(string(data)))
}