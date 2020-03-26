package types

type AuthenticationData struct {
	Type    string `json:"type"`
	Session string `json:"session"`
}

type RegisterReq struct {
	Auth                     AuthenticationData `json:"auth"`
	UserName                 string             `json:"username"`
	Password                 string             `json:"password"`
	DeviceId                 string             `json:"device_id"`
	InitialDeviceDisplayName string             `json:"initial_device_display_name"`
	InhibitLogin             bool               `json:"inhibit_login"`
}

type RegisterResp struct {
	UserId      string `json:"user_id"`
	AccessToken string `json:"access_token"`
	DeviceId    string `json:"device_id"`
}

type RegAvailableResp struct {
	Available bool `json:"available"`
}
