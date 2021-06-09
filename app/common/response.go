package common

import (
	"net/http"
)

type Response interface {
	JSON() []byte
	StatusCode() int
}

func WriteResponse(w http.ResponseWriter, res Response) {
	w.WriteHeader(res.StatusCode())
	_, _ = w.Write(res.JSON())
}
