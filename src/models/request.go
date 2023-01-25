package models

type Request struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Slot struct {
	SlotSize int `json:"slotsize"`
}
