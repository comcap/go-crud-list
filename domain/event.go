package domain

type Event struct {
	ID          string      `bson:"id"`
	CalendarId  string      `bson:"calendarId"`
	HTMLLink    string      `bson:"htmlLink"`
	Summary     string      `bson:"summary"`
	Description string      `bson:"description"`
	Start       DateTime    `bson:"start"`
	End         DateTime    `bson:"end"`
	Rooms       string      `bson:"rooms"`
	Reminders   Reminders   `bson:"reminders"`
	Creator     Creator     `bson:"creator"`
	Organizer   Creator     `bson:"organizer"`
	Attendees   []*Attendee `bson:"attendees"`
	Created     int64       `bson:"created"`
	Updated     int64       `bson:"updated"`
}

type Attendee struct {
	Email string `bson:"email"`
}

type DateTime struct {
	DateTime int64 `bson:"dateTime"`
}

type Reminders struct {
	Override []*Override `bson:"override"`
}

type Override struct {
	Method  string `bson:"method"`
	Minutes string `bson:"minutes"`
}

type Creator struct {
	Email string `bson:"email"`
	Self  bool   `bson:"self"`
}
