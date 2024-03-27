package httpserver

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Routes(appName string) http.Handler {
	mux := mux.NewRouter()
	return mux
}
