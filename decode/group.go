package decode

type Group struct {
	Posts      []Post     `json: posts`
	Comments   []Comment  `json: comments`
	Membership Membership `json: membership`
}

type Community struct {
	ID string `json: id`
}
