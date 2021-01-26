package dto

type Sermon struct {
	SermonID      string
	SermonName    string
	SermonDesc   string
	CreatedOn         string
	CreatedBy   string
	AppID  string
	Status string
}

type ListSermon struct {
	Sermon []Sermon
}