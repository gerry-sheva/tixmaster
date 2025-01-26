package event

import (
	"net/http"

	"github.com/gerry-sheva/tixmaster/pkg/common/apierror"
	"github.com/gerry-sheva/tixmaster/pkg/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EventApi struct {
	dbpool *pgxpool.Pool
}

func New(dbpool *pgxpool.Pool) *EventApi {
	return &EventApi{
		dbpool,
	}
}

func (api *EventApi) NewEvent(w http.ResponseWriter, r *http.Request) {
	var p NewEventInput
	if err := util.ReadJSON(w, r, &p); err != nil {
		apierror.Write(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := newEvent(r.Context(), api.dbpool, &p); err != nil {
		apierror.ServerErrorResponse(w)
		return
	}

	return
}
