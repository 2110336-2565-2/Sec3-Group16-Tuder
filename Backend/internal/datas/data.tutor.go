package datas

import (
	"context"
	"fmt"
	"log"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	// Tutor "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/tutor"
	"github.com/google/uuid"
)

func InsertTutor(client *ent.Client, ctx context.Context, user []*ent.User, schedule []*ent.Schedule) []*ent.Tutor {

	var tutor []*ent.Tutor

	tutor = append(tutor, CreateTutor(client, ctx, user[2], schedule[0], "1", "tokn_test_5jx9z8z5q2z7q7x6z3z", "I am a tutor 1"))
	tutor = append(tutor, CreateTutor(client, ctx, user[3], schedule[1], "2", "tokn_test_d9fm40fk40f45xkfe3f", "I am a tutor 2"))

	return tutor

}

func CreateTutor(client *ent.Client, ctx context.Context, user *ent.User, schedule *ent.Schedule, cid string, omiseId string, desc string) *ent.Tutor {

	tutorId := uuid.New()

	var tutor *ent.Tutor
	var err error

	tutor, err = client.Tutor.Create().
		SetID(tutorId).
		SetUserID(user.ID).
		SetScheduleID(schedule.ID).
		SetCitizenID(cid).
		SetOmiseBankToken(omiseId).
		SetDescription(desc).
		Save(ctx)

	if err != nil {
		log.Fatalf("failed creating tutor: %v", err)
	}

	fmt.Println("Tutor created: ", tutor.ID)

	return tutor
}
