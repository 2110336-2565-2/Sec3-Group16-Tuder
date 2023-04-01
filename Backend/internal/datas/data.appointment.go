package datas

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	appointment "github.com/2110336-2565-2/Sec3-Group16-Tuder/ent/appointment"
	"github.com/google/uuid"
)

func InsertAppointment(client *ent.Client, ctx context.Context) []*ent.Appointment {

	// Create appointment
	app1 := CreateApp(client, appointment.StatusOngoing, time.Now(), time.Now().Add(1*time.Hour))
	app2 := CreateApp(client, appointment.StatusOngoing, time.Now(), time.Now().Add(1*time.Hour))
	app3 := CreateApp(client, appointment.StatusOngoing, time.Now(), time.Now().Add(1*time.Hour))
	app4 := CreateApp(client, appointment.StatusOngoing, time.Now(), time.Now().Add(1*time.Hour))

	app, err := client.Appointment.CreateBulk(app1, app2, app3, app4).Save(ctx)

	if err != nil {
		log.Fatalf("failed creating appointment: %v", err)
	}

	for _, c := range app {
		fmt.Println("Appointment created: ", c.ID)
	}

	return app
}

func CreateApp(
	client *ent.Client,
	status appointment.Status,
	bg_at time.Time,
	ed_at time.Time,
) *ent.AppointmentCreate {

	// Generate a new UUID for the course.
	appId := uuid.New()

	// Create a new course.
	appCreate := client.Appointment.Create().
		SetID(appId).
		SetStatus(status).
		SetBeginAt(bg_at).
		SetEndAt(ed_at)

	return appCreate
}
