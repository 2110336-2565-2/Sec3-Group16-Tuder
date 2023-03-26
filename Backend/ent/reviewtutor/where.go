// Code generated by ent, DO NOT EDIT.

package reviewtutor

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldLTE(FieldID, id))
}

// Score applies equality check predicate on the "score" field. It's identical to ScoreEQ.
func Score(v float32) predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldEQ(FieldScore, v))
}

// ReviewMsg applies equality check predicate on the "review_msg" field. It's identical to ReviewMsgEQ.
func ReviewMsg(v string) predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldEQ(FieldReviewMsg, v))
}

// ScoreEQ applies the EQ predicate on the "score" field.
func ScoreEQ(v float32) predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldEQ(FieldScore, v))
}

// ScoreNEQ applies the NEQ predicate on the "score" field.
func ScoreNEQ(v float32) predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldNEQ(FieldScore, v))
}

// ScoreIn applies the In predicate on the "score" field.
func ScoreIn(vs ...float32) predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldIn(FieldScore, vs...))
}

// ScoreNotIn applies the NotIn predicate on the "score" field.
func ScoreNotIn(vs ...float32) predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldNotIn(FieldScore, vs...))
}

// ScoreGT applies the GT predicate on the "score" field.
func ScoreGT(v float32) predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldGT(FieldScore, v))
}

// ScoreGTE applies the GTE predicate on the "score" field.
func ScoreGTE(v float32) predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldGTE(FieldScore, v))
}

// ScoreLT applies the LT predicate on the "score" field.
func ScoreLT(v float32) predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldLT(FieldScore, v))
}

// ScoreLTE applies the LTE predicate on the "score" field.
func ScoreLTE(v float32) predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldLTE(FieldScore, v))
}

// ScoreIsNil applies the IsNil predicate on the "score" field.
func ScoreIsNil() predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldIsNull(FieldScore))
}

// ScoreNotNil applies the NotNil predicate on the "score" field.
func ScoreNotNil() predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldNotNull(FieldScore))
}

// ReviewMsgEQ applies the EQ predicate on the "review_msg" field.
func ReviewMsgEQ(v string) predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldEQ(FieldReviewMsg, v))
}

// ReviewMsgNEQ applies the NEQ predicate on the "review_msg" field.
func ReviewMsgNEQ(v string) predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldNEQ(FieldReviewMsg, v))
}

// ReviewMsgIn applies the In predicate on the "review_msg" field.
func ReviewMsgIn(vs ...string) predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldIn(FieldReviewMsg, vs...))
}

// ReviewMsgNotIn applies the NotIn predicate on the "review_msg" field.
func ReviewMsgNotIn(vs ...string) predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldNotIn(FieldReviewMsg, vs...))
}

// ReviewMsgGT applies the GT predicate on the "review_msg" field.
func ReviewMsgGT(v string) predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldGT(FieldReviewMsg, v))
}

// ReviewMsgGTE applies the GTE predicate on the "review_msg" field.
func ReviewMsgGTE(v string) predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldGTE(FieldReviewMsg, v))
}

// ReviewMsgLT applies the LT predicate on the "review_msg" field.
func ReviewMsgLT(v string) predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldLT(FieldReviewMsg, v))
}

// ReviewMsgLTE applies the LTE predicate on the "review_msg" field.
func ReviewMsgLTE(v string) predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldLTE(FieldReviewMsg, v))
}

// ReviewMsgContains applies the Contains predicate on the "review_msg" field.
func ReviewMsgContains(v string) predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldContains(FieldReviewMsg, v))
}

// ReviewMsgHasPrefix applies the HasPrefix predicate on the "review_msg" field.
func ReviewMsgHasPrefix(v string) predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldHasPrefix(FieldReviewMsg, v))
}

// ReviewMsgHasSuffix applies the HasSuffix predicate on the "review_msg" field.
func ReviewMsgHasSuffix(v string) predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldHasSuffix(FieldReviewMsg, v))
}

// ReviewMsgIsNil applies the IsNil predicate on the "review_msg" field.
func ReviewMsgIsNil() predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldIsNull(FieldReviewMsg))
}

// ReviewMsgNotNil applies the NotNil predicate on the "review_msg" field.
func ReviewMsgNotNil() predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldNotNull(FieldReviewMsg))
}

// ReviewMsgEqualFold applies the EqualFold predicate on the "review_msg" field.
func ReviewMsgEqualFold(v string) predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldEqualFold(FieldReviewMsg, v))
}

// ReviewMsgContainsFold applies the ContainsFold predicate on the "review_msg" field.
func ReviewMsgContainsFold(v string) predicate.ReviewTutor {
	return predicate.ReviewTutor(sql.FieldContainsFold(FieldReviewMsg, v))
}

// HasTutor applies the HasEdge predicate on the "tutor" edge.
func HasTutor() predicate.ReviewTutor {
	return predicate.ReviewTutor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, TutorTable, TutorPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTutorWith applies the HasEdge predicate on the "tutor" edge with a given conditions (other predicates).
func HasTutorWith(preds ...predicate.Tutor) predicate.ReviewTutor {
	return predicate.ReviewTutor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TutorInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, TutorTable, TutorPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStudent applies the HasEdge predicate on the "student" edge.
func HasStudent() predicate.ReviewTutor {
	return predicate.ReviewTutor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, StudentTable, StudentPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStudentWith applies the HasEdge predicate on the "student" edge with a given conditions (other predicates).
func HasStudentWith(preds ...predicate.Student) predicate.ReviewTutor {
	return predicate.ReviewTutor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StudentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, StudentTable, StudentPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ReviewTutor) predicate.ReviewTutor {
	return predicate.ReviewTutor(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ReviewTutor) predicate.ReviewTutor {
	return predicate.ReviewTutor(func(s *sql.Selector) {
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
func Not(p predicate.ReviewTutor) predicate.ReviewTutor {
	return predicate.ReviewTutor(func(s *sql.Selector) {
		p(s.Not())
	})
}
