package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"validate-service/src/models"
	valerror "validate-service/src/models/error"
	"validate-service/src/validators"
)

func ValidateHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody models.ValidateRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var valid bool
	var errorObj *valerror.Error
	var header int
	var response models.ValidateResponse

	// Validation
	valid, errorObj = validators.ValidateCard(mapRequestToCard(requestBody))

	// Header
	if errorObj == nil {
		header = http.StatusOK
		response = models.ValidateResponse{
			Valid: valid,
		}
	} else {
		switch errorObj.Code {
		case valerror.Empty:
			header = http.StatusBadRequest
		case valerror.Expired:
			header = http.StatusBadRequest
		case valerror.InvalidMonth:
			header = http.StatusBadRequest
		case valerror.InvalidYear:
			header = http.StatusBadRequest
		}
		response = models.ValidateResponse{
			Valid: valid,
			Error: mapErrorToErrorResponse(errorObj),
		}
	}

	w.WriteHeader(header)
	json.NewEncoder(w).Encode(response)
	return
}

func mapRequestToCard(r models.ValidateRequest) models.Card {
	return models.Card{
		CardNumber:          r.CardNumber,
		CardYearExpiration:  r.CardYearExpiration,
		CardMonthExpiration: r.CardMonthExpiration,
	}
}

func mapErrorToErrorResponse(e *valerror.Error) *valerror.ErrorResponse {
	if e != nil {
		return &valerror.ErrorResponse{
			Code:    formatCode(e.Code),
			Message: e.Message,
		}
	}
	return nil
}

func formatCode(code byte) string {
	return fmt.Sprintf("%03d", code)
}
