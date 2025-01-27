package host

import (
	"net/http"

	"github.com/gerry-sheva/tixmaster/pkg/common/apierror"
	"github.com/gerry-sheva/tixmaster/pkg/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

type HostApi struct {
	dbpool *pgxpool.Pool
}

func New(dbpool *pgxpool.Pool) *HostApi {
	return &HostApi{
		dbpool,
	}
}

func (api *HostApi) NewHost(w http.ResponseWriter, r *http.Request) {
	var input NewHostInput
	if err := util.ReadJSONForm(r.FormValue("host"), &input); err != nil {
		apierror.Write(w, http.StatusBadRequest, err.Error())
		return
	}

	host, err := newHost(r.Context(), api.dbpool, &input)
	if err != nil {
		apierror.ServerErrorResponse(w)
		return
	}

	util.WriteJSON(w, http.StatusOK, util.Envelope{"host": host}, nil)
}
