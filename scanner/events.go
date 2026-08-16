package scanner

import "time"

type Event struct {
	Deskripshion string
	UserInpur    string
	DateAdd      time.Time
}

func NewEvent(Deskripshion string, UserInpur string) Event {
	return Event{
		DateAdd:      time.Now(),
		UserInpur:    UserInpur,
		Deskripshion: Deskripshion,
	}
}
