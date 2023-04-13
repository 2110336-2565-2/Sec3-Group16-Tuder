package schemas

import (
	"time"

	"github.com/google/uuid"
)

type SchemaQRCode struct {
	QrCodeUrl string `json:"qr_code_url"`
}

type SchemaGetQRCode struct {
	Course_id uuid.UUID `json:"course_id"`
	Amount    int       `json:"amount" validate:"required"`
}

type Card struct {
	Country    string `json:"country"`
	City       string `json:"city"`
	Bank       string `json:"bank"`
	PostalCode string `json:"postal_code"`
	Financing  string `json:"financing"`
	LastDigits string `json:"last_digits"`
	Brand      string `json:"brand"`

	ExpirationMonth time.Month `json:"expiration_month"`
	ExpirationYear  int        `json:"expiration_year"`

	Fingerprint       string `json:"fingerprint"`
	Name              string `json:"name"`
	SecurityCodeCheck bool   `json:"security_code_check"`
}
