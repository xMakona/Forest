package auth

import (
	"fmt"
	"net/http"
)

func Auth(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(writer, "Method not accepted for this endpoint", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(writer, "Auth endpoint reached")
}
