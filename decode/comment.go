package decode

type Comment struct {
	CreatedTime  time.Time     `json: created_time`
	CommentID    CommentID     `json: comment_id`
	Community    Community     `json: community`
	DeletedTime  null.NullTime `json: deleted_time`
	From         UserID        `json: from`
	Message      string        `json: message`
	MessageTags  []MessageTag  `json: message_tags`
	Parent       Comment       `json: parent`
	PostID       PostID        `json: post_id`
	ParmalinkUrl string        `json: parmalink_url`
	Verb         Verb          `json: verb`
}
