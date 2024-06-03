package wrapper

import (
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var (
	ErrInvalidArgument = errors.New("invalid argument")
	ErrIsNotInteger    = errors.New("value is not an integer")
)

const (
	ClientSaltParam = "salt"
	ClientIDParam   = "client_id"
)

func GetStringQueryParam(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}

func GetUIntQueryParam(r *http.Request, key string) (res uint64, err error) {
	raw := r.URL.Query().Get(key)
	if raw == "" {
		return 0, nil
	}
	res, err = strconv.ParseUint(raw, 10, 64)
	if err != nil {
		return 0, ErrIsNotInteger
	}
	return res, nil
}

func GetIntQueryParam(r *http.Request, key string) (res int64, err error) {
	raw := r.URL.Query().Get(key)
	if raw == "" {
		return 0, nil
	}
	res, err = strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return 0, ErrIsNotInteger
	}
	return res, nil
}

func GetUIntPathParam(r *http.Request, param string) (id uint64, err error) {
	vars := mux.Vars(r)
	idStr, ok := vars[param]
	if !ok {
		return 0, ErrIsNotInteger
	}

	id, err = strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return 0, ErrIsNotInteger
	}
	return id, nil
}

func GetIntPathParam(r *http.Request, param string) (id int64, err error) {
	vars := mux.Vars(r)
	idStr, ok := vars[param]
	if !ok {
		return 0, ErrIsNotInteger
	}

	id, err = strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, ErrIsNotInteger
	}
	return id, nil
}
