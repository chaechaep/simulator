package types

type JoinRoomResp struct {
	RoomId string `json:"room_id"`
}

type JoinedRoomResp struct {
	JoinedRooms []string `json:"joined_rooms"`
}

type Rooms struct {
	Join   map[string]JoinedRoom `json:"join"`
	Invite map[string]InviteRoom `json:"invite"`
	Leave  map[string]LeftRoom   `json:"leave"`
}

type JoinedRoom struct {
	Summary             RoomSummary              `json:"summary"`
	State               State                    `json:"state"`
	Timeline            Timeline                 `json:"timeline"`
	Ephemeral           Ephemeral                `json:"ephemeral"`
	AccountData         AccountData              `json:"account_data"`
	UnreadNotifications UnreadNotificationCounts `json:"unread_notifications"`
}

type InviteRoom struct {
	InviteState InviteState `json:"invite_state"`
}

type LeftRoom struct {
	State       State       `json:"state"`
	Timeline    Timeline    `json:"timeline"`
	AccountData AccountData `json:"account_data"`
}

type RoomSummary struct {
	MHeros             []string `json:"m.heros"`
	MJoinedMemberCount int      `json:"m.joined_member_count"`
	MInviteMemberCount int      `json:"m.invite_member_count"`
}

type Timeline struct {
	Events    []RoomEvent `json:"events"`
	Limited   bool        `json:"limited"`
	PrevBatch string      `json:"prev_batch"`
}

type UnreadNotificationCounts struct {
	HighlightCount    int `json:"highlight_count"`
	NotificationCount int `json:"notification_count"`
}

type ReadMarkerReq struct {
	MFullyRead string `json:"m.fully_read"`
	MRead      string `json:"m.read"`
}
