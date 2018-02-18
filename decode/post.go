package decode

import (
	"time"

	"github.com/guregu/null"
)

type Post struct {
	CreatedTime  time.Time     `json: created_time`
	Community    Community     `json: commnity`
	DeletedTime  null.NullTime `json: deleted_time`
	From         UserID        `json: from`
	Message      string        `json: message`
	MessageTags  []MessageTag  `json: message_tags`
	ParmalinkUrl string        `json: parmalink_url`
	PostID       PostID        `json: post_id`
	// TargetType
	Type string `json: type`
	Verb Verb   `json: verb`
}
