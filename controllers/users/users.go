package users

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func userIndexAction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "User Index Page")
}

func userInfoAction(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	fmt.Fprintf(w, "User %s's info action", params["id"])
}
