package datas

import (
	"context"
	"fmt"
	"log"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
)

func InsertMatch(client *ent.Client, ctx context.Context, class []*ent.Class, course []*ent.Course, student []*ent.Student) []*ent.Match {
	match1 := CreateMatch(client, class[0], course[0], student[0])
	match2 := CreateMatch(client, class[0], course[0], student[1])
	match3 := CreateMatch(client, class[1], course[1], student[0])
	match4 := CreateMatch(client, class[2], course[2], student[1])

	match, err := client.Match.CreateBulk(match1, match2, match3, match4).Save(ctx)

	if err != nil {
		log.Fatalf("failed creating Match: %v", err)
	}

	for _, m := range match {
		fmt.Println("Match created: ", m.ID)
	}

	return match

}

func CreateMatch(client *ent.Client, class *ent.Class, course *ent.Course, student *ent.Student) *ent.MatchCreate {
	match := client.Match.
		Create().
		AddClasIDs(class.ID).
		AddCourseIDs(course.ID).
		SetStudentID(student.ID)

	return match
}
