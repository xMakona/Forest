package router

import (
	"auth"
	"log"
	"net/http"
	"time"
	"user"
	"utils"
)

type RouteHandler struct {
	authRouteHandler *auth.AuthRouteHandler
	userRouteHandler *user.UserRouteHandler
}

func (h *RouteHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	token, rest := utils.ParsePathToken(req.URL.Path)
	req.URL.Path = rest

	switch token {
	case "auth":
		h.authRouteHandler.ServeHTTP(writer, req)
	case "user":
		h.userRouteHandler.ServeHTTP(writer, req)
	default:
		http.Error(writer, "Endpoint not found", http.StatusNotFound)
	}
}

func InitRouter() {
	handler := &RouteHandler{
		authRouteHandler: new(auth.AuthRouteHandler),
		userRouteHandler: new(user.UserRouteHandler),
	}
	s := &http.Server{
		Addr:           ":8080",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
