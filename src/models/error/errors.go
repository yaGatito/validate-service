package error

const (
	Empty byte = 1 + iota
	InvalidNumber
	Expired
	InvalidMonth
	InvalidYear
)

const (
	EmptyMessage         string = "Empty card number"
	InvalidNumberMessage        = "Card number invalid"
	ExpiredMessage              = "Card is expired"
	InvalidMonthMessage         = "Invalid month"
	InvalidYearMessage          = "Invalid year"
)

type Error struct {
	Code    byte   `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
