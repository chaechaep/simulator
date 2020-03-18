package main

import (
	"fmt"
	config2 "github.com/CHAEUNPARK/simulator/config"
	"github.com/CHAEUNPARK/simulator/event"
	"github.com/CHAEUNPARK/simulator/types"
)

func main() {

	confFile := "C://Users/user/GolandProjects/go/src/github.com/CHAEUNPARK/simulator/config.json"
	err := config2.Init(confFile)
	if err != nil {
		fmt.Println("config load failed : ", err)
		return
	}

	/*
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
	*/
	login, err := event.Login()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(login.UserId)
	fmt.Println(login.AccessToken)
	user := types.User{
		UserId:      login.UserId,
		AccessToken: login.AccessToken,
		AvatarUrl:   "",
		DisplayName: "",
	}

}
