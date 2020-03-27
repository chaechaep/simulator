package types

type SendEventResp struct {
	EventId string `json:"event_id"`
}

type StateEvent struct {
	Content        interface{}  `json:"content"`
	Type           string       `json:"type"`
	EventId        string       `json:"event_id"`
	Sender         string       `json:"sender"`
	OriginServerTs int          `json:"origin_server_ts"`
	Unsigned       UnsignedData `json:"unsigned"`
	PrevContent    EventContent `json:"prev_content"`
	StateKey       string       `json:"state_key"`
}

type UnsignedData struct {
	Age             int    `json:"age"`
	RedactedBecause Event  `json:"redacted_because"`
	TransactionId   string `json:"transaction_id"`
}

type Event struct {
	Content interface{} `json:"content"`
	Type    string      `json:"type"`
}

type EventContent struct {
	AvatarUrl   string       `json:"avatar_url"`
	Displayname string       `json:"displayname"`
	Membership  string       `json:"membership"`
	IsDirect    bool         `json:"is_direct"`
	Unsigned    UnsignedData `json:"unsigned"`
}

type RoomEvent struct {
	Content        interface{}  `json:"content"`
	Type           string       `json:"type"`
	EventId        string       `json:"event_id"`
	Sender         string       `json:"sender"`
	OriginServerTs int          `json:"origin_server_ts"`
	Unsigned       UnsignedData `json:"unsigned"`
}

type Ephemeral struct {
	Events []Event `json:"events"`
}

type AccountData struct {
	Events []Event `json:"events"`
}

type Presence struct {
	Events []Event `json:"events"`
}
