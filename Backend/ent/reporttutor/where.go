// Code generated by ent, DO NOT EDIT.

package reporttutor

import (
	"entgo.io/ent/dialect/sql"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldLTE(FieldID, id))
}

// Score applies equality check predicate on the "score" field. It's identical to ScoreEQ.
func Score(v float32) predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldEQ(FieldScore, v))
}

// ReviewMsg applies equality check predicate on the "review_msg" field. It's identical to ReviewMsgEQ.
func ReviewMsg(v string) predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldEQ(FieldReviewMsg, v))
}

// ScoreEQ applies the EQ predicate on the "score" field.
func ScoreEQ(v float32) predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldEQ(FieldScore, v))
}

// ScoreNEQ applies the NEQ predicate on the "score" field.
func ScoreNEQ(v float32) predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldNEQ(FieldScore, v))
}

// ScoreIn applies the In predicate on the "score" field.
func ScoreIn(vs ...float32) predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldIn(FieldScore, vs...))
}

// ScoreNotIn applies the NotIn predicate on the "score" field.
func ScoreNotIn(vs ...float32) predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldNotIn(FieldScore, vs...))
}

// ScoreGT applies the GT predicate on the "score" field.
func ScoreGT(v float32) predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldGT(FieldScore, v))
}

// ScoreGTE applies the GTE predicate on the "score" field.
func ScoreGTE(v float32) predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldGTE(FieldScore, v))
}

// ScoreLT applies the LT predicate on the "score" field.
func ScoreLT(v float32) predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldLT(FieldScore, v))
}

// ScoreLTE applies the LTE predicate on the "score" field.
func ScoreLTE(v float32) predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldLTE(FieldScore, v))
}

// ScoreIsNil applies the IsNil predicate on the "score" field.
func ScoreIsNil() predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldIsNull(FieldScore))
}

// ScoreNotNil applies the NotNil predicate on the "score" field.
func ScoreNotNil() predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldNotNull(FieldScore))
}

// ReviewMsgEQ applies the EQ predicate on the "review_msg" field.
func ReviewMsgEQ(v string) predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldEQ(FieldReviewMsg, v))
}

// ReviewMsgNEQ applies the NEQ predicate on the "review_msg" field.
func ReviewMsgNEQ(v string) predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldNEQ(FieldReviewMsg, v))
}

// ReviewMsgIn applies the In predicate on the "review_msg" field.
func ReviewMsgIn(vs ...string) predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldIn(FieldReviewMsg, vs...))
}

// ReviewMsgNotIn applies the NotIn predicate on the "review_msg" field.
func ReviewMsgNotIn(vs ...string) predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldNotIn(FieldReviewMsg, vs...))
}

// ReviewMsgGT applies the GT predicate on the "review_msg" field.
func ReviewMsgGT(v string) predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldGT(FieldReviewMsg, v))
}

// ReviewMsgGTE applies the GTE predicate on the "review_msg" field.
func ReviewMsgGTE(v string) predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldGTE(FieldReviewMsg, v))
}

// ReviewMsgLT applies the LT predicate on the "review_msg" field.
func ReviewMsgLT(v string) predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldLT(FieldReviewMsg, v))
}

// ReviewMsgLTE applies the LTE predicate on the "review_msg" field.
func ReviewMsgLTE(v string) predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldLTE(FieldReviewMsg, v))
}

// ReviewMsgContains applies the Contains predicate on the "review_msg" field.
func ReviewMsgContains(v string) predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldContains(FieldReviewMsg, v))
}

// ReviewMsgHasPrefix applies the HasPrefix predicate on the "review_msg" field.
func ReviewMsgHasPrefix(v string) predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldHasPrefix(FieldReviewMsg, v))
}

// ReviewMsgHasSuffix applies the HasSuffix predicate on the "review_msg" field.
func ReviewMsgHasSuffix(v string) predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldHasSuffix(FieldReviewMsg, v))
}

// ReviewMsgIsNil applies the IsNil predicate on the "review_msg" field.
func ReviewMsgIsNil() predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldIsNull(FieldReviewMsg))
}

// ReviewMsgNotNil applies the NotNil predicate on the "review_msg" field.
func ReviewMsgNotNil() predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldNotNull(FieldReviewMsg))
}

// ReviewMsgEqualFold applies the EqualFold predicate on the "review_msg" field.
func ReviewMsgEqualFold(v string) predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldEqualFold(FieldReviewMsg, v))
}

// ReviewMsgContainsFold applies the ContainsFold predicate on the "review_msg" field.
func ReviewMsgContainsFold(v string) predicate.ReportTutor {
	return predicate.ReportTutor(sql.FieldContainsFold(FieldReviewMsg, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ReportTutor) predicate.ReportTutor {
	return predicate.ReportTutor(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ReportTutor) predicate.ReportTutor {
	return predicate.ReportTutor(func(s *sql.Selector) {
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
func Not(p predicate.ReportTutor) predicate.ReportTutor {
	return predicate.ReportTutor(func(s *sql.Selector) {
		p(s.Not())
	})
}
