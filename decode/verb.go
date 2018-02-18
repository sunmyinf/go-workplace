package decode

type Verb string

const (
	VerbAdd     Verb = "add"
	VerbBlock   Verb = "block"
	VerbEdit    Verb = "edit"
	VerbEdited  Verb = "edited"
	VerbDelete  Verb = "delete"
	VerbFollow  Verb = "follow"
	VerbHide    Verb = "hide"
	VerbMute    Verb = "mute"
	VerbRemove  Verb = "remove"
	VerbUnblock Verb = "unblock"
	VerbUnhide  Verb = "unhide"
	VerbUpdate  Verb = "update"
)
