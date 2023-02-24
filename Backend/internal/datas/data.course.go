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

	var course []*ent.Course

	course = append(course, CreateCourse(client, ctx, tutor[0], "Mathematics for Boys lovers", "Mathematics", "Algebra", 60, "Algebra is a branch of mathematics that studies the properties of objects under the action of groups, rings, and other algebraic structures.", 100, Course.LevelGrade10, "https://picsum.photos/seed/picsum/200/100"))
	course = append(course, CreateCourse(client, ctx, tutor[1], "Boys licking practice in daily life", "Romantic", "How's Boys smell", 284, "Boys is a kind of thing that we need to know more about it", 2400, Course.LevelGrade12, "https://picsum.photos/seed/picsum/200/100"))
	course = append(course, CreateCourse(client, ctx, tutor[0], "Introduction to Go programming", "Programming", "Golang", 120, "Learn the basics of Go programming language", 150, Course.LevelGrade11, "https://picsum.photos/seed/picsum/200/100"))

	return course
}

func CreateCourse(
	client *ent.Client,
	ctx context.Context,
	tutor *ent.Tutor,
	title string,
	subject string,
	topic string,
	estimatedTime int,
	description string,
	pricePerHour int,
	level Course.Level,
	coursePictureURL string,
) *ent.Course {

	// Generate a new UUID for the course.
	courseId := uuid.New()

	// Create a new course.
	course, err := client.Course.Create().
		SetID(courseId).
		SetTutorID(tutor.ID).
		SetTitle(title).
		SetSubject(subject).
		SetTopic(topic).
		SetEstimatedTime(estimatedTime).
		SetDescription(description).
		SetPricePerHour(pricePerHour).
		SetLevel(level).
		SetCoursePictureURL(coursePictureURL).
		Save(ctx)

	if err != nil {
		log.Fatalf("failed creating course: %v", err)
	}

	fmt.Println("Course created: ", course.ID)

	return course
}
