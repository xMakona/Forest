package auth

import (
	"net/http"
	"utils"
)

type AuthRouteHandler struct{}

func (h *AuthRouteHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	token, rest := utils.ParsePathToken(req.URL.Path)
	req.URL.Path = rest

	switch token {
	case "login":
		Auth(writer, req)
	default:
		http.Error(writer, "Endpoint not found", http.StatusNotFound)
	}
}
