package decode

import (
	"time"

	"github.com/guregu/null"
)

type Comment struct {
	CreatedTime  time.Time    `json: created_time`
	CommentID    string       `json: comment_id`
	Community    Community    `json: community`
	DeletedTime  null.Time    `json: deleted_time`
	From         string       `json: from`
	Message      string       `json: message`
	MessageTags  []MessageTag `json: message_tags`
	Parent       *Comment     `json: parent`
	PostID       string       `json: post_id`
	ParmalinkUrl string       `json: parmalink_url`
	Verb         Verb         `json: verb`
}

type Community struct {
	ID string `json: id`
}

type MessageTagType string

const (
	MessageTagTypeUser        MessageTagType = "user"
	MessageTagTypePage        MessageTagType = "page"
	MessageTagTypeEvent       MessageTagType = "event"
	MessageTagTypeGroup       MessageTagType = "group"
	MessageTagTypeApplication MessageTagType = "application"
)

type MessageTag struct {
	ID     string         `json: id`
	Length uint32         `json: length`
	Name   string         `json: name`
	Offset uint32         `json: offset`
	Type   MessageTagType `json: type`
}

type Membership struct {
	Actor      string    `json: actor`
	Community  Community `json: community`
	Member     string    `json: member`
	UpdateTime time.Time `json: update_time`
	Verb       Verb      `json: verb`
}

type Post struct {
	CreatedTime  time.Time    `json: created_time`
	Community    Community    `json: commnity`
	DeletedTime  null.Time    `json: deleted_time`
	From         string       `json: from`
	Message      string       `json: message`
	MessageTags  []MessageTag `json: message_tags`
	ParmalinkUrl string       `json: parmalink_url`
	PostID       string       `json: post_id`
	TargetType   string       `json:"target_type"`
	Type         string       `json: type`
	Verb         Verb         `json: verb`
}
