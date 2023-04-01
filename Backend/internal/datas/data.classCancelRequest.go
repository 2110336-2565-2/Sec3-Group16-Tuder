package datas

import (
	"context"
	"fmt"
	"log"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	cancelrequest "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/cancelrequest"
	"github.com/google/uuid"
)

func InsertCancelRequest(client *ent.Client, ctx context.Context, user []*ent.User, match []*ent.Match) []*ent.CancelRequest {

	r1 := CreateCancelRequest(client, "Class Cancel Request 1", user[0], match[0], "Description 1", "https://www.google.com", cancelrequest.StatusPending)

	ccr, err := client.CancelRequest.CreateBulk(r1).Save(ctx)

	if err != nil {
		log.Fatalf("failed creating cancel request: %v", err)
	}

	for _, c := range ccr {
		fmt.Println("Cancel Request created: ", c.ID)
	}

	return ccr
}

func CreateCancelRequest(
	client *ent.Client,
	title string,
	user *ent.User,
	match *ent.Match,
	description string,
	imgURL string,
	status cancelrequest.Status,

) *ent.CancelRequestCreate {

	// Generate a new UUID for the course.
	requestId := uuid.New()

	// Create a new course.
	requestCreate := client.CancelRequest.Create().
		SetID(requestId).
		SetDescription(description).
		SetMatchID(match.ID).
		SetStatus(status).
		SetImgURL(imgURL).
		SetTitle(title).
		SetUserID(user.ID)

	return requestCreate
}
