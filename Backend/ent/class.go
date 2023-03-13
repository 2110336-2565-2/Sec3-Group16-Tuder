// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/class"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/paymenthistory"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/schedule"
	"github.com/google/uuid"
)

// Class is the model entity for the Class schema.
type Class struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// ReviewAvaliable holds the value of the "review_avaliable" field.
	ReviewAvaliable bool `json:"review_avaliable,omitempty"`
	// TotalHour holds the value of the "total_hour" field.
	TotalHour time.Time `json:"total_hour,omitempty"`
	// SuccessHour holds the value of the "success_hour" field.
	SuccessHour time.Time `json:"success_hour,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ClassQuery when eager-loading is set.
	Edges                 ClassEdges `json:"edges"`
	payment_history_class *uuid.UUID
	schedule_class        *uuid.UUID
}

// ClassEdges holds the relations/edges for other nodes in the graph.
type ClassEdges struct {
	// Match holds the value of the match edge.
	Match []*Match `json:"match,omitempty"`
	// Schedule holds the value of the schedule edge.
	Schedule *Schedule `json:"schedule,omitempty"`
	// PaymentHistory holds the value of the payment_history edge.
	PaymentHistory *PaymentHistory `json:"payment_history,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// MatchOrErr returns the Match value or an error if the edge
// was not loaded in eager-loading.
func (e ClassEdges) MatchOrErr() ([]*Match, error) {
	if e.loadedTypes[0] {
		return e.Match, nil
	}
	return nil, &NotLoadedError{edge: "match"}
}

// ScheduleOrErr returns the Schedule value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ClassEdges) ScheduleOrErr() (*Schedule, error) {
	if e.loadedTypes[1] {
		if e.Schedule == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: schedule.Label}
		}
		return e.Schedule, nil
	}
	return nil, &NotLoadedError{edge: "schedule"}
}

// PaymentHistoryOrErr returns the PaymentHistory value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ClassEdges) PaymentHistoryOrErr() (*PaymentHistory, error) {
	if e.loadedTypes[2] {
		if e.PaymentHistory == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: paymenthistory.Label}
		}
		return e.PaymentHistory, nil
	}
	return nil, &NotLoadedError{edge: "payment_history"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Class) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case class.FieldReviewAvaliable:
			values[i] = new(sql.NullBool)
		case class.FieldTotalHour, class.FieldSuccessHour:
			values[i] = new(sql.NullTime)
		case class.FieldID:
			values[i] = new(uuid.UUID)
		case class.ForeignKeys[0]: // payment_history_class
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case class.ForeignKeys[1]: // schedule_class
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Class", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Class fields.
func (c *Class) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case class.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case class.FieldReviewAvaliable:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field review_avaliable", values[i])
			} else if value.Valid {
				c.ReviewAvaliable = value.Bool
			}
		case class.FieldTotalHour:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field total_hour", values[i])
			} else if value.Valid {
				c.TotalHour = value.Time
			}
		case class.FieldSuccessHour:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field success_hour", values[i])
			} else if value.Valid {
				c.SuccessHour = value.Time
			}
		case class.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field payment_history_class", values[i])
			} else if value.Valid {
				c.payment_history_class = new(uuid.UUID)
				*c.payment_history_class = *value.S.(*uuid.UUID)
			}
		case class.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field schedule_class", values[i])
			} else if value.Valid {
				c.schedule_class = new(uuid.UUID)
				*c.schedule_class = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryMatch queries the "match" edge of the Class entity.
func (c *Class) QueryMatch() *MatchQuery {
	return NewClassClient(c.config).QueryMatch(c)
}

// QuerySchedule queries the "schedule" edge of the Class entity.
func (c *Class) QuerySchedule() *ScheduleQuery {
	return NewClassClient(c.config).QuerySchedule(c)
}

// QueryPaymentHistory queries the "payment_history" edge of the Class entity.
func (c *Class) QueryPaymentHistory() *PaymentHistoryQuery {
	return NewClassClient(c.config).QueryPaymentHistory(c)
}

// Update returns a builder for updating this Class.
// Note that you need to call Class.Unwrap() before calling this method if this Class
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Class) Update() *ClassUpdateOne {
	return NewClassClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Class entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Class) Unwrap() *Class {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Class is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Class) String() string {
	var builder strings.Builder
	builder.WriteString("Class(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("review_avaliable=")
	builder.WriteString(fmt.Sprintf("%v", c.ReviewAvaliable))
	builder.WriteString(", ")
	builder.WriteString("total_hour=")
	builder.WriteString(c.TotalHour.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("success_hour=")
	builder.WriteString(c.SuccessHour.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Classes is a parsable slice of Class.
type Classes []*Class
