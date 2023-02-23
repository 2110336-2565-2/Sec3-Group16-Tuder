package datas

import (
	"context"
	"fmt"
	"log"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	Tutor "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/tutor"
	"github.com/google/uuid"
)

func InsertTutor(client *ent.Client, user *ent.User, cid string, omiseId string, desc string) *ent.Tutor {

	ctx := context.Background()

	tutorId := uuid.New()

	var tutor *ent.Tutor
	var err error
	tutor, err = client.Tutor.Query().Where(Tutor.CitizenIDEQ(cid)).Only(ctx)

	if err != nil {

		tutor, err = client.Tutor.Create().
			SetID(tutorId).
			SetUserID(user.ID).
			SetCitizenID(cid).
			SetOmiseBankToken(omiseId).
			SetDescription(desc).
			Save(ctx)

		if err != nil {
			log.Fatalf("failed creating tutor: %v", err)
		}
	}

	// update tutor user id
	client.Tutor.UpdateOne(tutor).SetUserID(user.ID).Exec(ctx)

	fmt.Println("Tutor created: ", tutor.ID)

	return tutor
}
