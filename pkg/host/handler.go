package host

import (
	"net/http"

	"github.com/gerry-sheva/tixmaster/pkg/common/apierror"
	"github.com/gerry-sheva/tixmaster/pkg/util"
	"github.com/imagekit-developer/imagekit-go"
	"github.com/jackc/pgx/v5/pgxpool"
)

type HostApi struct {
	dbpool *pgxpool.Pool
	ik     *imagekit.ImageKit
}

func New(dbpool *pgxpool.Pool, ik *imagekit.ImageKit) *HostApi {
	return &HostApi{
		dbpool,
		ik,
	}
}

func (api *HostApi) NewHost(w http.ResponseWriter, r *http.Request) {
	var input NewHostInput
	avatar, _, err := r.FormFile("host_avatar")
	if err != nil {
		apierror.Write(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := util.ReadJSONForm(r.FormValue("host_data"), &input); err != nil {
		apierror.Write(w, http.StatusBadRequest, err.Error())
		return
	}

	host, err := newHost(r.Context(), api.dbpool, api.ik, avatar, &input)
	if err != nil {
		apierror.ServerErrorResponse(w)
		return
	}

	util.WriteJSON(w, http.StatusOK, util.Envelope{"host": host}, nil)
}
