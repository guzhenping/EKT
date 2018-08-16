package api

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/EducationEKT/EKT/blockchain_manager"
	"github.com/EducationEKT/EKT/conf"
	"github.com/EducationEKT/EKT/core/userevent"
	"github.com/EducationEKT/EKT/param"
	"github.com/EducationEKT/EKT/util"
	"github.com/EducationEKT/xserver/x_err"
	"github.com/EducationEKT/xserver/x_http/x_req"
	"github.com/EducationEKT/xserver/x_http/x_resp"
	"github.com/EducationEKT/xserver/x_http/x_router"
	"strings"
)

func init() {
	x_router.All("/token/api/issue", broadcastTokenIssue, issueToken)
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

func broadcastTokenIssue(req *x_req.XReq) (*x_resp.XRespContainer, *x_err.XErr) {
	IP := strings.Split(req.R.RemoteAddr, ":")[0]
	broadcasted := false
	for _, peer := range param.MainChainDelegateNode {
		if peer.Address == IP {
			broadcasted = true
			break
		}
	}
	if !broadcasted {
		for _, peer := range param.MainChainDelegateNode {
			if !peer.Equal(conf.EKTConfig.Node) {
				url := fmt.Sprintf(`http://%s:%d/token/api/issue`, peer.Address, peer.Port)
				util.HttpPost(url, req.Body)
			}
		}
	}
	return nil, nil
}
