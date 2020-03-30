package types

type MsgText struct {
	MsgType string `json:"msgtype"`
	Body    string `json:"body"`
}

//type MsgImage struct {
//	MsgType string `json:"msgtype"`
//	Info ImageInfo
//}
//
//type ImageInfo struct {
//
//}

type TypingReq struct {
	Typing  bool `json:"typing"`
	Timeout int  `json:"timeout"`
}
