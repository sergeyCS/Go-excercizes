package notification

type SickJson struct {
	Employee    string `json:"employee"`
	Reason      string `json:"reason,omitempty"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate,omitempty"`
	Alternate   string `json:"alternate"`
	AlternateRm string `json:"alternateRm,omitempty"`
}
