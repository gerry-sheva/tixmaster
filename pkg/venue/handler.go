package venue

import (
	"net/http"

	"github.com/gerry-sheva/tixmaster/pkg/common/apierror"
	"github.com/gerry-sheva/tixmaster/pkg/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

type VenueAPI struct {
	dbpool *pgxpool.Pool
}

func New(dbpool *pgxpool.Pool) *VenueAPI {
	return &VenueAPI{
		dbpool,
	}
}

func (api *VenueAPI) NewVenue(w http.ResponseWriter, r *http.Request) {
	var input NewVenueInput
	if err := util.ReadJSON(w, r, &input); err != nil {
		apierror.Write(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := newVenue(r.Context(), api.dbpool, input); err != nil {
		apierror.ServerErrorResponse(w)
		return
	}
}
