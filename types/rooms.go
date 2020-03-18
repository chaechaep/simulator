package types

type JoinRoomResp struct {
	RoomId string `json:"room_id"`
}

type JoinedRoomResp struct {
	JoinedRoom []string `json:"joined_room"`
}
