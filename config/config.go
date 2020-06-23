package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Admin struct {
		AdminId       string `json:"admin_id"`
		AdminPassword string `json:"admin_password"`
	} `json:"admin"`
	Simulator struct {
		BaseUrl             string   `json:"base_url"`
		AccessTokenPrefix   string   `json:"access_token_prefix"`
		DefaultRoomId       string   `json:"default_room_id"`
		DefaultPassword     string   `json:"default_password"`
		RoomList            []string `json:"room_list"`
		SendMessageDuration int      `json:"send_message_duration"`
		LoginDuration       int      `json:"login_duration"`
		SyncDuration        int      `json:"sync_duration"`
		RoomMemberCount     int      `json:"room_member_count"`
		CreateUserCount     int      `json:"create_user_count"`
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
