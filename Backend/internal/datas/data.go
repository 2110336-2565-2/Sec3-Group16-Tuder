package datas

import (
	"context"
	"fmt"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
)

func InsertData(client *ent.Client) {

	ctx := context.Background()

	// Insert users
	user := InsertUser(client, ctx)

	// Insert students
	student := InsertStudent(client, ctx, user)

	// Insert schedules
	schedule := InsertSchedule(client, ctx)

	// Insert tutors
	tutor := InsertTutor(client, ctx, user, schedule)

	// Insert courses
	course := InsertCourse(client, ctx, tutor)

	// Insert payment
	payment := InsertPayment(client, ctx, user)

	// Insert payment history
	ph := InsertPaymentHistory(client, ctx, user, payment)

	// Insert classes
	class := InsertClass(client, ctx, schedule, ph)

	// Insert match
	InsertMatch(client, ctx, class, course, student)

	fmt.Print("\n\t::::::::: Data inserted! :::::::::\n")

}
