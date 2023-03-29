// Code generated by ent, DO NOT EDIT.

package course

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Course {
	return predicate.Course(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Course {
	return predicate.Course(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Course {
	return predicate.Course(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Course {
	return predicate.Course(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Course {
	return predicate.Course(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Course {
	return predicate.Course(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Course {
	return predicate.Course(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Course {
	return predicate.Course(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Course {
	return predicate.Course(sql.FieldLTE(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Course {
	return predicate.Course(sql.FieldEQ(FieldTitle, v))
}

// Subject applies equality check predicate on the "subject" field. It's identical to SubjectEQ.
func Subject(v string) predicate.Course {
	return predicate.Course(sql.FieldEQ(FieldSubject, v))
}

// Topic applies equality check predicate on the "topic" field. It's identical to TopicEQ.
func Topic(v string) predicate.Course {
	return predicate.Course(sql.FieldEQ(FieldTopic, v))
}

// EstimatedTime applies equality check predicate on the "estimated_time" field. It's identical to EstimatedTimeEQ.
func EstimatedTime(v int) predicate.Course {
	return predicate.Course(sql.FieldEQ(FieldEstimatedTime, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Course {
	return predicate.Course(sql.FieldEQ(FieldDescription, v))
}

// PricePerHour applies equality check predicate on the "price_per_hour" field. It's identical to PricePerHourEQ.
func PricePerHour(v int) predicate.Course {
	return predicate.Course(sql.FieldEQ(FieldPricePerHour, v))
}

// CoursePictureURL applies equality check predicate on the "course_picture_url" field. It's identical to CoursePictureURLEQ.
func CoursePictureURL(v string) predicate.Course {
	return predicate.Course(sql.FieldEQ(FieldCoursePictureURL, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Course {
	return predicate.Course(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Course {
	return predicate.Course(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Course {
	return predicate.Course(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Course {
	return predicate.Course(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Course {
	return predicate.Course(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Course {
	return predicate.Course(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Course {
	return predicate.Course(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Course {
	return predicate.Course(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Course {
	return predicate.Course(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Course {
	return predicate.Course(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Course {
	return predicate.Course(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Course {
	return predicate.Course(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Course {
	return predicate.Course(sql.FieldContainsFold(FieldTitle, v))
}

// SubjectEQ applies the EQ predicate on the "subject" field.
func SubjectEQ(v string) predicate.Course {
	return predicate.Course(sql.FieldEQ(FieldSubject, v))
}

// SubjectNEQ applies the NEQ predicate on the "subject" field.
func SubjectNEQ(v string) predicate.Course {
	return predicate.Course(sql.FieldNEQ(FieldSubject, v))
}

// SubjectIn applies the In predicate on the "subject" field.
func SubjectIn(vs ...string) predicate.Course {
	return predicate.Course(sql.FieldIn(FieldSubject, vs...))
}

// SubjectNotIn applies the NotIn predicate on the "subject" field.
func SubjectNotIn(vs ...string) predicate.Course {
	return predicate.Course(sql.FieldNotIn(FieldSubject, vs...))
}

// SubjectGT applies the GT predicate on the "subject" field.
func SubjectGT(v string) predicate.Course {
	return predicate.Course(sql.FieldGT(FieldSubject, v))
}

// SubjectGTE applies the GTE predicate on the "subject" field.
func SubjectGTE(v string) predicate.Course {
	return predicate.Course(sql.FieldGTE(FieldSubject, v))
}

// SubjectLT applies the LT predicate on the "subject" field.
func SubjectLT(v string) predicate.Course {
	return predicate.Course(sql.FieldLT(FieldSubject, v))
}

// SubjectLTE applies the LTE predicate on the "subject" field.
func SubjectLTE(v string) predicate.Course {
	return predicate.Course(sql.FieldLTE(FieldSubject, v))
}

// SubjectContains applies the Contains predicate on the "subject" field.
func SubjectContains(v string) predicate.Course {
	return predicate.Course(sql.FieldContains(FieldSubject, v))
}

// SubjectHasPrefix applies the HasPrefix predicate on the "subject" field.
func SubjectHasPrefix(v string) predicate.Course {
	return predicate.Course(sql.FieldHasPrefix(FieldSubject, v))
}

// SubjectHasSuffix applies the HasSuffix predicate on the "subject" field.
func SubjectHasSuffix(v string) predicate.Course {
	return predicate.Course(sql.FieldHasSuffix(FieldSubject, v))
}

// SubjectEqualFold applies the EqualFold predicate on the "subject" field.
func SubjectEqualFold(v string) predicate.Course {
	return predicate.Course(sql.FieldEqualFold(FieldSubject, v))
}

// SubjectContainsFold applies the ContainsFold predicate on the "subject" field.
func SubjectContainsFold(v string) predicate.Course {
	return predicate.Course(sql.FieldContainsFold(FieldSubject, v))
}

// TopicEQ applies the EQ predicate on the "topic" field.
func TopicEQ(v string) predicate.Course {
	return predicate.Course(sql.FieldEQ(FieldTopic, v))
}

// TopicNEQ applies the NEQ predicate on the "topic" field.
func TopicNEQ(v string) predicate.Course {
	return predicate.Course(sql.FieldNEQ(FieldTopic, v))
}

// TopicIn applies the In predicate on the "topic" field.
func TopicIn(vs ...string) predicate.Course {
	return predicate.Course(sql.FieldIn(FieldTopic, vs...))
}

// TopicNotIn applies the NotIn predicate on the "topic" field.
func TopicNotIn(vs ...string) predicate.Course {
	return predicate.Course(sql.FieldNotIn(FieldTopic, vs...))
}

// TopicGT applies the GT predicate on the "topic" field.
func TopicGT(v string) predicate.Course {
	return predicate.Course(sql.FieldGT(FieldTopic, v))
}

// TopicGTE applies the GTE predicate on the "topic" field.
func TopicGTE(v string) predicate.Course {
	return predicate.Course(sql.FieldGTE(FieldTopic, v))
}

// TopicLT applies the LT predicate on the "topic" field.
func TopicLT(v string) predicate.Course {
	return predicate.Course(sql.FieldLT(FieldTopic, v))
}

// TopicLTE applies the LTE predicate on the "topic" field.
func TopicLTE(v string) predicate.Course {
	return predicate.Course(sql.FieldLTE(FieldTopic, v))
}

// TopicContains applies the Contains predicate on the "topic" field.
func TopicContains(v string) predicate.Course {
	return predicate.Course(sql.FieldContains(FieldTopic, v))
}

// TopicHasPrefix applies the HasPrefix predicate on the "topic" field.
func TopicHasPrefix(v string) predicate.Course {
	return predicate.Course(sql.FieldHasPrefix(FieldTopic, v))
}

// TopicHasSuffix applies the HasSuffix predicate on the "topic" field.
func TopicHasSuffix(v string) predicate.Course {
	return predicate.Course(sql.FieldHasSuffix(FieldTopic, v))
}

// TopicEqualFold applies the EqualFold predicate on the "topic" field.
func TopicEqualFold(v string) predicate.Course {
	return predicate.Course(sql.FieldEqualFold(FieldTopic, v))
}

// TopicContainsFold applies the ContainsFold predicate on the "topic" field.
func TopicContainsFold(v string) predicate.Course {
	return predicate.Course(sql.FieldContainsFold(FieldTopic, v))
}

// EstimatedTimeEQ applies the EQ predicate on the "estimated_time" field.
func EstimatedTimeEQ(v int) predicate.Course {
	return predicate.Course(sql.FieldEQ(FieldEstimatedTime, v))
}

// EstimatedTimeNEQ applies the NEQ predicate on the "estimated_time" field.
func EstimatedTimeNEQ(v int) predicate.Course {
	return predicate.Course(sql.FieldNEQ(FieldEstimatedTime, v))
}

// EstimatedTimeIn applies the In predicate on the "estimated_time" field.
func EstimatedTimeIn(vs ...int) predicate.Course {
	return predicate.Course(sql.FieldIn(FieldEstimatedTime, vs...))
}

// EstimatedTimeNotIn applies the NotIn predicate on the "estimated_time" field.
func EstimatedTimeNotIn(vs ...int) predicate.Course {
	return predicate.Course(sql.FieldNotIn(FieldEstimatedTime, vs...))
}

// EstimatedTimeGT applies the GT predicate on the "estimated_time" field.
func EstimatedTimeGT(v int) predicate.Course {
	return predicate.Course(sql.FieldGT(FieldEstimatedTime, v))
}

// EstimatedTimeGTE applies the GTE predicate on the "estimated_time" field.
func EstimatedTimeGTE(v int) predicate.Course {
	return predicate.Course(sql.FieldGTE(FieldEstimatedTime, v))
}

// EstimatedTimeLT applies the LT predicate on the "estimated_time" field.
func EstimatedTimeLT(v int) predicate.Course {
	return predicate.Course(sql.FieldLT(FieldEstimatedTime, v))
}

// EstimatedTimeLTE applies the LTE predicate on the "estimated_time" field.
func EstimatedTimeLTE(v int) predicate.Course {
	return predicate.Course(sql.FieldLTE(FieldEstimatedTime, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Course {
	return predicate.Course(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Course {
	return predicate.Course(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Course {
	return predicate.Course(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Course {
	return predicate.Course(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Course {
	return predicate.Course(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Course {
	return predicate.Course(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Course {
	return predicate.Course(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Course {
	return predicate.Course(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Course {
	return predicate.Course(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Course {
	return predicate.Course(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Course {
	return predicate.Course(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Course {
	return predicate.Course(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Course {
	return predicate.Course(sql.FieldContainsFold(FieldDescription, v))
}

// PricePerHourEQ applies the EQ predicate on the "price_per_hour" field.
func PricePerHourEQ(v int) predicate.Course {
	return predicate.Course(sql.FieldEQ(FieldPricePerHour, v))
}

// PricePerHourNEQ applies the NEQ predicate on the "price_per_hour" field.
func PricePerHourNEQ(v int) predicate.Course {
	return predicate.Course(sql.FieldNEQ(FieldPricePerHour, v))
}

// PricePerHourIn applies the In predicate on the "price_per_hour" field.
func PricePerHourIn(vs ...int) predicate.Course {
	return predicate.Course(sql.FieldIn(FieldPricePerHour, vs...))
}

// PricePerHourNotIn applies the NotIn predicate on the "price_per_hour" field.
func PricePerHourNotIn(vs ...int) predicate.Course {
	return predicate.Course(sql.FieldNotIn(FieldPricePerHour, vs...))
}

// PricePerHourGT applies the GT predicate on the "price_per_hour" field.
func PricePerHourGT(v int) predicate.Course {
	return predicate.Course(sql.FieldGT(FieldPricePerHour, v))
}

// PricePerHourGTE applies the GTE predicate on the "price_per_hour" field.
func PricePerHourGTE(v int) predicate.Course {
	return predicate.Course(sql.FieldGTE(FieldPricePerHour, v))
}

// PricePerHourLT applies the LT predicate on the "price_per_hour" field.
func PricePerHourLT(v int) predicate.Course {
	return predicate.Course(sql.FieldLT(FieldPricePerHour, v))
}

// PricePerHourLTE applies the LTE predicate on the "price_per_hour" field.
func PricePerHourLTE(v int) predicate.Course {
	return predicate.Course(sql.FieldLTE(FieldPricePerHour, v))
}

// LevelEQ applies the EQ predicate on the "level" field.
func LevelEQ(v Level) predicate.Course {
	return predicate.Course(sql.FieldEQ(FieldLevel, v))
}

// LevelNEQ applies the NEQ predicate on the "level" field.
func LevelNEQ(v Level) predicate.Course {
	return predicate.Course(sql.FieldNEQ(FieldLevel, v))
}

// LevelIn applies the In predicate on the "level" field.
func LevelIn(vs ...Level) predicate.Course {
	return predicate.Course(sql.FieldIn(FieldLevel, vs...))
}

// LevelNotIn applies the NotIn predicate on the "level" field.
func LevelNotIn(vs ...Level) predicate.Course {
	return predicate.Course(sql.FieldNotIn(FieldLevel, vs...))
}

// LevelIsNil applies the IsNil predicate on the "level" field.
func LevelIsNil() predicate.Course {
	return predicate.Course(sql.FieldIsNull(FieldLevel))
}

// LevelNotNil applies the NotNil predicate on the "level" field.
func LevelNotNil() predicate.Course {
	return predicate.Course(sql.FieldNotNull(FieldLevel))
}

// CoursePictureURLEQ applies the EQ predicate on the "course_picture_url" field.
func CoursePictureURLEQ(v string) predicate.Course {
	return predicate.Course(sql.FieldEQ(FieldCoursePictureURL, v))
}

// CoursePictureURLNEQ applies the NEQ predicate on the "course_picture_url" field.
func CoursePictureURLNEQ(v string) predicate.Course {
	return predicate.Course(sql.FieldNEQ(FieldCoursePictureURL, v))
}

// CoursePictureURLIn applies the In predicate on the "course_picture_url" field.
func CoursePictureURLIn(vs ...string) predicate.Course {
	return predicate.Course(sql.FieldIn(FieldCoursePictureURL, vs...))
}

// CoursePictureURLNotIn applies the NotIn predicate on the "course_picture_url" field.
func CoursePictureURLNotIn(vs ...string) predicate.Course {
	return predicate.Course(sql.FieldNotIn(FieldCoursePictureURL, vs...))
}

// CoursePictureURLGT applies the GT predicate on the "course_picture_url" field.
func CoursePictureURLGT(v string) predicate.Course {
	return predicate.Course(sql.FieldGT(FieldCoursePictureURL, v))
}

// CoursePictureURLGTE applies the GTE predicate on the "course_picture_url" field.
func CoursePictureURLGTE(v string) predicate.Course {
	return predicate.Course(sql.FieldGTE(FieldCoursePictureURL, v))
}

// CoursePictureURLLT applies the LT predicate on the "course_picture_url" field.
func CoursePictureURLLT(v string) predicate.Course {
	return predicate.Course(sql.FieldLT(FieldCoursePictureURL, v))
}

// CoursePictureURLLTE applies the LTE predicate on the "course_picture_url" field.
func CoursePictureURLLTE(v string) predicate.Course {
	return predicate.Course(sql.FieldLTE(FieldCoursePictureURL, v))
}

// CoursePictureURLContains applies the Contains predicate on the "course_picture_url" field.
func CoursePictureURLContains(v string) predicate.Course {
	return predicate.Course(sql.FieldContains(FieldCoursePictureURL, v))
}

// CoursePictureURLHasPrefix applies the HasPrefix predicate on the "course_picture_url" field.
func CoursePictureURLHasPrefix(v string) predicate.Course {
	return predicate.Course(sql.FieldHasPrefix(FieldCoursePictureURL, v))
}

// CoursePictureURLHasSuffix applies the HasSuffix predicate on the "course_picture_url" field.
func CoursePictureURLHasSuffix(v string) predicate.Course {
	return predicate.Course(sql.FieldHasSuffix(FieldCoursePictureURL, v))
}

// CoursePictureURLIsNil applies the IsNil predicate on the "course_picture_url" field.
func CoursePictureURLIsNil() predicate.Course {
	return predicate.Course(sql.FieldIsNull(FieldCoursePictureURL))
}

// CoursePictureURLNotNil applies the NotNil predicate on the "course_picture_url" field.
func CoursePictureURLNotNil() predicate.Course {
	return predicate.Course(sql.FieldNotNull(FieldCoursePictureURL))
}

// CoursePictureURLEqualFold applies the EqualFold predicate on the "course_picture_url" field.
func CoursePictureURLEqualFold(v string) predicate.Course {
	return predicate.Course(sql.FieldEqualFold(FieldCoursePictureURL, v))
}

// CoursePictureURLContainsFold applies the ContainsFold predicate on the "course_picture_url" field.
func CoursePictureURLContainsFold(v string) predicate.Course {
	return predicate.Course(sql.FieldContainsFold(FieldCoursePictureURL, v))
}

// HasReviewCourse applies the HasEdge predicate on the "review_course" edge.
func HasReviewCourse() predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ReviewCourseTable, ReviewCoursePrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReviewCourseWith applies the HasEdge predicate on the "review_course" edge with a given conditions (other predicates).
func HasReviewCourseWith(preds ...predicate.ReviewCourse) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ReviewCourseInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ReviewCourseTable, ReviewCoursePrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMatch applies the HasEdge predicate on the "match" edge.
func HasMatch() predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MatchTable, MatchColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMatchWith applies the HasEdge predicate on the "match" edge with a given conditions (other predicates).
func HasMatchWith(preds ...predicate.Match) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MatchInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MatchTable, MatchColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTutor applies the HasEdge predicate on the "tutor" edge.
func HasTutor() predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TutorTable, TutorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTutorWith applies the HasEdge predicate on the "tutor" edge with a given conditions (other predicates).
func HasTutorWith(preds ...predicate.Tutor) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TutorInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TutorTable, TutorColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Course) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Course) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
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
func Not(p predicate.Course) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		p(s.Not())
	})
}
