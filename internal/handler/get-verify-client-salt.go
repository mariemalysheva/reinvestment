package handler

import (
	"github.com/mariemalysheva/tokenized-reinvestment/internal/handler/response"
	"github.com/mariemalysheva/tokenized-reinvestment/internal/handler/wrapper"
	"net/http"
)

func (i *Implementation) VerifyClientSalt(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	clientID, err := wrapper.GetIntPathParam(r, wrapper.ClientIDParam)
	if err != nil {
		handleErrResponse(rw, err)
		return
	}

	salt := wrapper.GetStringQueryParam(r, wrapper.ClientSaltParam)

	err = i.adminSvc.VerifyUserSalt(ctx, clientID, salt)
	if err != nil {
		handleErrResponse(rw, err)
		return
	}

	response.OK(rw, nil)
}
