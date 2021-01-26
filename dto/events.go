package dto

type Events struct {
	EventID      string
	EventName    string
	EventStart   string
	EventDetail         string
	EventLocation   string
	CreatedDate      string
	CreatedBy   string
	EventDate string
	EventEnd  string
	EventImg  string
	AppID  string
	Status string
}

type ListEvents struct {
	Events []Events
}