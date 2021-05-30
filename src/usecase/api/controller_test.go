package api

import (
	"bytes"
	"github.com/hansels/currency_exchange_ovo_backend/src/usecase"
	"github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPI_Ping(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	w := &httptest.ResponseRecorder{}
	r := &http.Request{Body: ioutil.NopCloser(new(bytes.Buffer))}

	module := &usecase.Module{}
	api := New(module)
	resp := api.Ping(w, r)

	g.Expect(resp.Data).Should(gomega.Equal("PING!"))
	g.Expect(resp.Error).Should(gomega.BeNil())
	g.Expect(resp.Code).Should(gomega.Equal("200"))
}
