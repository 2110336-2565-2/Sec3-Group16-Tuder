package datas

import (
	"context"
	"fmt"
	"log"

	"github.com/2110336-2565-2/Sec3-Group16-Tuder/ent"
	"github.com/google/uuid"
)

func InsertSchedule(client *ent.Client, ctx context.Context) []*ent.Schedule {

	free_day, busy_day, some_day := GenerateDay()

	// Create schedules
	schedule1 := CreateSchedule(client, free_day, busy_day, some_day, free_day, some_day, some_day, busy_day)
	schedule2 := CreateSchedule(client, busy_day, some_day, free_day, busy_day, free_day, free_day, free_day)
	schedule3 := CreateSchedule(client, free_day, busy_day, free_day, busy_day, free_day, busy_day, free_day)
	schedule4 := CreateSchedule(client, some_day, some_day, busy_day, free_day, busy_day, free_day, free_day)

	schedule, err := client.Schedule.CreateBulk(schedule1, schedule2, schedule3, schedule4).Save(ctx)

	if err != nil {
		log.Fatalf("failed creating schedule: %v", err)
	}

	for _, s := range schedule {
		fmt.Println("Schedule created: ", s.ID)
	}

	return schedule

}

func CreateSchedule(
	client *ent.Client,
	d0 [24]bool,
	d1 [24]bool,
	d2 [24]bool,
	d3 [24]bool,
	d4 [24]bool,
	d5 [24]bool,
	d6 [24]bool,
) *ent.ScheduleCreate {

	sId := uuid.New()

	schedule := client.Schedule.Create().
		SetID(sId).
		SetDay0(d0).
		SetDay1(d1).
		SetDay2(d2).
		SetDay3(d3).
		SetDay4(d4).
		SetDay5(d5).
		SetDay6(d6)

	return schedule
}

func GenerateDay() ([24]bool, [24]bool, [24]bool) {
	free_day := [24]bool{}
	busy_day := [24]bool{}
	some_day := [24]bool{}
	for idx := range free_day {
		free_day[idx] = true
		busy_day[idx] = false
		if idx%3 == 0 {
			some_day[idx] = true
		} else {
			some_day[idx] = false
		}
	}

	return free_day, busy_day, some_day
}
