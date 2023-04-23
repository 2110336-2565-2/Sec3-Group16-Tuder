package datas

import (
	"context"
	"fmt"
	"log"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/match"
	"github.com/google/uuid"
)

func InsertMatch(client *ent.Client, ctx context.Context, app []*ent.Appointment, course []*ent.Course, student []*ent.Student, sc []*ent.Schedule) []*ent.Match {
	match1 := CreateMatch(client, app[0], course[0], student[0], sc[0], match.StatusCancelling)
	match2 := CreateMatch(client, app[1], course[0], student[1], sc[1], match.StatusEnrolled)
	match3 := CreateMatch(client, app[2], course[1], student[0], sc[2], match.StatusEnrolled)
	match4 := CreateMatch(client, app[3], course[2], student[1], sc[3], match.StatusEnrolled)

	match, err := client.Match.CreateBulk(match1, match2, match3, match4).Save(ctx)

	if err != nil {
		log.Fatalf("failed creating Match: %v", err)
	}

	for _, m := range match {
		fmt.Println("Match created: ", m.ID)
	}

	return match

}

func CreateMatch(client *ent.Client, app *ent.Appointment, course *ent.Course, student *ent.Student, sc *ent.Schedule, s match.Status) *ent.MatchCreate {
	mId := uuid.New()
	match := client.Match.
		Create().
		SetID(mId).
		AddAppointmentIDs(app.ID).
		SetCourseID(course.ID).
		SetStudentID(student.ID).
		SetScheduleID(sc.ID).
		SetStatus(s)

	return match
}
