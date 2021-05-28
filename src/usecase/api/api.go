package api

import (
	myRouter "github.com/hansels/currency_exchange_ovo_backend/common/router"
	"github.com/hansels/currency_exchange_ovo_backend/src/usecase"
	"github.com/julienschmidt/httprouter"
)

func (a *API) Register(router *httprouter.Router) {
	router.GET("/ping", myRouter.HandleNow("/ping", a.Ping))
}

type API struct {
	Module *usecase.Module
}

func New(module *usecase.Module) *API {
	return &API{Module: module}
}
