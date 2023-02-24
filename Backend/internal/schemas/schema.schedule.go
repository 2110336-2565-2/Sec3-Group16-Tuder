package schemas

type TimeSlot struct {
	Day  int `json:"day,omitempty"`
	Hour int `json:"hour,omitempty"`
}

// a tuple specifying the occupied hour on a day of the week
type Schedule []TimeSlot

func (s *Schedule) AvailableTimeArrays() *[7][24]bool {
	aT := [7][24]bool{}
	for _, tS := range *s {
		aT[tS.Day][tS.Hour] = true
	}
	return &aT
}
