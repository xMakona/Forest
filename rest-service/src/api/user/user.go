package user

import (
	"fmt"
	"net/http"
)

func Register(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(writer, "Method not accepted for this endpoint", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(writer, "Registration endpoint reached")
	req.ParseForm()
}
