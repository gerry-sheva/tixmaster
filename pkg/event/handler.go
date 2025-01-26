package event

import (
	"net/http"

	e "github.com/gerry-sheva/tixmaster/pkg/common/error"
	"github.com/gerry-sheva/tixmaster/pkg/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EventApi struct {
	dbpool *pgxpool.Pool
	rwJSON *util.RwJSON
}

func New(dbpool *pgxpool.Pool, rwJSON *util.RwJSON) *EventApi {
	return &EventApi{
		dbpool,
		rwJSON,
	}
}

func (api *EventApi) NewEvent(w http.ResponseWriter, r *http.Request) {
	var p NewEventInput
	if err := api.rwJSON.Read(w, r, &p); err != nil {
		e.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := newEvent(r.Context(), api.dbpool, &p); err != nil {
		e.ServerErrorResponse(w)
		return
	}

	return
}
