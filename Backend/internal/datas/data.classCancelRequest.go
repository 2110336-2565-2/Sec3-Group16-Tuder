package datas

import (
	"context"
	"fmt"
	"log"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	classcancelrequest "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/classcancelrequest"
	"github.com/google/uuid"
)

func InsertClassCancelRequest(client *ent.Client, ctx context.Context, user []*ent.User, match []*ent.Match) []*ent.ClassCancelRequest {

	// Create class
	r1 := CreateClassCancelRequest(client, "Class Cancel Request 1", user[0], match[0], "Description 1", "https://www.google.com", classcancelrequest.StatusPending)

	ccr, err := client.ClassCancelRequest.CreateBulk(r1).Save(ctx)

	if err != nil {
		log.Fatalf("failed creating class: %v", err)
	}

	for _, c := range ccr {
		fmt.Println("Class created: ", c.ID)
	}

	return ccr
}

func CreateClassCancelRequest(
	client *ent.Client,
	title string,
	user *ent.User,
	match *ent.Match,
	description string,
	imgURL string,
	status classcancelrequest.Status,

) *ent.ClassCancelRequestCreate {

	// Generate a new UUID for the course.
	requestId := uuid.New()

	// Create a new course.
	requestCreate := client.ClassCancelRequest.Create().
		SetID(requestId).
		SetDescription(description).
		SetMatchID(match.ID).
		SetStatus(status).
		SetImgURL(imgURL).
		SetTitle(title).
		SetUserID(user.ID)

	return requestCreate
}
