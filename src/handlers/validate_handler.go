package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	cardNumber := r.URL.Query().Get("number")
	cardMonthExpiration, err := strconv.Atoi(r.URL.Query().Get("month"))
	cardYearExpiration, err := strconv.Atoi(r.URL.Query().Get("year"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if &cardNumber == nil || &cardMonthExpiration == nil || &cardYearExpiration == nil {
		http.Error(w, "Invalid parameters", http.StatusBadRequest)
		return
	}
	log.Printf("Validating card: %s", cardNumber)

	var (
		valid    bool
		errorObj *valerror.Error
		header   int
		response models.ValidateResponse
	)

	// Validation
	valid, errorObj = validators.ValidateCard(Card(cardNumber, cardYearExpiration, cardMonthExpiration))

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
		case valerror.InvalidNumber:
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

func Card(cardNum string, year int, month int) models.Card {
	return models.Card{
		CardNumber:          cardNum,
		CardYearExpiration:  year,
		CardMonthExpiration: month,
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
