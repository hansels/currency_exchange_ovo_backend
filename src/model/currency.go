package model

type Currency struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Value float32 `json:"value"`
}
