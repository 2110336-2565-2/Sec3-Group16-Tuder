package datas

import (
	"context"
	"fmt"
	"log"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/google/uuid"
)

func InsertPaymentHistory(client *ent.Client, ctx context.Context, user []*ent.User, payment []*ent.Payment) []*ent.PaymentHistory {
	// Create payment history
	ph1 := CreatePaymentHistory(client, user[0], payment[0])
	ph2 := CreatePaymentHistory(client, user[1], payment[1])
	ph3 := CreatePaymentHistory(client, user[1], payment[1])

	ph, err := client.PaymentHistory.CreateBulk(ph1, ph2, ph3).Save(ctx)

	if err != nil {
		log.Fatalf("failed creating payment_history: %v", err)
	}

	for _, c := range ph {
		fmt.Println("payment_history created: ", c.ID)
	}

	return ph
}

func CreatePaymentHistory(
	client *ent.Client,
	user *ent.User,
	payment *ent.Payment,
) *ent.PaymentHistoryCreate {

	// Generate a new UUID for the course.
	phId := uuid.New()

	// Create a new course.
	phCreate := client.PaymentHistory.Create().
		SetID(phId).
		SetUserID(user.ID).
		SetPaymentID(payment.ID)

	return phCreate
}
