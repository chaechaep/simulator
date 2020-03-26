package object

import (
	"fmt"
	"github.com/chaechaep/simulator/config"
	"github.com/chaechaep/simulator/event"
)

type User struct {
	UserId      string `json:"user_id"`
	AccessToken string `json:"access_token"`
	Password    string `json:"password"`
	DeviceId    string `json:"device_id"`
}

func (user *User) Login() error {
	result, err := event.Login(user.UserId, user.Password, user.DeviceId)
	if err != nil {
		return err
	}
	user.AccessToken = result.AccessToken
	fmt.Println(result)
	return nil
}

func (user *User) Logout() error {
	err := event.Logout(user.AccessToken)
	if err != nil {
		return err
	}
	return nil
}

func (user *User) SendMessage(msgType string, msg string) error {
	joinedRoomList, err := user.GetJoinedRooms()
	if err != nil {
		return err
	}
	roomId := ""
	if len(joinedRoomList) == 0 {
		user.JoinRoom(config.Cfg.DefaultRoomId)
	} else {
		roomId = joinedRoomList[0]
	}

	err = event.SendMessage(user.AccessToken, roomId, msgType, msg)
	if err != nil {
		return err
	}
	return nil
}

func (user *User) GetSync() error {
	return nil
}

func (user *User) GetDevices() (ret []string, err error) {
	return ret, nil
}

func (user *User) JoinRoom(roomId string) error {
	result, err := event.JoinRoom(user.AccessToken, roomId)
	if err != nil {
		return err
	}
	fmt.Println(result)

	return nil
}

func (user *User) GetJoinedRooms() (ret []string, err error) {
	result, err := event.GetJoinedRooms(user.AccessToken)
	if err != nil {
		return ret, err
	}
	fmt.Println(result)
	ret = result.JoinedRooms
	return ret, nil
}

func Register(userId string, password string, deviceId string) (user *User, err error) {
	result, err := event.Register(userId, password, deviceId)
	if err != nil {
		return user, err
	}
	fmt.Println(result)
	user = &User{
		UserId:      userId,
		AccessToken: result.AccessToken,
		Password:    password,
		DeviceId:    deviceId,
	}
	err = user.Login()
	if err != nil {
		return user, err
	}

	return user, nil
}
