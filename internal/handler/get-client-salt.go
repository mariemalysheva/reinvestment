package handler

import (
	"github.com/mariemalysheva/tokenized-reinvestment/internal/handler/response"
	"github.com/mariemalysheva/tokenized-reinvestment/internal/handler/wrapper"
	"net/http"
)

func (i *Implementation) GetClientSalt(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	clientID, err := wrapper.GetIntPathParam(r, wrapper.ClientIDParam)
	if err != nil {
		handleErrResponse(rw, err)
		return
	}

	client, err := i.adminSvc.GetUserSalt(ctx, clientID)
	if err != nil {
		handleErrResponse(rw, err)
		return
	}

	response.OK(rw, client)
}
