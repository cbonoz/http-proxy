package models

type CustomRequest struct {
	Url  string                  `json:"url"`
	Type string                  `json:"type"`
	Hash string                  `json:"hash"`
	Body *map[string]interface{} `json:"body"`
}
