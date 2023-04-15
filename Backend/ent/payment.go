// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/appointment"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/match"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/payment"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	"github.com/google/uuid"
)

// Payment is the model entity for the Payment schema.
type Payment struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// QrPictureURL holds the value of the "qr_picture_url" field.
	QrPictureURL *string `json:"qr_picture_url,omitempty"`
	// PaymentStatus holds the value of the "payment_status" field.
	PaymentStatus payment.PaymentStatus `json:"payment_status,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount int `json:"amount,omitempty"`
	// Currency holds the value of the "currency" field.
	Currency string `json:"currency,omitempty"`
	// ChargeID holds the value of the "charge_id" field.
	ChargeID string `json:"charge_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PaymentQuery when eager-loading is set.
	Edges        PaymentEdges `json:"edges"`
	user_payment *uuid.UUID
}

// PaymentEdges holds the relations/edges for other nodes in the graph.
type PaymentEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Match holds the value of the match edge.
	Match *Match `json:"match,omitempty"`
	// Appointment holds the value of the appointment edge.
	Appointment *Appointment `json:"appointment,omitempty"`
	// PaymentHistory holds the value of the payment_history edge.
	PaymentHistory []*PaymentHistory `json:"payment_history,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PaymentEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// MatchOrErr returns the Match value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PaymentEdges) MatchOrErr() (*Match, error) {
	if e.loadedTypes[1] {
		if e.Match == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: match.Label}
		}
		return e.Match, nil
	}
	return nil, &NotLoadedError{edge: "match"}
}

// AppointmentOrErr returns the Appointment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PaymentEdges) AppointmentOrErr() (*Appointment, error) {
	if e.loadedTypes[2] {
		if e.Appointment == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: appointment.Label}
		}
		return e.Appointment, nil
	}
	return nil, &NotLoadedError{edge: "appointment"}
}

// PaymentHistoryOrErr returns the PaymentHistory value or an error if the edge
// was not loaded in eager-loading.
func (e PaymentEdges) PaymentHistoryOrErr() ([]*PaymentHistory, error) {
	if e.loadedTypes[3] {
		return e.PaymentHistory, nil
	}
	return nil, &NotLoadedError{edge: "payment_history"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Payment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case payment.FieldAmount:
			values[i] = new(sql.NullInt64)
		case payment.FieldQrPictureURL, payment.FieldPaymentStatus, payment.FieldCurrency, payment.FieldChargeID:
			values[i] = new(sql.NullString)
		case payment.FieldCreatedAt, payment.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case payment.FieldID:
			values[i] = new(uuid.UUID)
		case payment.ForeignKeys[0]: // user_payment
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Payment", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Payment fields.
func (pa *Payment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case payment.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pa.ID = *value
			}
		case payment.FieldQrPictureURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field qr_picture_url", values[i])
			} else if value.Valid {
				pa.QrPictureURL = new(string)
				*pa.QrPictureURL = value.String
			}
		case payment.FieldPaymentStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field payment_status", values[i])
			} else if value.Valid {
				pa.PaymentStatus = payment.PaymentStatus(value.String)
			}
		case payment.FieldAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				pa.Amount = int(value.Int64)
			}
		case payment.FieldCurrency:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field currency", values[i])
			} else if value.Valid {
				pa.Currency = value.String
			}
		case payment.FieldChargeID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field charge_id", values[i])
			} else if value.Valid {
				pa.ChargeID = value.String
			}
		case payment.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pa.CreatedAt = value.Time
			}
		case payment.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pa.UpdatedAt = value.Time
			}
		case payment.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_payment", values[i])
			} else if value.Valid {
				pa.user_payment = new(uuid.UUID)
				*pa.user_payment = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Payment entity.
func (pa *Payment) QueryUser() *UserQuery {
	return NewPaymentClient(pa.config).QueryUser(pa)
}

// QueryMatch queries the "match" edge of the Payment entity.
func (pa *Payment) QueryMatch() *MatchQuery {
	return NewPaymentClient(pa.config).QueryMatch(pa)
}

// QueryAppointment queries the "appointment" edge of the Payment entity.
func (pa *Payment) QueryAppointment() *AppointmentQuery {
	return NewPaymentClient(pa.config).QueryAppointment(pa)
}

// QueryPaymentHistory queries the "payment_history" edge of the Payment entity.
func (pa *Payment) QueryPaymentHistory() *PaymentHistoryQuery {
	return NewPaymentClient(pa.config).QueryPaymentHistory(pa)
}

// Update returns a builder for updating this Payment.
// Note that you need to call Payment.Unwrap() before calling this method if this Payment
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Payment) Update() *PaymentUpdateOne {
	return NewPaymentClient(pa.config).UpdateOne(pa)
}

// Unwrap unwraps the Payment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pa *Payment) Unwrap() *Payment {
	_tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Payment is not a transactional entity")
	}
	pa.config.driver = _tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Payment) String() string {
	var builder strings.Builder
	builder.WriteString("Payment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pa.ID))
	if v := pa.QrPictureURL; v != nil {
		builder.WriteString("qr_picture_url=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("payment_status=")
	builder.WriteString(fmt.Sprintf("%v", pa.PaymentStatus))
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", pa.Amount))
	builder.WriteString(", ")
	builder.WriteString("currency=")
	builder.WriteString(pa.Currency)
	builder.WriteString(", ")
	builder.WriteString("charge_id=")
	builder.WriteString(pa.ChargeID)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pa.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pa.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Payments is a parsable slice of Payment.
type Payments []*Payment
