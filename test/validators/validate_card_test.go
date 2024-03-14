package validators

import (
	"testing"
	"time"
	"validate-service/src/models"
	valerror "validate-service/src/models/error"
	"validate-service/src/validators"
)

func TestValidateCardEmpty(t *testing.T) {
	card := models.Card{} // Empty card
	valid, err := validators.ValidateCard(card)

	if valid {
		t.Error("Expected card to be invalid")
	}
	if err == nil || err.Code != valerror.Empty {
		t.Error("Expected empty card error")
	}
}

func TestValidateCardExpired(t *testing.T) {
	card := models.Card{
		CardNumber:          "4111111111111111",
		CardYearExpiration:  time.Now().Year() - 1, // Past year
		CardMonthExpiration: int(time.Now().Month()),
	}
	valid, err := validators.ValidateCard(card)

	if valid {
		t.Error("Expected card to be invalid")
	}
	if err == nil || err.Code != valerror.Expired {
		t.Errorf("Expected %d error, but was %d", valerror.Expired, err.Code)
	}
}

func TestValidateCardInvalidMonth(t *testing.T) {
	card := models.Card{
		CardNumber:          "4111111111111111",
		CardYearExpiration:  time.Now().Year(),
		CardMonthExpiration: 13, // Invalid month
	}
	valid, err := validators.ValidateCard(card)

	if valid {
		t.Error("Expected card to be invalid")
	}
	if err == nil || err.Code != valerror.InvalidMonth {
		t.Error("Expected invalid month error")
	}
}

func TestValidateCardInvalidYear(t *testing.T) {
	card := models.Card{
		CardNumber:          "4111111111111111",
		CardYearExpiration:  3000, // Invalid year
		CardMonthExpiration: int(time.Now().Month()),
	}
	valid, err := validators.ValidateCard(card)

	if valid {
		t.Error("Expected card to be invalid")
	}
	if err == nil || err.Code != valerror.InvalidYear {
		t.Error("Expected invalid year error")
	}
}

func TestValidateCardValid(t *testing.T) {
	card := models.Card{
		CardNumber:          "1234567890123456",
		CardYearExpiration:  time.Now().Year() + 1, // Future year
		CardMonthExpiration: int(time.Now().Month()),
	}
	valid, err := validators.ValidateCard(card)

	if !valid {
		t.Error("Expected card to be valid")
	}
	if err != nil {
		t.Error("Expected no error")
	}
}
