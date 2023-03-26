package datas

import (
	"context"
	"fmt"
	"log"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
)

func InsertStudent(client *ent.Client, ctx context.Context, user []*ent.User) []*ent.Student {

	// Create students
	student1 := CreateStudent(client, user[0])
	student2 := CreateStudent(client, user[1])

	student, err := client.Student.CreateBulk(student1, student2).Save(ctx)

	if err != nil {
		log.Fatalf("failed creating student: %v", err)
	}

	for _, s := range student {
		fmt.Println("Student created: ", s.ID)
	}

	return student
}

func CreateStudent(client *ent.Client, user *ent.User) *ent.StudentCreate {

	studentCreate := client.Student.Create().
		SetUserID(user.ID)

	return studentCreate
}
