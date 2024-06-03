package handler

import (
	"encoding/json"
	"github.com/mariemalysheva/tokenized-reinvestment/internal/handler/response"
	"github.com/mariemalysheva/tokenized-reinvestment/internal/handler/wrapper"
	"net/http"
)

func (i *Implementation) PostRescheduleReinvestment(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req wrapper.PostSetScheduleReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		handleErrResponse(rw, err)
		return
	}

	err = i.schedulerSvc.RescheduleCronHandlerFunc(ctx, req.CronSchedule, i.schedulerSvc.ReinvestSavings)
	if err != nil {
		handleErrResponse(rw, err)
		return
	}

	response.OK(rw, nil)
}
