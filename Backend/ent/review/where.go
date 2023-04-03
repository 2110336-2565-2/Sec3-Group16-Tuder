// Code generated by ent, DO NOT EDIT.

package review

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Review {
	return predicate.Review(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Review {
	return predicate.Review(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Review {
	return predicate.Review(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Review {
	return predicate.Review(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Review {
	return predicate.Review(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Review {
	return predicate.Review(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Review {
	return predicate.Review(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Review {
	return predicate.Review(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Review {
	return predicate.Review(sql.FieldLTE(FieldID, id))
}

// Score applies equality check predicate on the "score" field. It's identical to ScoreEQ.
func Score(v int8) predicate.Review {
	return predicate.Review(sql.FieldEQ(FieldScore, v))
}

// ReviewMsg applies equality check predicate on the "review_msg" field. It's identical to ReviewMsgEQ.
func ReviewMsg(v string) predicate.Review {
	return predicate.Review(sql.FieldEQ(FieldReviewMsg, v))
}

// ReviewTimeAt applies equality check predicate on the "review_time_at" field. It's identical to ReviewTimeAtEQ.
func ReviewTimeAt(v time.Time) predicate.Review {
	return predicate.Review(sql.FieldEQ(FieldReviewTimeAt, v))
}

// ScoreEQ applies the EQ predicate on the "score" field.
func ScoreEQ(v int8) predicate.Review {
	return predicate.Review(sql.FieldEQ(FieldScore, v))
}

// ScoreNEQ applies the NEQ predicate on the "score" field.
func ScoreNEQ(v int8) predicate.Review {
	return predicate.Review(sql.FieldNEQ(FieldScore, v))
}

// ScoreIn applies the In predicate on the "score" field.
func ScoreIn(vs ...int8) predicate.Review {
	return predicate.Review(sql.FieldIn(FieldScore, vs...))
}

// ScoreNotIn applies the NotIn predicate on the "score" field.
func ScoreNotIn(vs ...int8) predicate.Review {
	return predicate.Review(sql.FieldNotIn(FieldScore, vs...))
}

// ScoreGT applies the GT predicate on the "score" field.
func ScoreGT(v int8) predicate.Review {
	return predicate.Review(sql.FieldGT(FieldScore, v))
}

// ScoreGTE applies the GTE predicate on the "score" field.
func ScoreGTE(v int8) predicate.Review {
	return predicate.Review(sql.FieldGTE(FieldScore, v))
}

// ScoreLT applies the LT predicate on the "score" field.
func ScoreLT(v int8) predicate.Review {
	return predicate.Review(sql.FieldLT(FieldScore, v))
}

// ScoreLTE applies the LTE predicate on the "score" field.
func ScoreLTE(v int8) predicate.Review {
	return predicate.Review(sql.FieldLTE(FieldScore, v))
}

// ScoreIsNil applies the IsNil predicate on the "score" field.
func ScoreIsNil() predicate.Review {
	return predicate.Review(sql.FieldIsNull(FieldScore))
}

// ScoreNotNil applies the NotNil predicate on the "score" field.
func ScoreNotNil() predicate.Review {
	return predicate.Review(sql.FieldNotNull(FieldScore))
}

// ReviewMsgEQ applies the EQ predicate on the "review_msg" field.
func ReviewMsgEQ(v string) predicate.Review {
	return predicate.Review(sql.FieldEQ(FieldReviewMsg, v))
}

// ReviewMsgNEQ applies the NEQ predicate on the "review_msg" field.
func ReviewMsgNEQ(v string) predicate.Review {
	return predicate.Review(sql.FieldNEQ(FieldReviewMsg, v))
}

// ReviewMsgIn applies the In predicate on the "review_msg" field.
func ReviewMsgIn(vs ...string) predicate.Review {
	return predicate.Review(sql.FieldIn(FieldReviewMsg, vs...))
}

// ReviewMsgNotIn applies the NotIn predicate on the "review_msg" field.
func ReviewMsgNotIn(vs ...string) predicate.Review {
	return predicate.Review(sql.FieldNotIn(FieldReviewMsg, vs...))
}

// ReviewMsgGT applies the GT predicate on the "review_msg" field.
func ReviewMsgGT(v string) predicate.Review {
	return predicate.Review(sql.FieldGT(FieldReviewMsg, v))
}

// ReviewMsgGTE applies the GTE predicate on the "review_msg" field.
func ReviewMsgGTE(v string) predicate.Review {
	return predicate.Review(sql.FieldGTE(FieldReviewMsg, v))
}

// ReviewMsgLT applies the LT predicate on the "review_msg" field.
func ReviewMsgLT(v string) predicate.Review {
	return predicate.Review(sql.FieldLT(FieldReviewMsg, v))
}

// ReviewMsgLTE applies the LTE predicate on the "review_msg" field.
func ReviewMsgLTE(v string) predicate.Review {
	return predicate.Review(sql.FieldLTE(FieldReviewMsg, v))
}

// ReviewMsgContains applies the Contains predicate on the "review_msg" field.
func ReviewMsgContains(v string) predicate.Review {
	return predicate.Review(sql.FieldContains(FieldReviewMsg, v))
}

// ReviewMsgHasPrefix applies the HasPrefix predicate on the "review_msg" field.
func ReviewMsgHasPrefix(v string) predicate.Review {
	return predicate.Review(sql.FieldHasPrefix(FieldReviewMsg, v))
}

// ReviewMsgHasSuffix applies the HasSuffix predicate on the "review_msg" field.
func ReviewMsgHasSuffix(v string) predicate.Review {
	return predicate.Review(sql.FieldHasSuffix(FieldReviewMsg, v))
}

// ReviewMsgIsNil applies the IsNil predicate on the "review_msg" field.
func ReviewMsgIsNil() predicate.Review {
	return predicate.Review(sql.FieldIsNull(FieldReviewMsg))
}

// ReviewMsgNotNil applies the NotNil predicate on the "review_msg" field.
func ReviewMsgNotNil() predicate.Review {
	return predicate.Review(sql.FieldNotNull(FieldReviewMsg))
}

// ReviewMsgEqualFold applies the EqualFold predicate on the "review_msg" field.
func ReviewMsgEqualFold(v string) predicate.Review {
	return predicate.Review(sql.FieldEqualFold(FieldReviewMsg, v))
}

// ReviewMsgContainsFold applies the ContainsFold predicate on the "review_msg" field.
func ReviewMsgContainsFold(v string) predicate.Review {
	return predicate.Review(sql.FieldContainsFold(FieldReviewMsg, v))
}

// ReviewTimeAtEQ applies the EQ predicate on the "review_time_at" field.
func ReviewTimeAtEQ(v time.Time) predicate.Review {
	return predicate.Review(sql.FieldEQ(FieldReviewTimeAt, v))
}

// ReviewTimeAtNEQ applies the NEQ predicate on the "review_time_at" field.
func ReviewTimeAtNEQ(v time.Time) predicate.Review {
	return predicate.Review(sql.FieldNEQ(FieldReviewTimeAt, v))
}

// ReviewTimeAtIn applies the In predicate on the "review_time_at" field.
func ReviewTimeAtIn(vs ...time.Time) predicate.Review {
	return predicate.Review(sql.FieldIn(FieldReviewTimeAt, vs...))
}

// ReviewTimeAtNotIn applies the NotIn predicate on the "review_time_at" field.
func ReviewTimeAtNotIn(vs ...time.Time) predicate.Review {
	return predicate.Review(sql.FieldNotIn(FieldReviewTimeAt, vs...))
}

// ReviewTimeAtGT applies the GT predicate on the "review_time_at" field.
func ReviewTimeAtGT(v time.Time) predicate.Review {
	return predicate.Review(sql.FieldGT(FieldReviewTimeAt, v))
}

// ReviewTimeAtGTE applies the GTE predicate on the "review_time_at" field.
func ReviewTimeAtGTE(v time.Time) predicate.Review {
	return predicate.Review(sql.FieldGTE(FieldReviewTimeAt, v))
}

// ReviewTimeAtLT applies the LT predicate on the "review_time_at" field.
func ReviewTimeAtLT(v time.Time) predicate.Review {
	return predicate.Review(sql.FieldLT(FieldReviewTimeAt, v))
}

// ReviewTimeAtLTE applies the LTE predicate on the "review_time_at" field.
func ReviewTimeAtLTE(v time.Time) predicate.Review {
	return predicate.Review(sql.FieldLTE(FieldReviewTimeAt, v))
}

// HasCourse applies the HasEdge predicate on the "course" edge.
func HasCourse() predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, CourseTable, CoursePrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCourseWith applies the HasEdge predicate on the "course" edge with a given conditions (other predicates).
func HasCourseWith(preds ...predicate.Course) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CourseInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, CourseTable, CoursePrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStudent applies the HasEdge predicate on the "student" edge.
func HasStudent() predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, StudentTable, StudentPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStudentWith applies the HasEdge predicate on the "student" edge with a given conditions (other predicates).
func HasStudentWith(preds ...predicate.Student) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
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
func And(predicates ...predicate.Review) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Review) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
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
func Not(p predicate.Review) predicate.Review {
	return predicate.Review(func(s *sql.Selector) {
		p(s.Not())
	})
}
