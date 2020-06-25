package object

import (
	"fmt"
	"github.com/chaechaep/simulator/config"
	"github.com/chaechaep/simulator/event"
	"github.com/chaechaep/simulator/log"
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
	regAvailableResp, err := event.GetRegAvailable(user.UserId)
	if err != nil {
		log.Log.Errorf("userId(%s) : %s", user.UserId, err)
	}
	log.Log.Info(user.UserId, "register available : ", regAvailableResp)

	if regAvailableResp {
		err := user.Register()
		if err != nil {
			return fmt.Errorf("login failed : %s", err)
		}
	} else {
		result, err := event.Login(user.UserId, user.Password, user.DeviceId)
		if err != nil {
			return err
		}
		user.AccessToken = result.AccessToken
		user.UserId = result.UserId
		user.DeviceId = result.DeviceId
	}

	fmt.Println("login success : ", user.UserId)
	return nil
}

func (user *User) Logout() error {
	err := event.Logout(user.UserId, user.AccessToken)
	if err != nil {
		return err
	}
	return nil
}

func (user *User) SendMessage(msgType string, msg string) error {
	//user.Typing()
	err := event.SendMessage(user.AccessToken, user.RoomId, user.UserId, msgType, msg)
	if err != nil {
		return err
	}
	return nil
}

func (user *User) GetSync() error {
	result, err := event.GetSync(user.UserId, user.AccessToken, user.Sync.NextBatch, config.Cfg.Simulator.SyncDuration)
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
	result, err := event.JoinRoom(user.UserId, user.AccessToken, roomId)
	if err != nil {
		return err
	}
	user.RoomId = result.RoomId
	if result.RoomId != "" {
		fmt.Println("join room success")
	}
	return nil
}

func (user *User) GetJoinedRooms() (ret []string, err error) {
	result, err := event.GetJoinedRooms(user.UserId, user.AccessToken)
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
	user.UserId = result.UserId
	user.DeviceId = result.DeviceId
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
			err = event.ReadMarker(user.UserId, user.AccessToken, eventId, user.RoomId)
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

func (user *User) ChangeJoinRule(joinRule string) error {
	var joinRules = []string{
		"public",
		"invite",
	}

	if user.RoomId == "" {
		return fmt.Errorf("roomId is not set")
	}
	if !event.Contains(joinRules, joinRule) {
		return fmt.Errorf("invalid join rule")
	}
	err := event.ChangeJoinRule(user.AccessToken, user.RoomId, user.UserId, joinRule)
	if err != nil {
		return err
	}

	return nil
}

func (user *User) GetJoinedMembers(roomId string) (ret int, err error) {
	ret, err = event.GetJoinedMembers(user.UserId, user.AccessToken, roomId)
	if err != nil {
		return 0, err
	}
	return ret, nil
}

func (user *User) CreateRoom() error {
	ret, err := event.CreateRoom(user.UserId, user.AccessToken)
	if err != nil {
		return err
	}
	user.RoomId = ret.RoomId
	return nil
}

func (user *User) GetPublicRooms(nextBatch string) (ret types.GetPublicRoomsResp, err error) {
	ret, err = event.GetPublicRooms(user.UserId, user.AccessToken, nextBatch)
	if err != nil {
		return types.GetPublicRoomsResp{}, err
	}
	return ret, nil
}
