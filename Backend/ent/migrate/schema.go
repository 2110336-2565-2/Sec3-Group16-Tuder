// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AppointmentsColumns holds the columns for the "appointments" table.
	AppointmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "begin_at", Type: field.TypeTime},
		{Name: "end_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"comingsoon", "ongoing", "verifying", "pending", "completed", "postponed", "considering", "canceled"}, Default: "comingsoon"},
		{Name: "appointment_match", Type: field.TypeUUID, Nullable: true},
	}
	// AppointmentsTable holds the schema information for the "appointments" table.
	AppointmentsTable = &schema.Table{
		Name:       "appointments",
		Columns:    AppointmentsColumns,
		PrimaryKey: []*schema.Column{AppointmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "appointments_matches_match",
				Columns:    []*schema.Column{AppointmentsColumns[4]},
				RefColumns: []*schema.Column{MatchesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CancelRequestsColumns holds the columns for the "cancel_requests" table.
	CancelRequestsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "title", Type: field.TypeString},
		{Name: "report_date", Type: field.TypeTime},
		{Name: "img_url", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"pending", "approved", "rejected"}},
		{Name: "cancel_request_match", Type: field.TypeUUID},
		{Name: "user_cancel_request", Type: field.TypeUUID},
	}
	// CancelRequestsTable holds the schema information for the "cancel_requests" table.
	CancelRequestsTable = &schema.Table{
		Name:       "cancel_requests",
		Columns:    CancelRequestsColumns,
		PrimaryKey: []*schema.Column{CancelRequestsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cancel_requests_matches_match",
				Columns:    []*schema.Column{CancelRequestsColumns[6]},
				RefColumns: []*schema.Column{MatchesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "cancel_requests_users_cancel_request",
				Columns:    []*schema.Column{CancelRequestsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
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
		{Name: "contact", Type: field.TypeString, Default: "No contact"},
		{Name: "report_date", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"ongoing", "completed", "rejected"}, Default: "ongoing"},
		{Name: "tutor_issue_report", Type: field.TypeUUID, Nullable: true},
	}
	// IssueReportsTable holds the schema information for the "issue_reports" table.
	IssueReportsTable = &schema.Table{
		Name:       "issue_reports",
		Columns:    IssueReportsColumns,
		PrimaryKey: []*schema.Column{IssueReportsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "issue_reports_tutors_issue_report",
				Columns:    []*schema.Column{IssueReportsColumns[6]},
				RefColumns: []*schema.Column{TutorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MatchesColumns holds the columns for the "matches" table.
	MatchesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "match_created_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"enrolled", "completed", "cancelling", "canceled"}, Default: "enrolled"},
		{Name: "course_match", Type: field.TypeUUID},
		{Name: "schedule_match", Type: field.TypeUUID, Unique: true},
		{Name: "student_match", Type: field.TypeUUID},
	}
	// MatchesTable holds the schema information for the "matches" table.
	MatchesTable = &schema.Table{
		Name:       "matches",
		Columns:    MatchesColumns,
		PrimaryKey: []*schema.Column{MatchesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "matches_courses_match",
				Columns:    []*schema.Column{MatchesColumns[3]},
				RefColumns: []*schema.Column{CoursesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "matches_schedules_match",
				Columns:    []*schema.Column{MatchesColumns[4]},
				RefColumns: []*schema.Column{SchedulesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "matches_students_match",
				Columns:    []*schema.Column{MatchesColumns[5]},
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
	// ReviewsColumns holds the columns for the "reviews" table.
	ReviewsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "score", Type: field.TypeInt8, Nullable: true},
		{Name: "review_msg", Type: field.TypeString, Nullable: true},
		{Name: "review_time_at", Type: field.TypeTime},
	}
	// ReviewsTable holds the schema information for the "reviews" table.
	ReviewsTable = &schema.Table{
		Name:       "reviews",
		Columns:    ReviewsColumns,
		PrimaryKey: []*schema.Column{ReviewsColumns[0]},
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
		{Name: "schedule_tutor", Type: field.TypeUUID, Unique: true},
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
	// CourseReviewColumns holds the columns for the "course_review" table.
	CourseReviewColumns = []*schema.Column{
		{Name: "course_id", Type: field.TypeUUID},
		{Name: "review_id", Type: field.TypeUUID},
	}
	// CourseReviewTable holds the schema information for the "course_review" table.
	CourseReviewTable = &schema.Table{
		Name:       "course_review",
		Columns:    CourseReviewColumns,
		PrimaryKey: []*schema.Column{CourseReviewColumns[0], CourseReviewColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "course_review_course_id",
				Columns:    []*schema.Column{CourseReviewColumns[0]},
				RefColumns: []*schema.Column{CoursesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "course_review_review_id",
				Columns:    []*schema.Column{CourseReviewColumns[1]},
				RefColumns: []*schema.Column{ReviewsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// StudentReviewColumns holds the columns for the "student_review" table.
	StudentReviewColumns = []*schema.Column{
		{Name: "student_id", Type: field.TypeUUID},
		{Name: "review_id", Type: field.TypeUUID},
	}
	// StudentReviewTable holds the schema information for the "student_review" table.
	StudentReviewTable = &schema.Table{
		Name:       "student_review",
		Columns:    StudentReviewColumns,
		PrimaryKey: []*schema.Column{StudentReviewColumns[0], StudentReviewColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "student_review_student_id",
				Columns:    []*schema.Column{StudentReviewColumns[0]},
				RefColumns: []*schema.Column{StudentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "student_review_review_id",
				Columns:    []*schema.Column{StudentReviewColumns[1]},
				RefColumns: []*schema.Column{ReviewsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AppointmentsTable,
		CancelRequestsTable,
		CoursesTable,
		IssueReportsTable,
		MatchesTable,
		PaymentsTable,
		PaymentHistoriesTable,
		ReviewsTable,
		SchedulesTable,
		StudentsTable,
		TutorsTable,
		UsersTable,
		CourseReviewTable,
		StudentReviewTable,
	}
)

func init() {
	AppointmentsTable.ForeignKeys[0].RefTable = MatchesTable
	CancelRequestsTable.ForeignKeys[0].RefTable = MatchesTable
	CancelRequestsTable.ForeignKeys[1].RefTable = UsersTable
	CoursesTable.ForeignKeys[0].RefTable = TutorsTable
	IssueReportsTable.ForeignKeys[0].RefTable = TutorsTable
	MatchesTable.ForeignKeys[0].RefTable = CoursesTable
	MatchesTable.ForeignKeys[1].RefTable = SchedulesTable
	MatchesTable.ForeignKeys[2].RefTable = StudentsTable
	PaymentsTable.ForeignKeys[0].RefTable = UsersTable
	PaymentHistoriesTable.ForeignKeys[0].RefTable = PaymentsTable
	PaymentHistoriesTable.ForeignKeys[1].RefTable = UsersTable
	StudentsTable.ForeignKeys[0].RefTable = UsersTable
	TutorsTable.ForeignKeys[0].RefTable = SchedulesTable
	TutorsTable.ForeignKeys[1].RefTable = UsersTable
	CourseReviewTable.ForeignKeys[0].RefTable = CoursesTable
	CourseReviewTable.ForeignKeys[1].RefTable = ReviewsTable
	StudentReviewTable.ForeignKeys[0].RefTable = StudentsTable
	StudentReviewTable.ForeignKeys[1].RefTable = ReviewsTable
}
