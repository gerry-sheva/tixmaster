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
	if err := util.ReadJSONForm(r.FormValue("venue"), &input); err != nil {
		apierror.Write(w, http.StatusBadRequest, err.Error())
		return
	}

	venue, err := newVenue(r.Context(), api.dbpool, input)
	if err != nil {
		apierror.ServerErrorResponse(w)
		return
	}

	util.WriteJSON(w, http.StatusOK, util.Envelope{"venue": venue}, nil)
}
