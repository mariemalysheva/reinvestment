package handler

import (
	"github.com/mariemalysheva/tokenized-reinvestment/internal/handler/response"
	"github.com/mariemalysheva/tokenized-reinvestment/internal/handler/wrapper"
	"net/http"
)

func (i *Implementation) GetReinvestmentPeriods(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	records, err := i.adminSvc.GetReinvestmentPeriods(ctx)
	if err != nil {
		handleErrResponse(rw, err)
		return
	}

	response.OK(rw, wrapper.RepackReinvestmentPeriods(records))
}
