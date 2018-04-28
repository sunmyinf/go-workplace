package decode

import (
	"time"

	"github.com/guregu/null"
)

type Conversations struct {
	PageID          int64  `json:"page_id"`
	ThreadID        int64  `json:"thread_id"`
	ScopedThreadKey string `json:"scoped_thread_key"`
}

type Feed struct {
	EditedTime null.nullTime `json:"edited_time"`
	From       From          `json:"from"`
	Post       struct {
		Type         string        `json:"type"`
		StatusType   string        `json:"status_type"`
		IsPublished  bool          `json:"is_published"`
		UpdatedTime  null.NullTime `json:"updated_time"`
		PermalinkURL string        `json:"permalink_url"`
	}
	Link            string    `json:"link"`
	Message         string    `json:"message"`
	Photo           string    `json:"photo"`
	PhotoIDs        []string  `json:"photo_ids"`
	Photos          []string  `json:"photos"`
	PostID          string    `json:"post_id"`
	Title           string    `json:"title"`
	Video           string    `json:"video"`
	VideoFlagReason int32     `json:"video_flag_reason"`
	Action          string    `json:"action"`
	AlbumID         string    `json:"album_id"`
	CommentID       string    `json:"comment_id"`
	CreatedTime     time.Time `json:"created_time"`
	EventID         string    `json:"event_id"`
	Item            string    `json:"item"`
	ParentID        string    `json:"parent_id"`
	ReactionType    string    `json:"reaction_type"`
	Published       int32     `json:"published"`
	RecipientID     string    `json:"recipient_id"`
	SharedID        string    `json:"shared_id"`
	Verb            Verb      `json:"verb"`
	VideoID         string    `json:"video_id"`
}

type From struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type LiveVideos struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type Mention struct {
	Community    Community    `json:"community"`
	From         From         `json:"from"`
	Link         string       `json:"link"`
	Message      string       `json:"message"`
	MessageTags  []MessageTag `json:"message_tags"`
	Photo        string       `json:"photo"`
	Photos       []string     `json:"photos"`
	PostID       string       `json:"post_id"`
	Video        string       `json:"video"`
	Action       string       `json:"action"`
	AlbumID      string       `json:"album_id"`
	CommentID    string       `json:"comment_id"`
	CreatedTime  time.Time    `json:"created_time"`
	EventID      string       `json:"event_id"`
	Item         string       `json:"item"`
	ParentID     string       `json:"parent_id"`
	PhotoID      string       `json:"photo_id"`
	ReactionType string       `json:"reaction_type"`
	Published    int32        `json:"published"`
	RecipientID  string       `json:"recipient_id"`
	SharedID     string       `json:"shared_id"`
	Verb         Verb         `json:"verb"`
	VideoID      string       `json:"video_id"`
}

type Videos struct {
	ID     string `json:"id"`
	Status struct {
		VideoStatus string `json:"video_status"`
	} `json:"video_status"`
}
