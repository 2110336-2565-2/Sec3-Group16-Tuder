package schemas

import (
	"time"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/google/uuid"
)

type SchemaQRCode struct {
	QrCodeUrl string       `json:"qr_code_url"`
	Charge_id string       `json:"charge_id"`
	Payment   *ent.Payment `json:"payment"`
}

type SchemaGetQRCodeForCoursePayment struct {
	Course_id uuid.UUID `json:"course_id"`
	Amount    int       `json:"amount" validate:"required"`
}

type SchemaGetQRCodeForTuitionFree struct {
	Appointment_id uuid.UUID `json:"appointment_id"`
	Tutor_id       uuid.UUID `json:"tutor_id"`
	Amount         int       `json:"amount" validate:"required"`
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

type SchemaChargeWebhook struct {
	Object            string        `json:"object"`
	ID                string        `json:"id"`
	Livemode          bool          `json:"livemode"`
	Location          string        `json:"location"`
	WebhookDeliveries []interface{} `json:"webhook_deliveries"`
	Data              struct {
		Object            string `json:"object"`
		ID                string `json:"id"`
		Location          string `json:"location"`
		Amount            int    `json:"amount"`
		AuthorizationType string `json:"authorization_type"`
		AuthorizedAmount  int    `json:"authorized_amount"`
		CapturedAmount    int    `json:"captured_amount"`
		Net               int    `json:"net"`
		Fee               int    `json:"fee"`
		FeeVat            int    `json:"fee_vat"`
		Interest          int    `json:"interest"`
		InterestVat       int    `json:"interest_vat"`
		FundingAmount     int    `json:"funding_amount"`
		RefundedAmount    int    `json:"refunded_amount"`
		TransactionFees   struct {
			FeeFlat string `json:"fee_flat"`
			FeeRate string `json:"fee_rate"`
			VatRate string `json:"vat_rate"`
		} `json:"transaction_fees"`
		PlatformFee struct {
			Fixed      string `json:"fixed"`
			Amount     string `json:"amount"`
			Percentage string `json:"percentage"`
		} `json:"platform_fee"`
		Currency        string `json:"currency"`
		FundingCurrency string `json:"funding_currency"`
		IP              string `json:"ip"`
		Refunds         struct {
			Object   string        `json:"object"`
			Data     []interface{} `json:"data"`
			Limit    int           `json:"limit"`
			Offset   int           `json:"offset"`
			Total    int           `json:"total"`
			Location string        `json:"location"`
			Order    string        `json:"order"`
			From     time.Time     `json:"from"`
			To       time.Time     `json:"to"`
		} `json:"refunds"`
		Link        string `json:"link"`
		Description string `json:"description"`
		Metadata    struct {
		} `json:"metadata"`
		Card   string `json:"card"`
		Source struct {
			Object          string    `json:"object"`
			ID              string    `json:"id"`
			Livemode        bool      `json:"livemode"`
			Location        string    `json:"location"`
			Amount          int       `json:"amount"`
			Barcode         string    `json:"barcode"`
			Bank            string    `json:"bank"`
			CreatedAt       time.Time `json:"created_at"`
			Currency        string    `json:"currency"`
			Email           string    `json:"email"`
			Flow            string    `json:"flow"`
			InstallmentTerm string    `json:"installment_term"`
			AbsorptionType  string    `json:"absorption_type"`
			Name            string    `json:"name"`
			MobileNumber    string    `json:"mobile_number"`
			PhoneNumber     string    `json:"phone_number"`
			PlatformType    string    `json:"platform_type"`
			ScannableCode   struct {
				Object string `json:"object"`
				Type   string `json:"type"`
				Image  struct {
					Object      string    `json:"object"`
					Livemode    bool      `json:"livemode"`
					ID          string    `json:"id"`
					Deleted     bool      `json:"deleted"`
					Filename    string    `json:"filename"`
					Location    string    `json:"location"`
					Kind        string    `json:"kind"`
					DownloadURI string    `json:"download_uri"`
					CreatedAt   time.Time `json:"created_at"`
				} `json:"image"`
			} `json:"scannable_code"`
			Billing                  string        `json:"billing"`
			Shipping                 string        `json:"shipping"`
			Items                    []interface{} `json:"items"`
			References               string        `json:"references"`
			StoreID                  string        `json:"store_id"`
			StoreName                string        `json:"store_name"`
			TerminalID               string        `json:"terminal_id"`
			Type                     string        `json:"type"`
			ZeroInterestInstallments string        `json:"zero_interest_installments"`
			ChargeStatus             string        `json:"charge_status"`
			ReceiptAmount            string        `json:"receipt_amount"`
			Discounts                []interface{} `json:"discounts"`
		} `json:"source"`
		Schedule                 string    `json:"schedule"`
		Customer                 string    `json:"customer"`
		Dispute                  string    `json:"dispute"`
		Transaction              string    `json:"transaction"`
		FailureCode              string    `json:"failure_code"`
		FailureMessage           string    `json:"failure_message"`
		Status                   string    `json:"status"`
		AuthorizeURI             string    `json:"authorize_uri"`
		ReturnURI                string    `json:"return_uri"`
		CreatedAt                time.Time `json:"created_at"`
		PaidAt                   string    `json:"paid_at"`
		ExpiresAt                time.Time `json:"expires_at"`
		ExpiredAt                string    `json:"expired_at"`
		ReversedAt               string    `json:"reversed_at"`
		ZeroInterestInstallments bool      `json:"zero_interest_installments"`
		Branch                   string    `json:"branch"`
		Terminal                 string    `json:"terminal"`
		Device                   string    `json:"device"`
		Authorized               bool      `json:"authorized"`
		Capturable               bool      `json:"capturable"`
		Capture                  bool      `json:"capture"`
		Disputable               bool      `json:"disputable"`
		Livemode                 bool      `json:"livemode"`
		Refundable               bool      `json:"refundable"`
		Reversed                 bool      `json:"reversed"`
		Reversible               bool      `json:"reversible"`
		Voided                   bool      `json:"voided"`
		Paid                     bool      `json:"paid"`
		Expired                  bool      `json:"expired"`
	} `json:"data"`
	Key       string    `json:"key"`
	CreatedAt time.Time `json:"created_at"`
}
