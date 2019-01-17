package common

type Message struct {
	Type     MSG_TYPE
	Code     MSG_CODE
	IsServer bool
	ClientId int
	Message  []byte
}
