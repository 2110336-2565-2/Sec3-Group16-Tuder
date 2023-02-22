// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/class"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/course"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/issuereport"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/payment"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/paymenthistory"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/reviewcourse"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/reviewtutor"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/schedule"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/schema"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/student"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/tutor"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/user"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	classFields := schema.Class{}.Fields()
	_ = classFields
	// classDescReviewAvaliable is the schema descriptor for review_avaliable field.
	classDescReviewAvaliable := classFields[1].Descriptor()
	// class.DefaultReviewAvaliable holds the default value on creation for the review_avaliable field.
	class.DefaultReviewAvaliable = classDescReviewAvaliable.Default.(bool)
	// classDescID is the schema descriptor for id field.
	classDescID := classFields[0].Descriptor()
	// class.DefaultID holds the default value on creation for the id field.
	class.DefaultID = classDescID.Default.(func() uuid.UUID)
	courseFields := schema.Course{}.Fields()
	_ = courseFields
	// courseDescTitle is the schema descriptor for title field.
	courseDescTitle := courseFields[1].Descriptor()
	// course.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	course.TitleValidator = courseDescTitle.Validators[0].(func(string) error)
	// courseDescSubject is the schema descriptor for subject field.
	courseDescSubject := courseFields[2].Descriptor()
	// course.SubjectValidator is a validator for the "subject" field. It is called by the builders before save.
	course.SubjectValidator = courseDescSubject.Validators[0].(func(string) error)
	// courseDescTopic is the schema descriptor for topic field.
	courseDescTopic := courseFields[3].Descriptor()
	// course.TopicValidator is a validator for the "topic" field. It is called by the builders before save.
	course.TopicValidator = courseDescTopic.Validators[0].(func(string) error)
	// courseDescDescription is the schema descriptor for description field.
	courseDescDescription := courseFields[5].Descriptor()
	// course.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	course.DescriptionValidator = courseDescDescription.Validators[0].(func(string) error)
	// courseDescCourseStatus is the schema descriptor for course_status field.
	courseDescCourseStatus := courseFields[6].Descriptor()
	// course.CourseStatusValidator is a validator for the "course_status" field. It is called by the builders before save.
	course.CourseStatusValidator = courseDescCourseStatus.Validators[0].(func(string) error)
	// courseDescPricePerHour is the schema descriptor for price_per_hour field.
	courseDescPricePerHour := courseFields[7].Descriptor()
	// course.PricePerHourValidator is a validator for the "price_per_hour" field. It is called by the builders before save.
	course.PricePerHourValidator = courseDescPricePerHour.Validators[0].(func(int) error)
	// courseDescID is the schema descriptor for id field.
	courseDescID := courseFields[0].Descriptor()
	// course.DefaultID holds the default value on creation for the id field.
	course.DefaultID = courseDescID.Default.(func() uuid.UUID)
	issuereportFields := schema.IssueReport{}.Fields()
	_ = issuereportFields
	// issuereportDescTitle is the schema descriptor for title field.
	issuereportDescTitle := issuereportFields[1].Descriptor()
	// issuereport.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	issuereport.TitleValidator = issuereportDescTitle.Validators[0].(func(string) error)
	// issuereportDescDescription is the schema descriptor for description field.
	issuereportDescDescription := issuereportFields[2].Descriptor()
	// issuereport.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	issuereport.DescriptionValidator = issuereportDescDescription.Validators[0].(func(string) error)
	// issuereportDescStatus is the schema descriptor for status field.
	issuereportDescStatus := issuereportFields[4].Descriptor()
	// issuereport.StatusValidator is a validator for the "status" field. It is called by the builders before save.
	issuereport.StatusValidator = issuereportDescStatus.Validators[0].(func(string) error)
	// issuereportDescID is the schema descriptor for id field.
	issuereportDescID := issuereportFields[0].Descriptor()
	// issuereport.DefaultID holds the default value on creation for the id field.
	issuereport.DefaultID = issuereportDescID.Default.(func() uuid.UUID)
	paymentFields := schema.Payment{}.Fields()
	_ = paymentFields
	// paymentDescID is the schema descriptor for id field.
	paymentDescID := paymentFields[0].Descriptor()
	// payment.DefaultID holds the default value on creation for the id field.
	payment.DefaultID = paymentDescID.Default.(func() uuid.UUID)
	paymenthistoryFields := schema.PaymentHistory{}.Fields()
	_ = paymenthistoryFields
	// paymenthistoryDescAmount is the schema descriptor for amount field.
	paymenthistoryDescAmount := paymenthistoryFields[1].Descriptor()
	// paymenthistory.AmountValidator is a validator for the "amount" field. It is called by the builders before save.
	paymenthistory.AmountValidator = paymenthistoryDescAmount.Validators[0].(func(float64) error)
	// paymenthistoryDescType is the schema descriptor for type field.
	paymenthistoryDescType := paymenthistoryFields[2].Descriptor()
	// paymenthistory.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	paymenthistory.TypeValidator = paymenthistoryDescType.Validators[0].(func(string) error)
	// paymenthistoryDescID is the schema descriptor for id field.
	paymenthistoryDescID := paymenthistoryFields[0].Descriptor()
	// paymenthistory.DefaultID holds the default value on creation for the id field.
	paymenthistory.DefaultID = paymenthistoryDescID.Default.(func() uuid.UUID)
	reviewcourseMixin := schema.ReviewCourse{}.Mixin()
	reviewcourseMixinFields0 := reviewcourseMixin[0].Fields()
	_ = reviewcourseMixinFields0
	reviewcourseFields := schema.ReviewCourse{}.Fields()
	_ = reviewcourseFields
	// reviewcourseDescScore is the schema descriptor for score field.
	reviewcourseDescScore := reviewcourseMixinFields0[0].Descriptor()
	// reviewcourse.ScoreValidator is a validator for the "score" field. It is called by the builders before save.
	reviewcourse.ScoreValidator = reviewcourseDescScore.Validators[0].(func(float32) error)
	reviewtutorMixin := schema.ReviewTutor{}.Mixin()
	reviewtutorMixinFields0 := reviewtutorMixin[0].Fields()
	_ = reviewtutorMixinFields0
	reviewtutorFields := schema.ReviewTutor{}.Fields()
	_ = reviewtutorFields
	// reviewtutorDescScore is the schema descriptor for score field.
	reviewtutorDescScore := reviewtutorMixinFields0[0].Descriptor()
	// reviewtutor.ScoreValidator is a validator for the "score" field. It is called by the builders before save.
	reviewtutor.ScoreValidator = reviewtutorDescScore.Validators[0].(func(float32) error)
	scheduleFields := schema.Schedule{}.Fields()
	_ = scheduleFields
	// scheduleDescID is the schema descriptor for id field.
	scheduleDescID := scheduleFields[0].Descriptor()
	// schedule.DefaultID holds the default value on creation for the id field.
	schedule.DefaultID = scheduleDescID.Default.(func() uuid.UUID)
	studentFields := schema.Student{}.Fields()
	_ = studentFields
	// studentDescID is the schema descriptor for id field.
	studentDescID := studentFields[0].Descriptor()
	// student.DefaultID holds the default value on creation for the id field.
	student.DefaultID = studentDescID.Default.(func() uuid.UUID)
	tutorFields := schema.Tutor{}.Fields()
	_ = tutorFields
	// tutorDescCitizenID is the schema descriptor for citizen_id field.
	tutorDescCitizenID := tutorFields[3].Descriptor()
	// tutor.CitizenIDValidator is a validator for the "citizen_id" field. It is called by the builders before save.
	tutor.CitizenIDValidator = tutorDescCitizenID.Validators[0].(func(string) error)
	// tutorDescID is the schema descriptor for id field.
	tutorDescID := tutorFields[0].Descriptor()
	// tutor.DefaultID holds the default value on creation for the id field.
	tutor.DefaultID = tutorDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[1].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[3].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescFirstName is the schema descriptor for first_name field.
	userDescFirstName := userFields[4].Descriptor()
	// user.FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	user.FirstNameValidator = userDescFirstName.Validators[0].(func(string) error)
	// userDescLastName is the schema descriptor for last_name field.
	userDescLastName := userFields[5].Descriptor()
	// user.LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	user.LastNameValidator = userDescLastName.Validators[0].(func(string) error)
	// userDescAddress is the schema descriptor for address field.
	userDescAddress := userFields[6].Descriptor()
	// user.AddressValidator is a validator for the "address" field. It is called by the builders before save.
	user.AddressValidator = userDescAddress.Validators[0].(func(string) error)
	// userDescPhone is the schema descriptor for phone field.
	userDescPhone := userFields[7].Descriptor()
	// user.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	user.PhoneValidator = userDescPhone.Validators[0].(func(string) error)
	// userDescGender is the schema descriptor for gender field.
	userDescGender := userFields[9].Descriptor()
	// user.GenderValidator is a validator for the "gender" field. It is called by the builders before save.
	user.GenderValidator = userDescGender.Validators[0].(func(string) error)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
