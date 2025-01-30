package event

import (
	"net/http"

	"github.com/gerry-sheva/tixmaster/pkg/common/apierror"
	"github.com/gerry-sheva/tixmaster/pkg/util"
	"github.com/imagekit-developer/imagekit-go"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/meilisearch/meilisearch-go"
)

type EventApi struct {
	dbpool      *pgxpool.Pool
	meilisearch meilisearch.ServiceManager
	ik          *imagekit.ImageKit
}

func New(
	dbpool *pgxpool.Pool,
	meilisearch meilisearch.ServiceManager,
	ik *imagekit.ImageKit,
) *EventApi {
	return &EventApi{
		dbpool,
		meilisearch,
		ik,
	}
}

func (api *EventApi) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var input NewEventInput
	thumbnail, _, err := r.FormFile("event_thumbnail")
	if err != nil {
		apierror.Write(w, http.StatusBadRequest, err.Error())
		return
	}

	banner, _, err := r.FormFile("event_banner")
	if err != nil {
		apierror.Write(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := util.ReadJSONForm(r.FormValue("event_data"), &input); err != nil {
		apierror.Write(w, http.StatusBadRequest, err.Error())
		return
	}

	event, err := NewEvent(
		r.Context(),
		api.dbpool,
		api.meilisearch,
		api.ik,
		thumbnail,
		banner,
		&input,
	)
	if err != nil {
		apierror.ServerErrorResponse(w)
		return
	}

	util.WriteJSON(w, http.StatusOK, util.Envelope{"event": event}, nil)
}
