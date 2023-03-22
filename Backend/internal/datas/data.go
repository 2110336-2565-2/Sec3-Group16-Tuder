package datas

import (
	"context"
	"fmt"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	util "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/utils"
)

func InsertData(client *ent.Client) {

	ctx := context.Background()

	// Delete all data in database
	util.ClearDB(client, ctx)

	// Insert users
	user := InsertUser(client, ctx)

	// Insert students
	InsertStudent(client, ctx, user)

	// Insert schedules
	schedule := InsertSchedule(client, ctx)

	// Insert tutors
	tutor := InsertTutor(client, ctx, user, schedule)

	// Insert courses
	InsertCourse(client, ctx, tutor)

	// Insert payment
	payment := InsertPayment(client, ctx, user)

	// Insert payment history
	ph := InsertPaymentHistory(client, ctx, user, payment)

	// Insert classes
	InsertClass(client, ctx, schedule, ph)

	fmt.Print("\n\t::::::::: Data inserted! :::::::::\n")

}
