package user

import (
	"net/http"
	"utils"
)

type UserRouteHandler struct{}

func (h *UserRouteHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	token, rest := utils.ParsePathToken(req.URL.Path)
	req.URL.Path = rest

	switch token {
	case "register":
		Register(writer, req)
	default:
		http.Error(writer, "Endpoint not found", http.StatusNotFound)
	}
}
