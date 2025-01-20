package auth

import (
	"net/http"

	e "github.com/gerry-sheva/tixmaster/pkg/common/error"
	"github.com/gerry-sheva/tixmaster/pkg/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UsersAPI struct {
	dbpool *pgxpool.Pool
	rwJSON *util.RwJSON
}

func New(dbpool *pgxpool.Pool, rwJSON *util.RwJSON) *UsersAPI {
	return &UsersAPI{
		dbpool,
		rwJSON,
	}
}

func (api *UsersAPI) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var p AuthInput
	if err := api.rwJSON.Read(w, r, &p); err != nil {
		e.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	v := util.NewValidator()
	p.validate(true, v)

	if !v.Valid() {
		http.Error(w, "Invalidddd", http.StatusBadRequest)
		return
	}

	err := register(r.Context(), api.dbpool, &p)
	if err != nil {
		e.ErrorResponse(w, http.StatusBadRequest, err.Error())
	}
	return
}

func (api *UsersAPI) LoginUser(w http.ResponseWriter, r *http.Request) {

}
