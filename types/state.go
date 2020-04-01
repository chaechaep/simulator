package types

type State struct {
	Events []StateEvent `json:"events"`
}

type InviteState struct {
	Events []StrippedState `json:"events"`
}

type StrippedState struct {
	Content  EventContent `json:"content"`
	StateKey string       `json:"state_key"`
	Type     string       `json:"type"`
	Sender   string       `json:"sender"`
}

type JoinRule struct {
	JoinRule string `json:"join_rule"`
}
