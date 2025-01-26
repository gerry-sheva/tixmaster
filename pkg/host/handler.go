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
	var p NewHostInput
	if err := util.ReadJSON(w, r, &p); err != nil {
		apierror.Write(w, http.StatusBadRequest, err.Error())
	}

	if err := newHost(r.Context(), api.dbpool, &p); err != nil {
		apierror.ServerErrorResponse(w)
		return
	}
}
