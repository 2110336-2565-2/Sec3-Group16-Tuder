// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ClassesColumns holds the columns for the "classes" table.
	ClassesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "review_avaliable", Type: field.TypeBool, Default: true},
		{Name: "total_hour", Type: field.TypeInt},
		{Name: "success_hour", Type: field.TypeInt},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"scheduled", "completed", "cancelling", "rejected", "cancelled"}},
		{Name: "payment_history_class", Type: field.TypeUUID},
		{Name: "schedule_class", Type: field.TypeUUID},
	}
	// ClassesTable holds the schema information for the "classes" table.
	ClassesTable = &schema.Table{
		Name:       "classes",
		Columns:    ClassesColumns,
		PrimaryKey: []*schema.Column{ClassesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "classes_payment_histories_class",
				Columns:    []*schema.Column{ClassesColumns[5]},
				RefColumns: []*schema.Column{PaymentHistoriesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "classes_schedules_class",
				Columns:    []*schema.Column{ClassesColumns[6]},
				RefColumns: []*schema.Column{SchedulesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// CoursesColumns holds the columns for the "courses" table.
	CoursesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "title", Type: field.TypeString},
		{Name: "subject", Type: field.TypeString},
		{Name: "topic", Type: field.TypeString},
		{Name: "estimated_time", Type: field.TypeInt},
		{Name: "description", Type: field.TypeString},
		{Name: "price_per_hour", Type: field.TypeInt},
		{Name: "level", Type: field.TypeEnum, Nullable: true, Enums: []string{"Grade1", "Grade2", "Grade3", "Grade4", "Grade5", "Grade6", "Grade7", "Grade8", "Grade9", "Grade10", "Grade11", "Grade12"}},
		{Name: "course_picture_url", Type: field.TypeString, Nullable: true},
		{Name: "tutor_course", Type: field.TypeUUID},
	}
	// CoursesTable holds the schema information for the "courses" table.
	CoursesTable = &schema.Table{
		Name:       "courses",
		Columns:    CoursesColumns,
		PrimaryKey: []*schema.Column{CoursesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "courses_tutors_course",
				Columns:    []*schema.Column{CoursesColumns[9]},
				RefColumns: []*schema.Column{TutorsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// IssueReportsColumns holds the columns for the "issue_reports" table.
	IssueReportsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "report_date", Type: field.TypeTime},
		{Name: "status", Type: field.TypeString},
		{Name: "tutor_issue_report", Type: field.TypeUUID, Nullable: true},
		{Name: "user_issue_report", Type: field.TypeUUID},
	}
	// IssueReportsTable holds the schema information for the "issue_reports" table.
	IssueReportsTable = &schema.Table{
		Name:       "issue_reports",
		Columns:    IssueReportsColumns,
		PrimaryKey: []*schema.Column{IssueReportsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "issue_reports_tutors_issue_report",
				Columns:    []*schema.Column{IssueReportsColumns[5]},
				RefColumns: []*schema.Column{TutorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "issue_reports_users_issue_report",
				Columns:    []*schema.Column{IssueReportsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// MatchesColumns holds the columns for the "matches" table.
	MatchesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "student_match", Type: field.TypeUUID},
	}
	// MatchesTable holds the schema information for the "matches" table.
	MatchesTable = &schema.Table{
		Name:       "matches",
		Columns:    MatchesColumns,
		PrimaryKey: []*schema.Column{MatchesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "matches_students_match",
				Columns:    []*schema.Column{MatchesColumns[1]},
				RefColumns: []*schema.Column{StudentsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PaymentsColumns holds the columns for the "payments" table.
	PaymentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "qr_picture_url", Type: field.TypeString, Nullable: true},
		{Name: "user_payment", Type: field.TypeUUID},
	}
	// PaymentsTable holds the schema information for the "payments" table.
	PaymentsTable = &schema.Table{
		Name:       "payments",
		Columns:    PaymentsColumns,
		PrimaryKey: []*schema.Column{PaymentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "payments_users_payment",
				Columns:    []*schema.Column{PaymentsColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PaymentHistoriesColumns holds the columns for the "payment_histories" table.
	PaymentHistoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "payment_payment_history", Type: field.TypeUUID},
		{Name: "user_payment_history", Type: field.TypeUUID},
	}
	// PaymentHistoriesTable holds the schema information for the "payment_histories" table.
	PaymentHistoriesTable = &schema.Table{
		Name:       "payment_histories",
		Columns:    PaymentHistoriesColumns,
		PrimaryKey: []*schema.Column{PaymentHistoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "payment_histories_payments_payment_history",
				Columns:    []*schema.Column{PaymentHistoriesColumns[1]},
				RefColumns: []*schema.Column{PaymentsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "payment_histories_users_payment_history",
				Columns:    []*schema.Column{PaymentHistoriesColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ReviewCoursesColumns holds the columns for the "review_courses" table.
	ReviewCoursesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "score", Type: field.TypeFloat32, Nullable: true},
		{Name: "review_msg", Type: field.TypeString, Nullable: true},
	}
	// ReviewCoursesTable holds the schema information for the "review_courses" table.
	ReviewCoursesTable = &schema.Table{
		Name:       "review_courses",
		Columns:    ReviewCoursesColumns,
		PrimaryKey: []*schema.Column{ReviewCoursesColumns[0]},
	}
	// ReviewTutorsColumns holds the columns for the "review_tutors" table.
	ReviewTutorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "score", Type: field.TypeFloat32, Nullable: true},
		{Name: "review_msg", Type: field.TypeString, Nullable: true},
	}
	// ReviewTutorsTable holds the schema information for the "review_tutors" table.
	ReviewTutorsTable = &schema.Table{
		Name:       "review_tutors",
		Columns:    ReviewTutorsColumns,
		PrimaryKey: []*schema.Column{ReviewTutorsColumns[0]},
	}
	// SchedulesColumns holds the columns for the "schedules" table.
	SchedulesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "day_0", Type: field.TypeJSON},
		{Name: "day_1", Type: field.TypeJSON},
		{Name: "day_2", Type: field.TypeJSON},
		{Name: "day_3", Type: field.TypeJSON},
		{Name: "day_4", Type: field.TypeJSON},
		{Name: "day_5", Type: field.TypeJSON},
		{Name: "day_6", Type: field.TypeJSON},
	}
	// SchedulesTable holds the schema information for the "schedules" table.
	SchedulesTable = &schema.Table{
		Name:       "schedules",
		Columns:    SchedulesColumns,
		PrimaryKey: []*schema.Column{SchedulesColumns[0]},
	}
	// StudentsColumns holds the columns for the "students" table.
	StudentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "user_student", Type: field.TypeUUID, Unique: true},
	}
	// StudentsTable holds the schema information for the "students" table.
	StudentsTable = &schema.Table{
		Name:       "students",
		Columns:    StudentsColumns,
		PrimaryKey: []*schema.Column{StudentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "students_users_student",
				Columns:    []*schema.Column{StudentsColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TutorsColumns holds the columns for the "tutors" table.
	TutorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "omise_bank_token", Type: field.TypeString, Nullable: true},
		{Name: "citizen_id", Type: field.TypeString, Unique: true},
		{Name: "schedule_tutor", Type: field.TypeUUID},
		{Name: "user_tutor", Type: field.TypeUUID, Unique: true},
	}
	// TutorsTable holds the schema information for the "tutors" table.
	TutorsTable = &schema.Table{
		Name:       "tutors",
		Columns:    TutorsColumns,
		PrimaryKey: []*schema.Column{TutorsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tutors_schedules_tutor",
				Columns:    []*schema.Column{TutorsColumns[4]},
				RefColumns: []*schema.Column{SchedulesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "tutors_users_tutor",
				Columns:    []*schema.Column{TutorsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "first_name", Type: field.TypeString},
		{Name: "last_name", Type: field.TypeString},
		{Name: "address", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
		{Name: "birth_date", Type: field.TypeTime},
		{Name: "gender", Type: field.TypeString},
		{Name: "profile_picture_url", Type: field.TypeString, Nullable: true},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"student", "tutor", "admin"}},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_username_role",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[1], UsersColumns[11]},
			},
		},
	}
	// ClassMatchColumns holds the columns for the "class_match" table.
	ClassMatchColumns = []*schema.Column{
		{Name: "class_id", Type: field.TypeUUID},
		{Name: "match_id", Type: field.TypeUUID},
	}
	// ClassMatchTable holds the schema information for the "class_match" table.
	ClassMatchTable = &schema.Table{
		Name:       "class_match",
		Columns:    ClassMatchColumns,
		PrimaryKey: []*schema.Column{ClassMatchColumns[0], ClassMatchColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "class_match_class_id",
				Columns:    []*schema.Column{ClassMatchColumns[0]},
				RefColumns: []*schema.Column{ClassesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "class_match_match_id",
				Columns:    []*schema.Column{ClassMatchColumns[1]},
				RefColumns: []*schema.Column{MatchesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// CourseReviewCourseColumns holds the columns for the "course_review_course" table.
	CourseReviewCourseColumns = []*schema.Column{
		{Name: "course_id", Type: field.TypeUUID},
		{Name: "review_course_id", Type: field.TypeInt},
	}
	// CourseReviewCourseTable holds the schema information for the "course_review_course" table.
	CourseReviewCourseTable = &schema.Table{
		Name:       "course_review_course",
		Columns:    CourseReviewCourseColumns,
		PrimaryKey: []*schema.Column{CourseReviewCourseColumns[0], CourseReviewCourseColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "course_review_course_course_id",
				Columns:    []*schema.Column{CourseReviewCourseColumns[0]},
				RefColumns: []*schema.Column{CoursesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "course_review_course_review_course_id",
				Columns:    []*schema.Column{CourseReviewCourseColumns[1]},
				RefColumns: []*schema.Column{ReviewCoursesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// CourseMatchColumns holds the columns for the "course_match" table.
	CourseMatchColumns = []*schema.Column{
		{Name: "course_id", Type: field.TypeUUID},
		{Name: "match_id", Type: field.TypeUUID},
	}
	// CourseMatchTable holds the schema information for the "course_match" table.
	CourseMatchTable = &schema.Table{
		Name:       "course_match",
		Columns:    CourseMatchColumns,
		PrimaryKey: []*schema.Column{CourseMatchColumns[0], CourseMatchColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "course_match_course_id",
				Columns:    []*schema.Column{CourseMatchColumns[0]},
				RefColumns: []*schema.Column{CoursesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "course_match_match_id",
				Columns:    []*schema.Column{CourseMatchColumns[1]},
				RefColumns: []*schema.Column{MatchesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// StudentReviewCourseColumns holds the columns for the "student_review_course" table.
	StudentReviewCourseColumns = []*schema.Column{
		{Name: "student_id", Type: field.TypeUUID},
		{Name: "review_course_id", Type: field.TypeInt},
	}
	// StudentReviewCourseTable holds the schema information for the "student_review_course" table.
	StudentReviewCourseTable = &schema.Table{
		Name:       "student_review_course",
		Columns:    StudentReviewCourseColumns,
		PrimaryKey: []*schema.Column{StudentReviewCourseColumns[0], StudentReviewCourseColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "student_review_course_student_id",
				Columns:    []*schema.Column{StudentReviewCourseColumns[0]},
				RefColumns: []*schema.Column{StudentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "student_review_course_review_course_id",
				Columns:    []*schema.Column{StudentReviewCourseColumns[1]},
				RefColumns: []*schema.Column{ReviewCoursesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// StudentReviewTutorColumns holds the columns for the "student_review_tutor" table.
	StudentReviewTutorColumns = []*schema.Column{
		{Name: "student_id", Type: field.TypeUUID},
		{Name: "review_tutor_id", Type: field.TypeInt},
	}
	// StudentReviewTutorTable holds the schema information for the "student_review_tutor" table.
	StudentReviewTutorTable = &schema.Table{
		Name:       "student_review_tutor",
		Columns:    StudentReviewTutorColumns,
		PrimaryKey: []*schema.Column{StudentReviewTutorColumns[0], StudentReviewTutorColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "student_review_tutor_student_id",
				Columns:    []*schema.Column{StudentReviewTutorColumns[0]},
				RefColumns: []*schema.Column{StudentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "student_review_tutor_review_tutor_id",
				Columns:    []*schema.Column{StudentReviewTutorColumns[1]},
				RefColumns: []*schema.Column{ReviewTutorsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TutorReviewTutorColumns holds the columns for the "tutor_review_tutor" table.
	TutorReviewTutorColumns = []*schema.Column{
		{Name: "tutor_id", Type: field.TypeUUID},
		{Name: "review_tutor_id", Type: field.TypeInt},
	}
	// TutorReviewTutorTable holds the schema information for the "tutor_review_tutor" table.
	TutorReviewTutorTable = &schema.Table{
		Name:       "tutor_review_tutor",
		Columns:    TutorReviewTutorColumns,
		PrimaryKey: []*schema.Column{TutorReviewTutorColumns[0], TutorReviewTutorColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tutor_review_tutor_tutor_id",
				Columns:    []*schema.Column{TutorReviewTutorColumns[0]},
				RefColumns: []*schema.Column{TutorsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "tutor_review_tutor_review_tutor_id",
				Columns:    []*schema.Column{TutorReviewTutorColumns[1]},
				RefColumns: []*schema.Column{ReviewTutorsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ClassesTable,
		CoursesTable,
		IssueReportsTable,
		MatchesTable,
		PaymentsTable,
		PaymentHistoriesTable,
		ReviewCoursesTable,
		ReviewTutorsTable,
		SchedulesTable,
		StudentsTable,
		TutorsTable,
		UsersTable,
		ClassMatchTable,
		CourseReviewCourseTable,
		CourseMatchTable,
		StudentReviewCourseTable,
		StudentReviewTutorTable,
		TutorReviewTutorTable,
	}
)

func init() {
	ClassesTable.ForeignKeys[0].RefTable = PaymentHistoriesTable
	ClassesTable.ForeignKeys[1].RefTable = SchedulesTable
	CoursesTable.ForeignKeys[0].RefTable = TutorsTable
	IssueReportsTable.ForeignKeys[0].RefTable = TutorsTable
	IssueReportsTable.ForeignKeys[1].RefTable = UsersTable
	MatchesTable.ForeignKeys[0].RefTable = StudentsTable
	PaymentsTable.ForeignKeys[0].RefTable = UsersTable
	PaymentHistoriesTable.ForeignKeys[0].RefTable = PaymentsTable
	PaymentHistoriesTable.ForeignKeys[1].RefTable = UsersTable
	StudentsTable.ForeignKeys[0].RefTable = UsersTable
	TutorsTable.ForeignKeys[0].RefTable = SchedulesTable
	TutorsTable.ForeignKeys[1].RefTable = UsersTable
	ClassMatchTable.ForeignKeys[0].RefTable = ClassesTable
	ClassMatchTable.ForeignKeys[1].RefTable = MatchesTable
	CourseReviewCourseTable.ForeignKeys[0].RefTable = CoursesTable
	CourseReviewCourseTable.ForeignKeys[1].RefTable = ReviewCoursesTable
	CourseMatchTable.ForeignKeys[0].RefTable = CoursesTable
	CourseMatchTable.ForeignKeys[1].RefTable = MatchesTable
	StudentReviewCourseTable.ForeignKeys[0].RefTable = StudentsTable
	StudentReviewCourseTable.ForeignKeys[1].RefTable = ReviewCoursesTable
	StudentReviewTutorTable.ForeignKeys[0].RefTable = StudentsTable
	StudentReviewTutorTable.ForeignKeys[1].RefTable = ReviewTutorsTable
	TutorReviewTutorTable.ForeignKeys[0].RefTable = TutorsTable
	TutorReviewTutorTable.ForeignKeys[1].RefTable = ReviewTutorsTable
}
