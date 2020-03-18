package types

import "github.com/CHAEUNPARK/simulator/event"

type User struct {
	UserId      string `json:"user_id"`
	AccessToken string `json:"access_token"`
	Password    string `json:"password"`
}

func (user *User) Login() error {
	event.Login(user)

	return nil
}

func (user *User) Logout() error {
	event.Logout(user)
	return nil
}

func (user *User) SendMessage(msg string) error {
	return nil
}

func (user *User) GetSync() error {
	return nil
}

func (user *User) GetDevices() (ret []string, err error) {
	return ret, nil
}

func (user *User) JoinRoom(roomId string) error {
	return nil
}

func (user *User) GetJoinedRooms() (ret []string, err error) {
	return ret, nil
}
