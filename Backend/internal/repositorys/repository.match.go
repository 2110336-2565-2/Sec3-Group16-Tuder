package repositorys

import (
	"context"
	"fmt"
	"time"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/course"
	entMatch "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/match"
	entPayment "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/payment"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/schedule"
	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/student"
	"github.com/google/uuid"

	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

type RepositoryMatch interface {
	CreateMatch(sr *schema.SchemaCreateMatch) (*ent.Match, error)
	GetMatchByID(id uuid.UUID) (*ent.Match, error)
}

type repositoryMatch struct {
	client *ent.Client
	ctx    context.Context
}

func NewRepositoryMatch(c *ent.Client) *repositoryMatch {
	return &repositoryMatch{client: c, ctx: context.Background()}
}

func (r *repositoryMatch) CreateMatch(sr *schema.SchemaCreateMatch) (*ent.Match, error) {
	// create a transaction
	tx, err := r.client.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("starting a transaction: %w", err)
	}

	// wrap the client with the transaction
	txc := tx.Client()

	//Check if course and student exist
	course, err := txc.Course.
		Query().
		Where(course.IDEQ(sr.Course_id)).
		WithTutor(func(tq *ent.TutorQuery) {
			tq.WithSchedule()
		}).
		Only(r.ctx)

	if err != nil {
		return nil, err
	}

	student, err := txc.Student.
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

	// Create Appointment
	appointments, err := r.CreateAppointment(&schema.SchemaCreateAppointment{
		Schedule:   sr.Schedule,
		Total_hour: sr.Total_hour,
	})

	if err != nil {
		fmt.Println("Create Appointment: %w", err)
		return nil, err
	}

	// updating tutor schedule and check if it is available
	_, err = r.UpdateTutorSchedule(&schema.SchemaUpdateTutorSchedule{
		Tutor_schedule_id: course.Edges.Tutor.Edges.Schedule.ID,
		Schedule:          sr.Schedule,
	})

	if err != nil {
		fmt.Println("Update Tutor Schedule: %w", err)
		return nil, err
	}

	//Check if payment exist
	payment, err := txc.Payment.
		Query().
		Where(entPayment.IDEQ(sr.Payment_id)).
		Only(r.ctx)

	if err != nil {
		return nil, err
	}

	//Create Match
	match, err := txc.Match.
		Create().
		SetCourse(course).
		AddAppointment(appointments...).
		SetStudent(student).
		SetScheduleID(schedule.ID).
		SetPaymentID(payment.ID).
		Save(r.ctx)

	if err != nil {
		fmt.Println("Create Match: %w", err)
		return nil, err
	}

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, fmt.Errorf("creating match: %w", err)
	}

	return match, tx.Commit()
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
	// create a transaction
	tx, err := r.client.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("starting a transaction: %w", err)
	}
	// wrap the client with the transaction
	txc := tx.Client()

	ata := sr.Schedule.AvailableTimeArrays()
	schedule, err := txc.Schedule.
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

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, fmt.Errorf("creating schedule: %w", err)
	}

	return schedule, tx.Commit()
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
	// create a transaction
	tx, err := r.client.Tx(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("starting a transaction: %w", err)
	}
	// wrap the client with the transaction
	txc := tx.Client()

	ata := sr.Schedule.AvailableTimeArrays()
	appointments := make([]*ent.Appointment, 0)

	dayNow := r.GetNow()
	day := (int(dayNow.Weekday()) + 1) % 7
	hoursCount := 0
	totalHour := sr.Total_hour
	dayCount := 1

	for totalHour > 0 {
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

		appointment, err := txc.Appointment.
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

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, fmt.Errorf("creating course: %w", err)
	}
	return appointments, tx.Commit()
}

func (r *repositoryMatch) GetNow() time.Time {
	// Load the Thai timezone location
	loc, _ := time.LoadLocation("Asia/Bangkok")

	// Get the current time in the Thai timezone
	now := time.Now().In(loc)

	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	return now
}

func (r *repositoryMatch) GetMatchByID(id uuid.UUID) (*ent.Match, error) {
	match, err := r.client.Match.
		Query().
		Where(entMatch.IDEQ(id)).
		WithStudent().
		WithCourse().
		WithSchedule().
		WithAppointment().
		Only(r.ctx)

	if err != nil {
		return nil, err
	}

	return match, nil
}
