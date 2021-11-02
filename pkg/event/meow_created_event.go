package event

import "time"

const MeowCreatedEventSubject = "meow_created_subject"

type MeowCreatedEvent struct {
	ID        int
	Body      string
	CreatedOn time.Time
}
