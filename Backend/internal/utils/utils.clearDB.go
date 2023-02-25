package utils

import (
	"context"
	"fmt"
	"log"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	
)

func ClearDB(client *ent.Client, ctx context.Context) {
	
	// Clear all data in the database.
	
	if _, err := client.Course.Delete().Exec(ctx); err!= nil{
		log.Fatalf("failed clearing course: %v", err)
	}

	if _, err := client.Tutor.Delete().Exec(ctx); err!= nil{
		log.Fatalf("failed clearing tutor: %v", err)
	}
	
	if _, err := client.Schedule.Delete().Exec(ctx); err!= nil{
		log.Fatalf("failed clearing schedule: %v", err)
	}

	if _, err :=client.Student.Delete().Exec(ctx); err!= nil{
		log.Fatalf("failed clearing student: %v", err)
	}

	if _, err := client.User.Delete().Exec(ctx); err!= nil{
		log.Fatalf("failed clearing user: %v", err)
	}

	if _, err := client.IssueReport.Delete().Exec(ctx); err!= nil{
		log.Fatalf("failed clearing issue report: %v", err)
	}

	if _, err := client.Payment.Delete().Exec(ctx); err!= nil{
		log.Fatalf("failed clearing payment: %v", err)
	}

	if _, err := client.PaymentHistory.Delete().Exec(ctx); err!= nil{
		log.Fatalf("failed clearing payment history: %v", err)
	}

	if _, err := client.ReviewCourse.Delete().Exec(ctx); err!= nil{
		log.Fatalf("failed clearing review course: %v", err)
	}

	if _, err := client.ReviewTutor.Delete().Exec(ctx); err!= nil{
		log.Fatalf("failed clearing review tutor: %v", err)
	}


	if _, err := client.Class.Delete().Exec(ctx); err!= nil{
		log.Fatalf("failed clearing class: %v", err)
	}

	fmt.Print("\n\t::::::::: Database cleared! :::::::::\n")
}
