package notification

type Sick struct {
	Reason      string
	StartDate   string
	EndDate     string
	Alternate   string
	AlternateRm string
}

func (s *Sick) GetMessageText() string {
	return "and easy to do anywhere, even with Go"
}
