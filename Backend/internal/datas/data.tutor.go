package datas

import (
	"context"
	"fmt"
	"log"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/google/uuid"
)

func InsertTutor(client *ent.Client, ctx context.Context, user []*ent.User, schedule []*ent.Schedule) []*ent.Tutor {

	// Create tutors
	tutor1 := CreateTutor(client, user[2], schedule[0], "6048764759382", "tokn_test_5jx9z8z5q2z7q7x6z3z", "I am a tutor 1", "cust_test_5vfpp7eghorusdbkggy")
	tutor2 := CreateTutor(client, user[3], schedule[1], "2194892411578", "tokn_test_d9fm40fk40f45xkfe3f", "I am a tutor 2", "cust_test_5vg2emiuoi3o7h43q18")

	tutor, err := client.Tutor.CreateBulk(tutor1, tutor2).Save(ctx)

	if err != nil {
		log.Fatalf("failed creating tutor: %v", err)
	}

	for _, t := range tutor {
		fmt.Println("Tutor created: ", t.ID)
	}

	return tutor
}

func CreateTutor(
	client *ent.Client,
	user *ent.User,
	schedule *ent.Schedule,
	cid string,
	omiseId string,
	desc string,
	custId string) *ent.TutorCreate {

	tutorId := uuid.New()

	tutorCreate := client.Tutor.Create().
		SetID(tutorId).
		SetUserID(user.ID).
		SetScheduleID(schedule.ID).
		SetCitizenID(cid).
		SetOmiseBankToken(omiseId).
		SetOmiseCustomerID(custId).
		SetDescription(desc)

	return tutorCreate
}
