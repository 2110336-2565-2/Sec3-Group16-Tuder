package datas

import (
	"context"
	"fmt"
	"log"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/google/uuid"


)

func InsertPayment(client *ent.Client, ctx context.Context, user []*ent.User) []*ent.Payment {
	// Create payment
	p1 := CreatePayment(client, "https://picsum.photos/seed/picsum/200/300", user[0])
	p2 := CreatePayment(client, "https://picsum.photos/seed/picsum/200/300", user[1])

	p, err := client.Payment.CreateBulk(p1, p2).Save(ctx)

	if err != nil {
		log.Fatalf("failed creating payment: %v", err)
	}

	for _, c := range p {
		fmt.Println("payment created: ", c.ID)
	}

	return p
}

func CreatePayment(
	client *ent.Client,
	img string,
	user *ent.User,
) *ent.PaymentCreate {

	// Generate a new UUID for the course.
	pId := uuid.New()

	// Create a new course.
	pCreate := client.Payment.Create().
		SetID(pId).
		SetQrPictureURL(img).
		SetUserID(user.ID)

	return pCreate
}