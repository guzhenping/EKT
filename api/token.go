package api

import (
	"encoding/hex"
	"encoding/json"
	"github.com/EducationEKT/EKT/blockchain_manager"
	"github.com/EducationEKT/EKT/core/userevent"
	"github.com/EducationEKT/xserver/x_err"
	"github.com/EducationEKT/xserver/x_http/x_req"
	"github.com/EducationEKT/xserver/x_http/x_resp"
	"github.com/EducationEKT/xserver/x_http/x_router"
)

func init() {
	x_router.All("/token/api/issue", issueToken)
}

func issueToken(req *x_req.XReq) (*x_resp.XRespContainer, *x_err.XErr) {
	var tokenIssue userevent.TokenIssue
	err := json.Unmarshal(req.Body, &tokenIssue)
	if err != nil {
		return x_resp.Fail(-1, "error request", nil), nil
	}
	if userevent.Validate(&tokenIssue) {
		blockchain_manager.GetMainChain().NewUserEvent(&tokenIssue)
		return x_resp.Success(hex.EncodeToString(tokenIssue.Token.Address())), nil
	}
	return x_resp.Fail(-1, "error signature", nil), nil
}
