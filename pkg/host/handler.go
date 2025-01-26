package host

import (
	"net/http"

	e "github.com/gerry-sheva/tixmaster/pkg/common/error"
	"github.com/gerry-sheva/tixmaster/pkg/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

type HostApi struct {
	dbpool *pgxpool.Pool
	rwJSON *util.RwJSON
}

func New(dbpool *pgxpool.Pool, rwJSON *util.RwJSON) *HostApi {
	return &HostApi{
		dbpool,
		rwJSON,
	}
}

func (api *HostApi) NewHost(w http.ResponseWriter, r *http.Request) {
	var p NewHostInput
	if err := api.rwJSON.Read(w, r, &p); err != nil {
		e.ErrorResponse(w, http.StatusBadRequest, err.Error())
	}

	if err := newHost(r.Context(), api.dbpool, &p); err != nil {
		e.ServerErrorResponse(w)
		return
	}
}
