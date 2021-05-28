package api

import (
	"github.com/hansels/currency_exchange_ovo_backend/common/response"
	"log"
	"net/http"
)

func (a *API) Ping(w http.ResponseWriter, r *http.Request) *response.JSONResponse {
	log.Println("PING Called!")
	return response.NewJSONResponse().SetData("PING!")
}
