package api

import (
	"net/http"

	"github.com/gerry-sheva/tixmaster/pkg/util"
)

func (app *app) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "status: available")
	email := r.Context().Value("sub")
	util.WriteJSON(w, http.StatusOK, util.Envelope{"sub": email}, nil)
}
