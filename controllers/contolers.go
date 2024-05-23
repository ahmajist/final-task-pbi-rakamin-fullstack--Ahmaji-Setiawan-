package controllers

import (
	"fmt"
	"net/http"
)

// ProtectedHandler
func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	// mendapatkan informasi pengguna dari context
	ctx := r.Context()
	username := ctx.Value("username").(string)

	fmt.Fprintf(w, "Protected resource accessed by user: %s\n", username)
}
