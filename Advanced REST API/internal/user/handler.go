package user

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type handler struct {
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET("/users", h.GetList)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is list of users"))
}
