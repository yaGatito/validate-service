package error

const (
	Empty byte = 1 + iota
	Expired
	InvalidMonth
	InvalidYear
)

const (
	EmptyMessage        string = "Empty card number"
	ExpiredMessage             = "Card is expired"
	InvalidMonthMessage        = "Invalid month"
	InvalidYearMessage         = "Invalid year"
)

type Error struct {
	Code    byte   `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
