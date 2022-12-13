package utils

import "net/http"

type HTTPResponse struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewHttpResponse(status int, data interface{}) HTTPResponse {
	switch status {
	case http.StatusBadRequest,
		http.StatusNotFound,
		http.StatusInternalServerError,
		http.StatusUnauthorized,
		http.StatusForbidden,
		http.StatusRequestTimeout:

		if e, ok := data.(error); ok {
			return HTTPResponse{
				Status:  status,
				Success: false,
				Message: e.Error(),
			}
		}

		return HTTPResponse{
			Status:  status,
			Success: false,
			Message: data.(string),
		}
	default:
		return HTTPResponse{
			Status:  status,
			Success: true,
			Data:    data,
		}
	}
}
