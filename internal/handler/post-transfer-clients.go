package handler

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mariemalysheva/tokenized-reinvestment/internal/handler/response"
	"github.com/mariemalysheva/tokenized-reinvestment/internal/handler/wrapper"
	"net/http"
)

func (i *Implementation) PostTransferClients(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req wrapper.PostTransferUsersReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		handleErrResponse(rw, err)
		return
	}

	txHash, err := i.adminSvc.PostTransferUsers(ctx, req.ClientIDs, common.HexToAddress(req.ToAddress))
	if err != nil {
		handleErrResponse(rw, err)
		return
	}

	response.OK(rw, wrapper.RepackTxHashResp(txHash))
}
