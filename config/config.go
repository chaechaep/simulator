package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Simulator struct {
		BaseUrl             string `json:"base_url"`
		AccessTokenPrefix   string `json:"access_token_prefix"`
		DefaultRoomId       string `json:"default_room_id"`
		DefaultPassword     string `json:"default_password"`
		SendMessageDuration int    `json:"send_message_duration"`
		LoginDuration       int    `json:"login_duration"`
		SyncDuration        int    `json:"sync_duration"`
		CreateUserCount     int    `json:"create_user_count"`
		RoomNamePrefix      string `json:"room_name_prefix"`
		UserNamePrefix      string `json:"user_name_prefix"`
		CreateUserName      string `json:"create_user_name"`
		JoinUserName        string `json:"join_user_name"`
	} `json:"simulator"`
	Log struct {
		LogFile     string `json:"log_file"`
		LogLevel    string `json:"log_level"`
		ProcessName string `json:"process_name"`
	}
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
