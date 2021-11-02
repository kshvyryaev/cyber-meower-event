package event

import "time"

type MeowCreatedEvent struct {
	ID        int
	Body      string
	CreatedOn time.Time
}
