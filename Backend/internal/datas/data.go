package datas

import (
	"context"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	Course "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/course"
)

func InsertData(client *ent.Client) {

	ctx := context.Background()

	// Delete all data in database
	client.Tutor.Delete().Exec(ctx)
	client.Student.Delete().Exec(ctx)
	client.User.Delete().Exec(ctx)
	client.Course.Delete().Exec(ctx)

	// Insert users
	user := InsertUser(client)

	// Insert students
	InsertStudent(client, user[0])
	InsertStudent(client, user[1])

	// Insert tutors
	tutor1 := InsertTutor(client, user[2], "1", "tokn_test_5jx9j8y8x9y8x9y8x9y8", "I am a tutor1")
	tutor2 := InsertTutor(client, user[3], "2", "tokn_test_5jx9j8y8x9y8x9y8x9y8", "I am a tutor2")

	// Insert courses
	InsertCourse(client, ctx, tutor1, "Mathematics for Boys lovers", "Mathematics", "Algebra", 60, "Algebra is a branch of mathematics that studies the properties of objects under the action of groups, rings, and other algebraic structures.", 100, Course.LevelGrade10, "https://www.mathsisfun.com/algebra/images/algebra-1.svg")
	InsertCourse(client, ctx, tutor2, "Boys licking in daily life", "Romantic", "How's Boys smell", 284, "Boys is a kind of thing that we need to know more about it", 2400, Course.LevelGrade12, "Boys")
	InsertCourse(client, ctx, tutor1, "Introduction to Go programming", "Programming", "Golang", 120, "Learn the basics of Go programming language", 150, Course.LevelGrade11, "https://golang.org/doc/gopher/gopher.png")

}
