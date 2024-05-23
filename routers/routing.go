package routers

import (
	"net/http"
	"FINALTASK-BTPN-SYARIAH/controllers"
	"FINALTASK-BTPN-SYARIAH/middlewares"
)

// DefineRoutes akan mengatur semua rute / endpoint API
func DefineRoutes() {
	http.HandleFunc("/protected", middlewares.JWTAuthMiddleware(controllers.ProtectedHandler))
}
