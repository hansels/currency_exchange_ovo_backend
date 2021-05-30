package api

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"github.com/hansels/currency_exchange_ovo_backend/common/response"
	"github.com/hansels/currency_exchange_ovo_backend/src/model"
	"log"
	"net/http"
)

var currenciesMap map[string]map[string]interface{}

func (a *API) Ping(w http.ResponseWriter, r *http.Request) *response.JSONResponse {
	log.Println("[GET] /ping Called!")
	return response.NewJSONResponse().SetData("PING!")
}

func (a *API) Currencies(w http.ResponseWriter, r *http.Request) *response.JSONResponse {
	err := setCurrencyMap(a.Module.Firestore)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest).SetMessage("Bad Request")
	}

	log.Println("[GET] /currencies Called!")
	return response.NewJSONResponse().SetData(currenciesMap)
}

func (a *API) Currency(w http.ResponseWriter, r *http.Request) *response.JSONResponse {
	err := setCurrencyMap(a.Module.Firestore)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest).SetMessage("Bad Request")
	}

	var currencySingle model.CurrencySingleData
	err = json.NewDecoder(r.Body).Decode(&currencySingle)
	if err != nil {
		log.Fatalf("Currency Single Data Json Decode Error : %+v", err)
		return response.NewJSONResponse().SetError(response.ErrBadRequest).SetMessage("Bad Request")
	}

	data := currenciesMap[currencySingle.Id]
	if data == nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest).SetMessage("Bad Request")
	}

	log.Println("[POST] /currency Called!")
	return response.NewJSONResponse().SetData(data)
}

func (a *API) Calculate(w http.ResponseWriter, r *http.Request) *response.JSONResponse {
	err := setCurrencyMap(a.Module.Firestore)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest).SetMessage("Bad Request")
	}

	var currencyCount model.CurrencyCountData
	err = json.NewDecoder(r.Body).Decode(&currencyCount)
	if err != nil {
		log.Fatalf("Currency Count Data Json Decode Error : %+v", err)
		return response.NewJSONResponse().SetError(response.ErrBadRequest).SetMessage("Bad Request")
	}

	var data []model.CurrencyCountResponse
	for _, element := range currencyCount.Ids {
		exchangeCurrency := (currenciesMap[element]["value"]).(float64)
		count := exchangeCurrency * currencyCount.Value
		data = append(data, model.CurrencyCountResponse{Id: element, Count: count})
	}

	log.Println("[POST] /count Called!")
	return response.NewJSONResponse().SetData(data)
}

func setCurrencyMap(firestore *firestore.Client) error {
	if currenciesMap != nil {
		return nil
	}
	currenciesMap = map[string]map[string]interface{}{}

	ctx := context.Background()
	docs := firestore.Collection("currency").Documents(ctx)

	listDoc, err := docs.GetAll()
	if err != nil {
		log.Fatalf("Read to Firestore error : %+v", err)
		return err
	}

	for _, snapshot := range listDoc {
		if snapshot.Exists() {
			data := snapshot.Data()
			id := data["id"].(string)
			currenciesMap[id] = data
		}
	}

	return nil
}
