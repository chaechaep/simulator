package types

type UserIdentifier struct {
	Type string `json:"type"`
	User string `json:"user"`
}

type LoginReq struct {
	Type                     string         `json:"type"`
	Identifier               UserIdentifier `json:"identifier"`
	Password                 string         `json:"password"`
	DeviceId                 string         `json:"device_id"`
	InitialDeviceDisplayName string         `json:"initial_device_display_name"`
}

type LoginResp struct {
	UserId      string               `json:"user_id"`
	AccessToken string               `json:"access_token"`
	DeviceId    string               `json:"device_id"`
	WellKnown   DiscoveryInformation `json:"well_known"`
}

type DiscoveryInformation struct {
	Homeserver     HomeserverInformation     `json:"m.homeserver"`
	IdentityServer IdentityServerInformation `json:"m.identity_server"`
}

type HomeserverInformation struct {
	BaseUrl string `json:"base_url"`
}

type IdentityServerInformation struct {
	BaseUrl string `json:"base_url"`
}