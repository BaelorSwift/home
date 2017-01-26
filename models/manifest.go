package models

import (
	"encoding/json"
	"io/ioutil"
)

// Manifest ..
type Manifest struct {
	Version      string                   `json:"version"`
	BaseURL      string                   `json:"baseUrl"`
	Pagination   pagination               `json:"pagination"`
	GlobalErrors globalErrors             `json:"globalErrors"`
	Groups       map[string]manifestGroup `json:"groups"`
}
type pagination struct {
	Example string `json:"example"`
}
type globalErrors struct {
	Example string  `json:"example"`
	Errors  []error `json:"errors"`
}
type error struct {
	Code        string `json:"code"`
	Status      int    `json:"status"`
	Description string `json:"description"`
}
type manifestGroup struct {
	Name        string                   `json:"name"`
	Description string                   `json:"description"`
	Idents      map[string]string        `json:"idents"`
	Endpoints   map[string]groupEndpoint `json:"endpoints"`
}
type groupEndpoint struct {
	Name             string          `json:"name"`
	Description      string          `json:"description"`
	Method           string          `json:"method"`
	Path             string          `json:"path"`
	IsAuthenticated  bool            `json:"isAuthenticated"`
	AllowsPagination bool            `json:"allowsPagination"`
	URLParams        []interface{}   `json:"urlParams"`
	Request          endpointRequest `json:"request"`
	Response         httpExample     `json:"response"`
	Errors           []error         `json:"errors"`
}
type endpointRequest struct {
	Example httpExample           `json:"example"`
	Body    []endpointRequestBody `json:"body"`
	Schema  string                `json:"schema"`
}
type endpointRequestBody struct {
	Name        string `json:"name"`
	Required    bool   `json:"required"`
	Description string `json:"description"`
	Example     string `json:"example"`
}
type httpExample struct {
	HTTP string `json:"http"`
	JSON string `json:"json"`
}

// NewManifest ..
func NewManifest() Manifest {
	file, err := ioutil.ReadFile("./manifest.json")
	if err != nil {
		panic(err)
	}

	var manifest Manifest
	json.Unmarshal(file, &manifest)
	return manifest
}
