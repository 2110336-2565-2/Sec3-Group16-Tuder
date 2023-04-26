package datas

import (
	"context"
	"fmt"
	"log"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	Course "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/course"
	"github.com/google/uuid"
)

func InsertCourse(client *ent.Client, ctx context.Context, tutor []*ent.Tutor) []*ent.Course {

	// Create courses
	course1 := CreateCourse(client, tutor[0], "Mathematics for Boys lovers", "Mathematics", "Algebra", 60, "Algebra is a branch of mathematics that studies the properties of objects under the action of groups, rings, and other algebraic structures.", 100, Course.LevelGrade10, "https://se2-tuder.s3.us-west-1.amazonaws.com/course1.jpg")
	course2 := CreateCourse(client, tutor[1], "English Conversation practice", "Romantic", "English", 284, "English is a language that we need to know more about it", 2400, Course.LevelGrade12, "https://se2-tuder.s3.us-west-1.amazonaws.com/course2.jpg")
	course3 := CreateCourse(client, tutor[0], "Introduction to Go programming", "Programming", "Golang", 120, "Learn the basics of Go programming language", 150, Course.LevelGrade11, "https://se2-tuder.s3.us-west-1.amazonaws.com/course3.jpg")

	course, err := client.Course.CreateBulk(course1, course2, course3).Save(ctx)

	if err != nil {
		log.Fatalf("failed creating course: %v", err)
	}

	for _, c := range course {
		fmt.Println("Course created: ", c.ID)
	}

	return course
}

func CreateCourse(
	client *ent.Client,
	tutor *ent.Tutor,
	title string,
	subject string,
	topic string,
	estimatedTime int,
	description string,
	pricePerHour int,
	level Course.Level,
	coursePictureURL string,
) *ent.CourseCreate {

	// Generate a new UUID for the course.
	courseId := uuid.New()

	// Create a new course.
	courseCreate := client.Course.Create().
		SetID(courseId).
		SetTutorID(tutor.ID).
		SetTitle(title).
		SetSubject(subject).
		SetTopic(topic).
		SetEstimatedTime(estimatedTime).
		SetDescription(description).
		SetPricePerHour(pricePerHour).
		SetLevel(level).
		SetCoursePictureURL(coursePictureURL)

	return courseCreate
}
