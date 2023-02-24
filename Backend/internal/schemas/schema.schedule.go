package schemas

type SchemaGetSchedule struct {
	Username string `json:"username"`
}

type SchemaUpdateSchedule struct {
	Username string   `json:"username"` // TODO this must be removed if jwt is available
	Schedule Schedule `json:"schedule"`
}

type SchemaRawSchedule struct {
	Sunday    [24]bool `json:"sunday"`
	Monday    [24]bool `json:"monday"`
	Tuesday   [24]bool `json:"tuesday"`
	Wednesday [24]bool `json:"wednesday"`
	Thursday  [24]bool `json:"thursday"`
	Friday    [24]bool `json:"friday"`
	Saturday  [24]bool `json:"saturday"`
}

// --------- Internal Struct ------------
// these are not the base of api payload format
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
