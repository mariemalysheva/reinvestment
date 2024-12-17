package handler

import (
	"errors"
	"github.com/mariemalysheva/tokenized-reinvestment/internal/handler/response"
	"github.com/opentracing/opentracing-go/log"
	"net/http"
)

var (
	ErrInvalidRequest = errors.New("failed to unmarshal request body")
	ErrInvalidAddress = errors.New("invalid address")
)

func handleErrResponse(rw http.ResponseWriter, err error) {
	defer log.Error(err)
	switch err {
	case
		ErrInvalidRequest,
		ErrInvalidAddress:
		response.BadRequest(rw, err.Error())
	default:
		response.Internal(rw, err.Error())
	}
}
