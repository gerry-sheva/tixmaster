package event

import (
	"net/http"

	"github.com/gerry-sheva/tixmaster/pkg/common/apierror"
	"github.com/gerry-sheva/tixmaster/pkg/util"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/meilisearch/meilisearch-go"
)

type EventApi struct {
	dbpool      *pgxpool.Pool
	meilisearch meilisearch.ServiceManager
}

func New(dbpool *pgxpool.Pool, meilisearch meilisearch.ServiceManager) *EventApi {
	return &EventApi{
		dbpool,
		meilisearch,
	}
}

func (api *EventApi) NewEvent(w http.ResponseWriter, r *http.Request) {
	var input NewEventInput
	if err := util.ReadJSONForm(r.FormValue("event"), &input); err != nil {
		apierror.Write(w, http.StatusBadRequest, err.Error())
		return
	}

	event, err := newEvent(r.Context(), api.dbpool, api.meilisearch, &input)
	if err != nil {
		apierror.ServerErrorResponse(w)
		return
	}

	util.WriteJSON(w, http.StatusOK, util.Envelope{"event": event}, nil)
}
