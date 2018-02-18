package decode

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
