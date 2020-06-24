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

type GetRoomIdResp struct {
	RoomId  string   `json:"room_id"`
	Servers []string `json:"servers"`
}

type GetJoinedMembersResp struct {
	Joined map[string]RoomMember `json:"joined"`
}

type RoomMember struct {
	DisplayName string `json:"display_name"`
	AvatarUrl   string `json:"avatar_url"`
}

type CreateRoomReq struct {
	Visibility    string `json:"visibility"` //["public", "private"]
	RoomAliasName string `json:"room_alias_name"`
	Name          string `json:"name"`
	Topic         string `json:"topic"`
	Preset        string `json:"preset"` //["private_chat", "public_chat", "trusted_private_chat"]
}

type CreateRoomResp struct {
	RoomId    string `json:"room_id"`
	RoomAlias string `json:"room_alias"`
}

type GetPublicRoomsResp struct {
	Chunk                  []PublicRoomsChunk `json:"chunk"`
	NextBatch              string             `json:"next_batch"`
	PrevBatch              string             `json:"prev_batch"`
	TotalRoomCountEstimate int                `json:"total_room_count_estimate"`
}

type PublicRoomsChunk struct {
	Aliases          []string `json:"aliases"`
	CanonicalAlias   string   `json:"canonical_alias"`
	Name             string   `json:"name"`
	NumJoinedMembers int      `json:"num_joined_members"`
	RoomId           string   `json:"room_id"`
	Topic            string   `json:"topic"`
	WorldReadable    bool     `json:"world_readable"`
	GuestCanJoin     bool     `json:"guest_can_join"`
	AvatarUrl        string   `json:"avatar_url"`
}
