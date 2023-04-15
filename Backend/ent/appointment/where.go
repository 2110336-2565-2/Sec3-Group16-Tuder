// Code generated by ent, DO NOT EDIT.

package appointment

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Appointment {
	return predicate.Appointment(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Appointment {
	return predicate.Appointment(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Appointment {
	return predicate.Appointment(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Appointment {
	return predicate.Appointment(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Appointment {
	return predicate.Appointment(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Appointment {
	return predicate.Appointment(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Appointment {
	return predicate.Appointment(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Appointment {
	return predicate.Appointment(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Appointment {
	return predicate.Appointment(sql.FieldLTE(FieldID, id))
}

// BeginAt applies equality check predicate on the "begin_at" field. It's identical to BeginAtEQ.
func BeginAt(v time.Time) predicate.Appointment {
	return predicate.Appointment(sql.FieldEQ(FieldBeginAt, v))
}

// EndAt applies equality check predicate on the "end_at" field. It's identical to EndAtEQ.
func EndAt(v time.Time) predicate.Appointment {
	return predicate.Appointment(sql.FieldEQ(FieldEndAt, v))
}

// BeginAtEQ applies the EQ predicate on the "begin_at" field.
func BeginAtEQ(v time.Time) predicate.Appointment {
	return predicate.Appointment(sql.FieldEQ(FieldBeginAt, v))
}

// BeginAtNEQ applies the NEQ predicate on the "begin_at" field.
func BeginAtNEQ(v time.Time) predicate.Appointment {
	return predicate.Appointment(sql.FieldNEQ(FieldBeginAt, v))
}

// BeginAtIn applies the In predicate on the "begin_at" field.
func BeginAtIn(vs ...time.Time) predicate.Appointment {
	return predicate.Appointment(sql.FieldIn(FieldBeginAt, vs...))
}

// BeginAtNotIn applies the NotIn predicate on the "begin_at" field.
func BeginAtNotIn(vs ...time.Time) predicate.Appointment {
	return predicate.Appointment(sql.FieldNotIn(FieldBeginAt, vs...))
}

// BeginAtGT applies the GT predicate on the "begin_at" field.
func BeginAtGT(v time.Time) predicate.Appointment {
	return predicate.Appointment(sql.FieldGT(FieldBeginAt, v))
}

// BeginAtGTE applies the GTE predicate on the "begin_at" field.
func BeginAtGTE(v time.Time) predicate.Appointment {
	return predicate.Appointment(sql.FieldGTE(FieldBeginAt, v))
}

// BeginAtLT applies the LT predicate on the "begin_at" field.
func BeginAtLT(v time.Time) predicate.Appointment {
	return predicate.Appointment(sql.FieldLT(FieldBeginAt, v))
}

// BeginAtLTE applies the LTE predicate on the "begin_at" field.
func BeginAtLTE(v time.Time) predicate.Appointment {
	return predicate.Appointment(sql.FieldLTE(FieldBeginAt, v))
}

// EndAtEQ applies the EQ predicate on the "end_at" field.
func EndAtEQ(v time.Time) predicate.Appointment {
	return predicate.Appointment(sql.FieldEQ(FieldEndAt, v))
}

// EndAtNEQ applies the NEQ predicate on the "end_at" field.
func EndAtNEQ(v time.Time) predicate.Appointment {
	return predicate.Appointment(sql.FieldNEQ(FieldEndAt, v))
}

// EndAtIn applies the In predicate on the "end_at" field.
func EndAtIn(vs ...time.Time) predicate.Appointment {
	return predicate.Appointment(sql.FieldIn(FieldEndAt, vs...))
}

// EndAtNotIn applies the NotIn predicate on the "end_at" field.
func EndAtNotIn(vs ...time.Time) predicate.Appointment {
	return predicate.Appointment(sql.FieldNotIn(FieldEndAt, vs...))
}

// EndAtGT applies the GT predicate on the "end_at" field.
func EndAtGT(v time.Time) predicate.Appointment {
	return predicate.Appointment(sql.FieldGT(FieldEndAt, v))
}

// EndAtGTE applies the GTE predicate on the "end_at" field.
func EndAtGTE(v time.Time) predicate.Appointment {
	return predicate.Appointment(sql.FieldGTE(FieldEndAt, v))
}

// EndAtLT applies the LT predicate on the "end_at" field.
func EndAtLT(v time.Time) predicate.Appointment {
	return predicate.Appointment(sql.FieldLT(FieldEndAt, v))
}

// EndAtLTE applies the LTE predicate on the "end_at" field.
func EndAtLTE(v time.Time) predicate.Appointment {
	return predicate.Appointment(sql.FieldLTE(FieldEndAt, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Appointment {
	return predicate.Appointment(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Appointment {
	return predicate.Appointment(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Appointment {
	return predicate.Appointment(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Appointment {
	return predicate.Appointment(sql.FieldNotIn(FieldStatus, vs...))
}

// HasMatch applies the HasEdge predicate on the "match" edge.
func HasMatch() predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, MatchTable, MatchColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMatchWith applies the HasEdge predicate on the "match" edge with a given conditions (other predicates).
func HasMatchWith(preds ...predicate.Match) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MatchInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, MatchTable, MatchColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPayment applies the HasEdge predicate on the "payment" edge.
func HasPayment() predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, PaymentTable, PaymentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPaymentWith applies the HasEdge predicate on the "payment" edge with a given conditions (other predicates).
func HasPaymentWith(preds ...predicate.Payment) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PaymentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, PaymentTable, PaymentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Appointment) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Appointment) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Appointment) predicate.Appointment {
	return predicate.Appointment(func(s *sql.Selector) {
		p(s.Not())
	})
}
