package handler

import (
	"github.com/mariemalysheva/tokenized-reinvestment/internal/handler/response"
	"github.com/mariemalysheva/tokenized-reinvestment/internal/handler/wrapper"
	"net/http"
)

func (i *Implementation) GetReinvestmentSchedule(rw http.ResponseWriter, r *http.Request) {
	_ = r.Context()

	schedule := i.schedulerSvc.GetCurrentReinvestmentSchedule()

	response.OK(rw, wrapper.PostSetScheduleReq{CronSchedule: schedule})
}
