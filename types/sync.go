package types

type SyncResp struct {
	NextBatch   string      `json:"next_batch"`
	Rooms       Rooms       `json:"rooms"`
	Presence    Presence    `json:"presence"`
	AccountData AccountData `json:"account_data"`
}
