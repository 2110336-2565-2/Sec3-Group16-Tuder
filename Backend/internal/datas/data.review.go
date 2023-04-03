package datas

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
)

func InsertReview(client *ent.Client, ctx context.Context, student []*ent.Student, course []*ent.Course) []*ent.Review {

	// Create reviews
	review1 := CreateReview(client, student[0], course[0], "I am a review 1", 3)
	review2 := CreateReview(client, student[1], course[2], "I am a review 2", 2)
	review3 := CreateReview(client, student[0], course[0], "I am a review 3", 5)
	review4 := CreateReview(client, student[1], course[2], "I am a review 4", 3)

	review, err := client.Review.CreateBulk(review1, review2, review3, review4).Save(ctx)

	if err != nil {
		log.Fatalf("failed creating review: %v", err)
	}

	for _, t := range review {
		fmt.Println("review created: ", t.ID)
	}

	return review
}

func CreateReview(
	client *ent.Client,
	student *ent.Student,
	course *ent.Course,
	reviewMsg string,
	score int8) *ent.ReviewCreate {

	reviewCreate := client.Review.Create().
		SetReviewMsg(reviewMsg).
		SetScore(int8(score)).
		SetReviewTimeAt(time.Now()).
		AddCourse(course).
		AddStudent(student)
	return reviewCreate
}
