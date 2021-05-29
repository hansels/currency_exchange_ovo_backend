package model

type Currency struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type CurrencySingleData struct {
	Id string `json:"id"`
}

type CurrencyCountData struct {
	Ids   []string `json:"ids"`
	Value float64  `json:"value"`
}

type CurrencyCountResponse struct {
	Id    string  `json:"id"`
	Count float64 `json:"count"`
}
