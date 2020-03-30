package object

import (
	"fmt"
	"github.com/chaechaep/simulator/config"
	"github.com/chaechaep/simulator/event"
	"github.com/chaechaep/simulator/types"
)

type User struct {
	UserId      string
	AccessToken string
	Password    string
	DeviceId    string
	RoomId      string
	Sync        types.SyncResp
}

func (user *User) Login() error {
	regAvailableResp, _ := event.GetRegAvailable(user.UserId)

	if regAvailableResp {
		err := user.Register()
		if err != nil {
			return fmt.Errorf("login failed : %s", err)
		}
		return nil
	} else {
		result, err := event.Login(user.UserId, user.Password, user.DeviceId)
		if err != nil {
			return err
		}
		user.AccessToken = result.AccessToken
		user.UserId = result.UserId
		user.DeviceId = result.DeviceId
	}

	joinedRoomList, err := user.GetJoinedRooms()
	if err != nil {
		return err
	}

	if len(joinedRoomList) == 0 {
		err = user.JoinRoom(config.Cfg.DefaultRoomId)
		if err != nil {
			return err
		}
		user.RoomId = config.Cfg.DefaultRoomId
	} else {
		user.RoomId = joinedRoomList[0]
	}
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
	user.Typing()
	err := event.SendMessage(user.AccessToken, user.RoomId, msgType, msg)
	if err != nil {
		return err
	}
	return nil
}

func (user *User) GetSync() error {
	result, err := event.GetSync(user.AccessToken, user.Sync.NextBatch)
	if err != nil {
		return err
	}
	user.Sync = result
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
	if result.RoomId != "" {
		fmt.Println("join room success")
	}
	return nil
}

func (user *User) GetJoinedRooms() (ret []string, err error) {
	result, err := event.GetJoinedRooms(user.AccessToken)
	if err != nil {
		return ret, err
	}
	ret = result.JoinedRooms
	return ret, nil
}

func (user *User) Register() (err error) {
	result, err := event.Register(user.UserId, user.Password, user.DeviceId)
	if err != nil {
		return err
	}
	user.AccessToken = result.AccessToken
	return nil
}

func (user *User) ReadMarker() (err error) {
	if len(user.Sync.Rooms.Join) != 0 {
		for key, _ := range user.Sync.Rooms.Join {
			user.RoomId = key
			break
		}
		if len(user.Sync.Rooms.Join[user.RoomId].Timeline.Events) != 0 {
			eventId := user.Sync.Rooms.Join[user.RoomId].Timeline.Events[0].EventId
			err = event.ReadMarker(user.AccessToken, eventId, user.RoomId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (user *User) Typing() error {
	err := event.Typing(user.AccessToken, user.RoomId, user.UserId)
	if err != nil {
		return err
	}
	return nil
}
