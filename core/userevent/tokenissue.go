package userevent

import (
	"encoding/hex"
	"encoding/json"
	"github.com/EducationEKT/EKT/core/types"
)

type TokenIssue struct {
	Token     types.Token    `json:"token"`
	From      types.HexBytes `json:"from"`
	Signature types.HexBytes `json:"signature"`
	Nonce     int64          `json:"nonce"`
	success   bool           `json:"-"`
}

func (event TokenIssue) GetNonce() int64 {
	return event.Nonce
}

func (event TokenIssue) Msg() []byte {
	return event.Token.Address()
}

func (event TokenIssue) GetSign() []byte {
	return event.Signature
}

func (event TokenIssue) GetFrom() []byte {
	return event.From
}

func (event TokenIssue) Type() string {
	return TYPE_USEREVENT_PUBLIC_TOKEN
}

func (event TokenIssue) EventId() string {
	data, _ := json.Marshal(event)
	return hex.EncodeToString(data)
}

func (event TokenIssue) SetSuccess() {
	event.success = true
}
