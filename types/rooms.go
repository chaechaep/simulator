package types

type JoinRoomResp struct {
	RoomId string `json:"room_id"`
}

type JoinedRoomResp struct {
	JoinedRooms []string `json:"joined_rooms"`
}
