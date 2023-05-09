package routes

import "net/http"

func HomeHandlers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GTD APIRest"))

}
