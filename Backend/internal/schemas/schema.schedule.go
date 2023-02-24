package schemas

//type TimeSlot struct {
//	Day  int `json:"day,omitempty"`
//	Hour int `json:"hour,omitempty"`
//}

// a tuple specifying the occupied hour on a day of the week
type Schedule struct {
	Day  int `json:"day,omitempty"`
	Hour int `json:"hour,omitempty"`
}
