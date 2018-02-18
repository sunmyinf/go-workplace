package decode

import "time"

type Membership struct {
	Actor      UserID    `json: actor`
	Community  Community `json: community`
	Member     UserID    `json: member`
	UpdateTime time.Time `json: update_time`
	Verb       Verb      `json: verb`
}
