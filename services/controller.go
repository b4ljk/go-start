package services

import (
	"back/structs"
	"net/http"
)

type Controller struct {
	DB       *DBService
	Response *structs.Response
}

func (co Controller) SetBody(body interface{}) {
	co.Response.StatusCode = http.StatusOK
	co.Response.Body.Message = ""
	co.Response.Body.Body = body
}

// SetError error response
func (co Controller) SetError(code int, message string) {
	co.Response.StatusCode = code
	co.Response.Body.Message = message
	co.Response.Body.Body = nil
}

// GetBody in response
func (co Controller) GetBody() (int, interface{}) {
	return co.Response.StatusCode, co.Response.Body
}

func NewController() (*Controller, error) {
	db, err := NewDBService()
	if err != nil {
		return nil, err
	}

	return &Controller{
		DB: db,
		Response: &structs.Response{StatusCode: http.StatusOK,
			Body: structs.ResponseBody{Message: "", Body: nil}},
	}, nil
}
