package handler

import (
	"github.com/mariemalysheva/tokenized-reinvestment/internal/handler/response"
	"net/http"
)

func (i *Implementation) GetClients(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := i.adminSvc.GetUsers(ctx)
	if err != nil {
		handleErrResponse(rw, err)
		return
	}

	response.OK(rw, users)
}
