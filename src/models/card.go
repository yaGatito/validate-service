package models

import (
	valerror "validate-service/src/models/error"
)

type ValidateRequest struct {
	CardNumber          string `json:"number"`
	CardYearExpiration  int    `json:"year"`
	CardMonthExpiration int    `json:"month"`
}

type Card struct {
	CardNumber          string
	CardYearExpiration  int
	CardMonthExpiration int
}

type ValidateResponse struct {
	Valid bool                    `json:"valid"`
	Error *valerror.ErrorResponse `json:"error,omitempty"`
}
