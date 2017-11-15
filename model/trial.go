package model

// Trial model
type Trial struct {
	ID          string `json:"id,omitempty"`
	OwnerUD     string `json:"ownerId,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	StartDate   int    `json:"startDate,omitempty"`
	EndDate     int    `json:"endDate,omitempty"`
}
