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
	course1 := CreateCourse(client, tutor[0], "Mathematics for Everyone", "Mathematics", "Algebra", 60, "Algebra is a branch of mathematics that studies the properties of objects under the action of groups, rings, and other algebraic structures.", 100, Course.LevelGrade10, "https://se2-tuder.s3.us-west-1.amazonaws.com/course1.jpg")
	course2 := CreateCourse(client, tutor[1], "English Conversation practice", "English", "Romantic purpose", 284, "English is a language that we need to know more about it", 2400, Course.LevelGrade12, "https://se2-tuder.s3.us-west-1.amazonaws.com/course2.jpg")
	course3 := CreateCourse(client, tutor[0], "Introduction to Go", "Programming", "Golang", 120, "Learn the basics of Go programming language", 150, Course.LevelGrade11, "https://se2-tuder.s3.us-west-1.amazonaws.com/course3.jpg")
	course4 := CreateCourse(client, tutor[1], "Deeplearning Specialization", "Deeplearning", "Artificial Intelligence", 420, "All you need to know about DL", 1500, Course.LevelGrade12, "https://d3njjcbhbojbot.cloudfront.net/api/utilities/v1/imageproxy/https://d15cw65ipctsrr.cloudfront.net/a4/079d5e7c7b45ac9107f22bfcfeab91/Specialization-logo.png?auto=format%2Ccompress&dpr=1&w=330&h=330&fit=fill&q=25")
	course5 := CreateCourse(client, tutor[0], "Introduction to Python", "Programming", "Python", 320, "Learn the basics of Python programming language", 450, Course.LevelGrade9, "https://i.ytimg.com/vi/uYjRzbP5aZs/maxresdefault.jpg")
	course6 := CreateCourse(client, tutor[1], "Introduction to Java", "Programming", "Java", 20, "Learn the basics of Java programming language", 450, Course.LevelGrade10, "https://artoftesting.com/wp-content/uploads/2020/03/introduction-to-java-featured.jpg")
	course7 := CreateCourse(client, tutor[0], "Introduction to C", "Programming", "C", 30, "Learn the basics of C programming language", 550, Course.LevelGrade7, "https://i.ytimg.com/vi/gEJBFKDkqTE/maxresdefault.jpg")

	course, err := client.Course.CreateBulk(course1, course2, course3, course4, course5, course6, course7).Save(ctx)

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
