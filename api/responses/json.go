package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Message string                 `json:"menssage"`
			Status  int                    `json:"status"`
			Error   bool                   `json:"error"`
			Data    map[string]interface{} `json:"data"`
		}{
			Message: err.Error(),
			Status:  500,
			Error:   true,
			Data:    nil,
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
