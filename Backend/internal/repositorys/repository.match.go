package repositorys

import (
	"context"
	"fmt"
	"time"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/course"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/schedule"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/student"

	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type RepositoryMatch interface {
	CreateMatch(sr *schema.SchemaCreateMatch) (*ent.Match, error)
}

type repositoryMatch struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepositoryMatch(c *ent.Client) *repositoryMatch {
	return &repositoryMatch{client: c, ctx: context.Background()}
}

func (r *repositoryMatch) CreateMatch(sr *schema.SchemaCreateMatch) (*ent.Match, error) {
	//Check if course and student exist
	course, err := r.client.Course.
		Query().
		Where(course.IDEQ(sr.Course_id)).
		Only(r.ctx)

	if err != nil {
		return nil, err
	}

	student, err := r.client.Student.
		Query().
		Where(student.IDEQ(sr.Student_id)).
		Only(r.ctx)

	if err != nil {
		return nil, err
	}

	// Create Schedule
	schedule, err := r.CreateSchedule(&schema.SchemaCreateSchedule{
		Schedule: sr.Schedule,
	})

	if err != nil {
		return nil, err
	}

	//Create Match
	match, err := r.client.Match.
		Create().
		SetCourse(course).
		SetStudent(student).
		SetSchedule(schedule).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	// Create Appointment
	appointments, err := r.CreateAppointment(&schema.SchemaCreateAppointment{
		Match:      match,
		Schedule:   sr.Schedule,
		Total_hour: course.EstimatedTime,
	})

	if err != nil {
		return nil, err
	}

	// updating tutor schedule and check if it is available
	_, err = r.UpdateTutorSchedule(&schema.SchemaUpdateTutorSchedule{
		Tutor_schedule_id: course.Edges.Tutor.Edges.Schedule.ID,
		Schedule:          sr.Schedule,
	})

	if err != nil {
		return nil, err
	}

	// Add Appointment to Match
	match, err = match.Update().
		AddAppointment(appointments...).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return match, nil
}

// If frontend want to create appointment by new approach, use this function
// func (r *repositoryMatch) CreateAppointment(sr *schema.SchemaCreateAppointment) ([]*ent.Appointment, error) {
// 	// Create Appointment
// 	appointments := make([]*ent.Appointment, 0)
// 	for _, time := range sr.Appointment_list {
// 		appointment, err := r.client.Appointment.
// 			Create().
// 			SetBeginAt(time.Begin_at).
// 			SetEndAt(time.End_at).
// 			SetStatus("pending").
// 			Save(r.ctx)

// 		appointments = append(appointments, appointment)

// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	return appointments, nil
// }

func (r *repositoryMatch) CreateSchedule(sr *schema.SchemaCreateSchedule) (*ent.Schedule, error) {
	ata := sr.Schedule.AvailableTimeArrays()
	schedule, err := r.client.Schedule.
		Create().
		SetDay0(ata[0]).
		SetDay1(ata[1]).
		SetDay2(ata[2]).
		SetDay3(ata[3]).
		SetDay4(ata[4]).
		SetDay5(ata[5]).
		SetDay6(ata[6]).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return schedule, nil
}

// Update tutor schedule when creating match
func (r *repositoryMatch) UpdateTutorSchedule(sr *schema.SchemaUpdateTutorSchedule) (*ent.Schedule, error) {
	tutorSchedule := r.client.Schedule.
		Query().
		Where(schedule.IDEQ(sr.Tutor_schedule_id)).
		OnlyX(r.ctx)

	aT := [7][24]bool{}
	aT[0] = tutorSchedule.Day0
	aT[1] = tutorSchedule.Day1
	aT[2] = tutorSchedule.Day2
	aT[3] = tutorSchedule.Day3
	aT[4] = tutorSchedule.Day4
	aT[5] = tutorSchedule.Day5
	aT[6] = tutorSchedule.Day6

	for _, time := range sr.Schedule {
		// Check if time is available
		if !aT[time.Day][time.Hour] {
			return nil, fmt.Errorf("time is not available")
		}

		// Set time to unavailable
		aT[time.Day][time.Hour] = false
	}

	schedule, err := tutorSchedule.Update().
		SetDay0(aT[0]).
		SetDay1(aT[1]).
		SetDay2(aT[2]).
		SetDay3(aT[3]).
		SetDay4(aT[4]).
		SetDay5(aT[5]).
		SetDay6(aT[6]).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return schedule, nil
}

func (r *repositoryMatch) CreateAppointment(sr *schema.SchemaCreateAppointment) ([]*ent.Appointment, error) {
	ata := sr.Schedule.AvailableTimeArrays()
	appointments := make([]*ent.Appointment, 0)

	dayNow := r.GetNow()
	day := (int(dayNow.Weekday()) + 1) % 7
	hoursCount := 0
	totalHour := sr.Total_hour
	dayCount := 1

	for totalHour >= 0 {
		// Find the first available time
		for !ata[day][hoursCount] {
			hoursCount++
			if hoursCount == 24 {
				day = (day + 1) % 7
				dayCount++
				hoursCount = 0
			}
		}

		//Create Appointment
		beginAtDay := dayNow.AddDate(0, 0, dayCount)
		beginAt := time.Date(beginAtDay.Year(), beginAtDay.Month(), beginAtDay.Day(), hoursCount, 0, 0, 0, beginAtDay.Location())

		endAt := time.Date(beginAtDay.Year(), beginAtDay.Month(), beginAtDay.Day(), hoursCount+1, 0, 0, 0, beginAtDay.Location())

		appointment, err := r.client.Appointment.
			Create().
			SetBeginAt(beginAt).
			SetEndAt(endAt).
			SetStatus("ongoing").
			Save(r.ctx)

		if err != nil {
			return nil, err
		}

		appointments = append(appointments, appointment)

		totalHour -= 1
		hoursCount++
		if hoursCount == 24 {
			day = (day + 1) % 7
			dayCount++
			hoursCount = 0
		}

	}
	return appointments, nil
}

func (r *repositoryMatch) GetNow() time.Time {
	// Load the Thai timezone location
	loc, _ := time.LoadLocation("Asia/Bangkok")

	// Get the current time in the Thai timezone
	now := time.Now().In(loc)

	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	return now
}
