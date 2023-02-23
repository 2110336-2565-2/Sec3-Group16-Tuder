package datas

import (
	"context"
	"fmt"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	Course "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/course"
	util "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/utils"
)

func InsertData(client *ent.Client) {

	ctx := context.Background()

	// Delete all data in database
	util.ClearDB(client, ctx)

	// Insert users
	user := InsertUser(client)

	// Insert students
	InsertStudent(client, user[0])
	InsertStudent(client, user[1])

	// Insert tutors
	var tutor []*ent.Tutor
	for i := 2; i < len(user); i++ {
		tutor = append(tutor, InsertTutor(client, user[i], fmt.Sprint(i), "tokn_test_5jx9j8y8x9y8x9y8x9y8"+fmt.Sprint(i), "I am a tutor "+fmt.Sprint(i)))
	}

	// Insert courses
	InsertCourse(client, ctx, tutor[0], "Mathematics for Boys lovers", "Mathematics", "Algebra", 60, "Algebra is a branch of mathematics that studies the properties of objects under the action of groups, rings, and other algebraic structures.", 100, Course.LevelGrade10, "https://www.mathsisfun.com/algebra/images/algebra-1.svg")
	InsertCourse(client, ctx, tutor[1], "Boys licking practice in daily life", "Romantic", "How's Boys smell", 284, "Boys is a kind of thing that we need to know more about it", 2400, Course.LevelGrade12, "Boys")
	InsertCourse(client, ctx, tutor[0], "Introduction to Go programming", "Programming", "Golang", 120, "Learn the basics of Go programming language", 150, Course.LevelGrade11, "https://golang.org/doc/gopher/gopher.png")

	// Insert schedules
	InsertSchedule(client, ctx, tutor)

}
