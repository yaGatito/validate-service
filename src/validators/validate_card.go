package validators

import (
	"regexp"
	"time"
	"validate-service/src/models"
	valerror "validate-service/src/models/error"
)

const (
	// The defaults for time diffs
	Day     = 1
	Hour    = 1
	Minute  = 1
	Second  = 1
	MSecond = 1
)

func ValidateCard(card models.Card) (bool, *valerror.Error) {
	if isEmpty(card) {
		return false, &valerror.Error{
			Code:    valerror.Empty,
			Message: valerror.EmptyMessage,
		}
	}

	if isValidNumber(card) {
		return false, &valerror.Error{
			Code:    valerror.InvalidNumber,
			Message: valerror.InvalidNumberMessage,
		}
	}

	if isMonthInvalid(card) {
		return false, &valerror.Error{
			Code:    valerror.InvalidMonth,
			Message: valerror.InvalidMonthMessage,
		}
	}

	if isYearInvalid(card) {
		return false, &valerror.Error{
			Code:    valerror.InvalidYear,
			Message: valerror.InvalidYearMessage,
		}
	}

	if isExpired(card) {
		return false, &valerror.Error{
			Code:    valerror.Expired,
			Message: valerror.ExpiredMessage,
		}
	}
	return true, nil
}

func isExpired(card models.Card) bool {
	return time.Date(card.CardYearExpiration, time.Month(card.CardMonthExpiration)+1, Day, Hour, Minute, Second, MSecond, time.Local).Before(time.Now())
}

func isMonthInvalid(card models.Card) bool {
	return !(card.CardMonthExpiration >= 1 && card.CardMonthExpiration <= 12)
}

func isYearInvalid(card models.Card) bool {
	return !(card.CardYearExpiration > 2000 && card.CardYearExpiration < 2100)
}

func isEmpty(card models.Card) bool {
	return card.CardNumber == ""
}

func isValidNumber(card models.Card) bool {
	pattern := `^\d{12,16}$`
	regex := regexp.MustCompile(pattern)
	return !regex.MatchString(card.CardNumber)
}
