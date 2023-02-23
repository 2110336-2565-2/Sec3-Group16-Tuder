package utils

import (
	"context"
	"fmt"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
)

func ClearDB(client *ent.Client, ctx context.Context) {

	// Clear all data in the database.

	client.Student.Delete().Exec(ctx)

	client.Tutor.Delete().Exec(ctx)

	client.Course.Delete().Exec(ctx)

	client.User.Delete().Exec(ctx)

	client.IssueReport.Delete().Exec(ctx)

	client.Payment.Delete().Exec(ctx)

	client.PaymentHistory.Delete().Exec(ctx)

	client.ReviewCourse.Delete().Exec(ctx)

	client.ReviewTutor.Delete().Exec(ctx)

	client.Schedule.Delete().Exec(ctx)

	client.Class.Delete().Exec(ctx)

	fmt.Print("\n\t::::::::: Database cleared! :::::::::\n")
}
