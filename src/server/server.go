package server

import (
	"github.com/hansels/currency_exchange_ovo_backend/src/usecase"
	"github.com/hansels/currency_exchange_ovo_backend/src/usecase/api"
	"github.com/julienschmidt/httprouter"
	"log"

	"github.com/rs/cors"
	"net/http"
)

type Opts struct {
	ListenAddress string
	Modules       *usecase.Module
}

type Handler struct {
	options     *Opts
	listenErrCh chan error
}

func New(o *Opts) *Handler {
	handler := &Handler{options: o}
	return handler
}

func (h *Handler) Run() {
	log.Printf("Listening on %s", h.options.ListenAddress)

	c := cors.New(cors.Options{
		AllowedHeaders: []string{"X-Requested-With", "Authorization", "Content-Type", "X-Authorization"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"FETCH", "GET", "POST", "DELETE", "PUT", "OPTIONS"},
	})

	router := httprouter.New()
	api.New(h.options.Modules).Register(router)

	handler := c.Handler(router)
	h.listenErrCh <- http.ListenAndServe(h.options.ListenAddress, handler)
}

func (h *Handler) ListenError() <-chan error {
	return h.listenErrCh
}
