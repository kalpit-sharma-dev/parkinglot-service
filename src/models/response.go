package models

type Vehicle struct {
	Slot   string `json:"slot,omitempty"`
	Number string `json:"number,omitempty"`
	Color  string `json:"color,omitempty"`
	Model  string `json:"model,omitempty"`
	//CreatedAt time.Time `json:"created_at"`
}
