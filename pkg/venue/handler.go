package venue

import (
	"net/http"

	"github.com/gerry-sheva/tixmaster/pkg/common/apierror"
	"github.com/gerry-sheva/tixmaster/pkg/util"
	"github.com/imagekit-developer/imagekit-go"
	"github.com/jackc/pgx/v5/pgxpool"
)

type VenueAPI struct {
	dbpool *pgxpool.Pool
	ik     *imagekit.ImageKit
}

func New(dbpool *pgxpool.Pool, ik *imagekit.ImageKit) *VenueAPI {
	return &VenueAPI{
		dbpool,
		ik,
	}
}

func (api *VenueAPI) NewVenue(w http.ResponseWriter, r *http.Request) {
	var input NewVenueInput
	img, _, err := r.FormFile("venue_image")
	if err != nil {
		apierror.Write(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := util.ReadJSONForm(r.FormValue("venue_data"), &input); err != nil {
		apierror.Write(w, http.StatusBadRequest, err.Error())
		return
	}

	venue, err := newVenue(r.Context(), api.dbpool, api.ik, img, input)
	if err != nil {
		apierror.ServerErrorResponse(w)
		return
	}

	util.WriteJSON(w, http.StatusOK, util.Envelope{"venue": venue}, nil)
}
