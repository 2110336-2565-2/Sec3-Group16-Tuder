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

	var schedule []*ent.Schedule

	schedule = append(schedule, CreateSchedule(client, ctx, free_day, busy_day, some_day, free_day, busy_day, some_day, busy_day))

	schedule = append(schedule, CreateSchedule(client, ctx, busy_day, some_day, free_day, busy_day, busy_day, free_day, free_day))

	return schedule

}

func CreateSchedule(
	client *ent.Client,
	ctx context.Context,
	d0 [24]bool,
	d1 [24]bool,
	d2 [24]bool,
	d3 [24]bool,
	d4 [24]bool,
	d5 [24]bool,
	d6 [24]bool,
) *ent.Schedule {

	scheduleId := uuid.New()

	schedule, err := client.Schedule.Create().
		SetID(scheduleId).
		SetDay0(d0).
		SetDay1(d1).
		SetDay2(d2).
		SetDay3(d3).
		SetDay4(d4).
		SetDay5(d5).
		SetDay6(d6).
		Save(ctx)

	if err != nil {
		log.Fatalf("failed creating schedule: %v", err)
	}

	fmt.Println("Schedule created: ", schedule.ID)

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
