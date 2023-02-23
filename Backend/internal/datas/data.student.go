package datas

import (
	"context"
	"fmt"
	"log"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
)

func InsertStudent(client *ent.Client, ctx context.Context, user []*ent.User) ([]*ent.Student) {

	var student []*ent.Student

	student = append(student, CreateStudent(client, user[0]))
	student = append(student, CreateStudent(client, user[1]))

	return student
}

func CreateStudent(client *ent.Client, user *ent.User) (*ent.Student) {

	ctx := context.Background()

	student, err := client.Student.Create().
		SetUserID(user.ID).
		Save(ctx)

	if err != nil {
		log.Fatalf("failed creating student: %v", err)
	}

	fmt.Println("Student created: ", student.ID)

	return student
}