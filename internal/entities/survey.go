package entities

type Survey struct {
	ID         string  `json:"id"`
	MerchantID string  `json:"merchant_id"`
	UserID     string  `json:"user_id"`
	Rating     float32 `json:"rating"`
}

type SurveyAvg struct {
	avg float32 `json:"avg"`
}
