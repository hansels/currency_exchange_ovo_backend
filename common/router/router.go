package router

import (
	"context"
	"errors"
	"github.com/hansels/currency_exchange_ovo_backend/common/response"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

type Handle func(http.ResponseWriter, *http.Request) *response.JSONResponse

type WrittenResponseWriter struct {
	http.ResponseWriter
	written bool
}

func (w *WrittenResponseWriter) Written() bool {
	return w.written
}

func HandleNow(fullPath string, handle Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*time.Duration(30))
		defer cancel()

		ctx = context.WithValue(ctx, "HTTPParams", ps)

		r.Header.Set("routePath", fullPath)
		r = r.WithContext(ctx)

		respChan := make(chan *response.JSONResponse)
		go func() {
			defer panicRecover(respChan, r, fullPath)
			resp := handle(w, r)
			respChan <- resp
		}()

		select {
		case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				w.WriteHeader(http.StatusGatewayTimeout)
				w.Write([]byte("timeout"))
			}
		case resp := <-respChan:
			if resp != nil {
				resp.Send(w)
			} else {
				if w, ok := w.(*WrittenResponseWriter); ok && !w.Written() {
					log.Println("Error nil response from the handler")
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(""))
				}
			}
		}

		return
	}
}

func panicRecover(resp chan *response.JSONResponse, r *http.Request, path string) {
	if recov := recover(); recov != nil {
		var e error
		switch t := recov.(type) {
		case string:
			e = errors.New(t)
		case error:
			e = t
		default:
			e = errors.New("Unknown error")
		}

		response.NewJSONResponse().SetError(e).SetMessage("app panic due to " + e.Error()).SetData(r.URL.Query())
	}
}
