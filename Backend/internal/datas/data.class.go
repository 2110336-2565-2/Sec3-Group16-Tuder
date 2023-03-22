package datas

import (
	"context"
	"fmt"
	"log"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	class "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/class"
	"github.com/google/uuid"
)

func InsertClass(client *ent.Client, ctx context.Context, schedule []*ent.Schedule, ph []*ent.PaymentHistory) []*ent.Class {

	// Create class
	class1 := CreateClass(client, true, 30, 20, class.StatusOngoing, schedule[0], ph[0])
	class2 := CreateClass(client, true, 50, 25, class.StatusOngoing, schedule[1], ph[1])

	class, err := client.Class.CreateBulk(class1, class2).Save(ctx)

	if err != nil {
		log.Fatalf("failed creating class: %v", err)
	}

	for _, c := range class {
		fmt.Println("Class created: ", c.ID)
	}

	return class
}

func CreateClass(
	client *ent.Client,
	review_avaliable bool,
	total_hour int,
	success_hour int,
	status class.Status,
	schedule *ent.Schedule,
	ph *ent.PaymentHistory,
) *ent.ClassCreate {

	// Generate a new UUID for the course.
	classId := uuid.New()

	// Create a new course.
	classCreate := client.Class.Create().
		SetID(classId).
		SetReviewAvaliable(review_avaliable).
		SetTotalHour(total_hour).
		SetSuccessHour(success_hour).
		SetStatus(status).
		SetPaymentHistoryID(ph.ID).
		SetScheduleID(schedule.ID)

	return classCreate
}
