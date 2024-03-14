package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"validate-service/src/models"
	valerror "validate-service/src/models/error"
	"validate-service/src/validators"
)

func HandleValidatePath(w http.ResponseWriter, r *http.Request) {
	// HTTP Method Switch
	switch r.Method {
	case http.MethodGet:
		log.Println("Handling GET request.")
		validateCard(w, r)
	default:
		w.Header().Set("Allow", "GET")
		http.Error(w, "", http.StatusMethodNotAllowed)
		log.Printf("Handler isn't listening to %s HTTP Methods", r.Method)
		return
	}
}

func validateCard(w http.ResponseWriter, r *http.Request) {
	// Deserialization
	var requestBody models.ValidateRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("Validating card: %s", requestBody.CardNumber)

	var valid bool
	var errorObj *valerror.Error
	var header int
	var response models.ValidateResponse

	// Validation
	valid, errorObj = validators.ValidateCard(mapRequestToCard(requestBody))

	// Headers
	if errorObj == nil {
		header = http.StatusOK
		response = models.ValidateResponse{
			Valid: valid,
		}

		log.Printf("Card valid")
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

		log.Printf("Card invalid reason: %s", response.Error.Message)
	}
	w.WriteHeader(header)
	w.Header().Set("Content-Type", "application/json")

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
