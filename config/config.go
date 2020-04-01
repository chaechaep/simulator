package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	BaseUrl           string   `json:"base_url"`
	AccessTokenPrefix string   `json:"access_token_prefix"`
	DefaultRoomId     string   `json:"default_room_id"`
	DefaultPassword   string   `json:"default_password"`
	RoomList          []string `json:"room_list"`
	UserList          []string `json:"user_list"`
}

var Cfg *Config

func Init(fileName string) error {
	fs, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Load Config Failed : ", err.Error())
		return err
	}
	defer fs.Close()

	js := json.NewDecoder(fs)
	err = js.Decode(&Cfg)
	if err != nil {
		fmt.Println("config json decoding failed : ", err.Error())
		return err
	}

	return nil
}
